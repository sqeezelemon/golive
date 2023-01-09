package golive

import (
	"net/http"
	"os"
	"testing"
)

func client() *Client {
	apikey := os.Getenv("APIKEY")
	return NewClient(apikey, &http.Client{})
}

func TestSessions(t *testing.T) {
	sessions, err := client().GetSessions()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetSession(sessions[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestFlights(t *testing.T) {
	sessions, err := client().GetSessions()

	flights, err := client().GetFlights(sessions[0].Id)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetFlight(sessions[0].Id, flights[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestFlightRoute(t *testing.T) {
	sessions, err := client().GetSessions()
	flights, err := client().GetFlights(sessions[0].Id)

	_, err = client().GetFlightRoute(sessions[0].Id, flights[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestFlightPlan(t *testing.T) {
	sessions, err := client().GetSessions()
	flights, err := client().GetFlights(sessions[0].Id)

	_, err = client().GetFlightPlan(sessions[0].Id, flights[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestActiveAtc(t *testing.T) {
	sessions, err := client().GetSessions()

	_, err = client().GetActiveAtc(sessions[0].Id)
	if err != nil {
		t.Error(err)
	}
}

//func TestAtis(t *testing.T) {
//	sessions, err := client().GetSessions()
//
//	_, err = client().GetAtis(sessions[0].Id, "KLAX")
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestUserStats(t *testing.T) {
	_, err := client().GetUserStats([]string{
		"2a11e620-1cc1-4ac6-90d1-18c4ed9cb913",
		"5917d076-88a5-40e7-95e0-8818748f8e99",
	}, []string{
		"KaiM",
		"Laura",
	}, []string{
		"F0081CAA",
		"E2087C9F",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUserGrade(t *testing.T) {
	_, err := client().GetUserGrade("2a11e620-1cc1-4ac6-90d1-18c4ed9cb913")
	if err != nil {
		t.Error(err)
	}
}

func TestAirportStatus(t *testing.T) {
	sessions, err := client().GetSessions()

	_, err = client().GetWorldStatus(sessions[0].Id)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetAirportStatus(sessions[0].Id, "KLAX")
	if err != nil {
		t.Error(err)
	}
}

func TestTracks(t *testing.T) {
	_, err := client().GetTracks()
	if err != nil {
		t.Error(err)
	}
}

func TestUserFlights(t *testing.T) {
	flights, err := client().GetUserFlights("2a11e620-1cc1-4ac6-90d1-18c4ed9cb913", 1)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetUserFlight("2a11e620-1cc1-4ac6-90d1-18c4ed9cb913", flights.Data[0].Id)
}

func TestUserAtcSessions(t *testing.T) {
	sessions, err := client().GetUserAtcSessions("2a11e620-1cc1-4ac6-90d1-18c4ed9cb913", 1)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetUserAtcSession("2a11e620-1cc1-4ac6-90d1-18c4ed9cb913", sessions.Data[0].Id)
}

func TestNotams(t *testing.T) {
	sessions, err := client().GetSessions()

	_, err = client().GetNotams(sessions[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestAircraft(t *testing.T) {
	_, err := client().GetAircraft()
	if err != nil {
		t.Error(err)
	}
}

func TestLiveries(t *testing.T) {
	liveries, err := client().GetLiveries()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client().GetAircraftLiveries(liveries[0].AircraftID)
	if err != nil {
		t.Error(err)
	}
}
