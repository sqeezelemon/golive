package golive

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const baseUrl = "https://api.infiniteflight.com/public/v2/"

type Client struct {
	client *http.Client
	Key    string
}

// NewClient creates a new golive.Client with the given API key and http.Client
// See [User Guide] for help obtaining the API key
//
// [User Guide]: https://infiniteflight.com/guide/developer-reference/live-api/overview
func NewClient(apikey string, client *http.Client) *Client {
	return &Client{
		client: client,
		Key:    apikey,
	}
}

// Internal method for GET requests
func (c *Client) get(path string) (*json.Decoder, error) {
	var decoder *json.Decoder
	url := baseUrl + path
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return decoder, err
	}

	request.Header.Add("Authorization", "Bearer "+c.Key)
	response, err := c.client.Do(request)
	if err != nil {
		return decoder, err
	}

	decoder = json.NewDecoder(response.Body)
	return decoder, err
}

// Internal method for POST requests
func (c *Client) post(path string, body io.Reader) (*json.Decoder, error) {
	var decoder *json.Decoder
	url := baseUrl + path
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return decoder, err
	}

	request.Header.Add("Authorization", "Bearer "+c.Key)
	request.Header.Add("Content-Type", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		return decoder, err
	}

	decoder = json.NewDecoder(response.Body)
	return decoder, err
}

// GetSessions retrieves all public sessions
func (c *Client) GetSessions() ([]Session, error) {
	var result apiResponse[[]Session]
	decoder, err := c.get("sessions")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetSession retrieves information about a session.
func (c *Client) GetSession(sessionId string) (Session, error) {
	var result apiResponse[Session]
	decoder, err := c.get("sessions/" + sessionId)
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetFlights retrieves all flights for a session.
func (c *Client) GetFlights(sessionId string) ([]Flight, error) {
	var result apiResponse[[]Flight]
	decoder, err := c.get("sessions/" + sessionId + "/flights")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetFlight retrieves information about a specific flight in a session.
func (c *Client) GetFlight(sessionId string, flightId string) (Flight, error) {
	var result apiResponse[Flight]
	decoder, err := c.get("sessions/" + sessionId + "/flights/" + flightId)
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetFlightRoute retrieves the flown path for a flight.
func (c *Client) GetFlightRoute(sessionId string, flightId string) ([]PositionReport, error) {
	var result apiResponse[[]PositionReport]
	decoder, err := c.get("sessions/" + sessionId + "/flights/" + flightId + "/route")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetFlightPlan retrieves a detailed flight plan for a flight.
func (c *Client) GetFlightPlan(sessionId string, flightId string) (FlightPlan, error) {
	var result apiResponse[FlightPlan]
	decoder, err := c.get("sessions/" + sessionId + "/flights/" + flightId + "/flightplan")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetActiveAtc retrieves all active ATC frequencies for a session.
func (c *Client) GetActiveAtc(sessionId string) ([]ActiveAtcFacility, error) {
	var result apiResponse[[]ActiveAtcFacility]
	decoder, err := c.get("sessions/" + sessionId + "/atc")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserStats retrieves stats about up to 25 users at once.
func (c *Client) GetUserStats(userIds []string, usernames []string, hashes []string) ([]UserStats, error) {
	bodyMap := map[string][]string{
		"userIds":        userIds,
		"discourseNames": usernames,
		"userHashes":     hashes,
	}
	body, _ := json.Marshal(bodyMap)
	var result apiResponse[[]UserStats]
	decoder, err := c.post("users", bytes.NewReader(body))
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserGrade retrieves detailed grade table for a user.
func (c *Client) GetUserGrade(userId string) (UserGrade, error) {
	var result apiResponse[UserGrade]
	decoder, err := c.get("users/" + userId)
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetAtis retrieves ATIS for an airport in a session.
func (c *Client) GetAtis(sessionId string, icao string) (string, error) {
	var result apiResponse[string]
	decoder, err := c.get("sessions/" + sessionId + "/airport/" + icao + "/atis")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetAirportStatus retrieves ATC and inbound/outbound aircraft information for an airport.
func (c *Client) GetAirportStatus(sessionId string, icao string) (AirportStatus, error) {
	var result apiResponse[AirportStatus]
	decoder, err := c.get("sessions/" + sessionId + "/airport/" + icao + "/status")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetWorldStatus retrieves ATC and inbound/outbound aircraft information for all airports in a session.
func (c *Client) GetWorldStatus(sessionId string) ([]AirportStatus, error) {
	var result apiResponse[[]AirportStatus]
	decoder, err := c.get("sessions/" + sessionId + "/world")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetTracks retrieves all currently active tracks.
func (c *Client) GetTracks() ([]Track, error) {
	var result apiResponse[[]Track]
	decoder, err := c.get("tracks")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserFlights retrieves a page from the flight logbook for a user.
func (c *Client) GetUserFlights(userId string, page int) (FlightLogbookPage, error) {
	var result apiResponse[FlightLogbookPage]
	decoder, err := c.get("users/" + userId + "/flights?page=" + strconv.Itoa(page))
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserFlight retrieves a flight from the user's logbook.
func (c *Client) GetUserFlight(userId string, flightId string) (LoggedFlight, error) {
	var result apiResponse[LoggedFlight]
	decoder, err := c.get("users/" + userId + "/flights/" + flightId)
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserAtcSessions retrieves a page from the ATC logbook for a user.
func (c *Client) GetUserAtcSessions(userId string, page int) (AtcLogbookPage, error) {
	var result apiResponse[AtcLogbookPage]
	decoder, err := c.get("users/" + userId + "/atc?page=" + strconv.Itoa(page))
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetUserAtcSession retrieves an ATC session from the user's logbook.
func (c *Client) GetUserAtcSession(userId string, atcSessionId string) (LoggedAtcSession, error) {
	var result apiResponse[LoggedAtcSession]
	decoder, err := c.get("users/" + userId + "/atc/" + atcSessionId)
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetNotams retrieves NOTAMs for a session.
func (c *Client) GetNotams(sessionId string) ([]Notam, error) {
	var result apiResponse[[]Notam]
	decoder, err := c.get("sessions/" + sessionId + "/notams")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetAircraft retrieves all aircraft models.
func (c *Client) GetAircraft() ([]Aircraft, error) {
	var result apiResponse[[]Aircraft]
	decoder, err := c.get("aircraft")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetAircraftLiveries retrieves all liveries for an aircraft.
func (c *Client) GetAircraftLiveries(aircraftId string) ([]Livery, error) {
	var result apiResponse[[]Livery]
	decoder, err := c.get("aircraft/" + aircraftId + "/liveries")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}

// GetLiveries retrieves all liveries.
func (c *Client) GetLiveries() ([]Livery, error) {
	var result apiResponse[[]Livery]
	decoder, err := c.get("aircraft/liveries")
	if err != nil {
		return result.Result, err
	}
	err = decoder.Decode(&result)
	if result.ErrorCode != 0 {
		err = ApiError(result.ErrorCode)
	}
	return result.Result, err
}
