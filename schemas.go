package golive

import "time"

type apiResponse[T any] struct {
	ErrorCode int `json:"errorCode"`
	Result    T   `json:"result"`
}

type Session struct {
	MaxUsers  int    `json:"maxUsers"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserCount int    `json:"UserCount"`
	Type      int    `json:"type"`
}

type Flight struct {
	Username            string       `json:"username"`
	Callsign            string       `json:"callsign"`
	Latitude            float64      `json:"latitude"`
	Longitude           float64      `json:"longitude"`
	Altitude            float64      `json:"altitude"`
	Speed               float64      `json:"speed"`
	VerticalSpeed       float64      `json:"verticalSpeed"`
	Track               float64      `json:"track"`
	LastReport          TimeWithoutT `json:"lastReport"`
	Id                  string       `json:"flightId"`
	UserId              string       `json:"userId"`
	AircraftId          string       `json:"aircraftId"`
	LiveryId            string       `json:"liveryId"`
	Heading             float64      `json:"heading"`
	VirtualOrganization string       `json:"virtualOrganization"`
}

type PositionReport struct {
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Altitude    float64   `json:"altitude"`
	Track       float64   `json:"track"`
	GroundSpeed float64   `json:"groundSpeed"`
	Date        time.Time `json:"date"`
}

type FlightPlan struct {
	Id              string           `json:"flightPlanId"`
	FlightId        string           `json:"flightId"`
	Waypoints       []string         `json:"waypoints"`
	LastUpdate      TimeWithoutT     `json:"lastUpdate"`
	FlightPlanItems []FlightPlanItem `json:"flightPlanItems"`
}

type FlightPlanItem struct {
	Name       string           `json:"name"`
	Type       int              `json:"type"`
	Children   []FlightPlanItem `json:"children"`
	Identifier string           `json:"identifier"`
	Altitude   int              `json:"altitude"`
	Location   Location         `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

type ActiveAtcFacility struct {
	FrequencyId         string       `json:"frequencyId"`
	UserId              string       `json:"userId"`
	Username            string       `json:"username"`
	VirtualOrganization string       `json:"virtualOrganization"`
	AirportName         string       `json:"airportName"`
	Type                int          `json:"type"`
	Latitude            float64      `json:"latitude"`
	Longitude           float64      `json:"longitude"`
	StartTime           TimeWithoutT `json:"startTime"`
}

type UserStats struct {
	OnlineFlights         int            `json:"onlineFlights"`
	Violations            int            `json:"violations"`
	Xp                    int            `json:"xp"`
	LandingCount          int            `json:"landingCount"`
	FlightTime            int            `json:"flightTime"`
	AtcOperations         int            `json:"atcOperations"`
	AtcRank               int            `json:"atcRank"`
	Grade                 int            `json:"grade"`
	Hash                  string         `json:"hash"`
	ViolationCountByLevel ViolationCount `json:"violationCountByLevel"`
	Roles                 []int          `json:"roles"`
	UserId                string         `json:"userId"`
	VirtualOrganization   string         `json:"virtualOrganization"`
	DiscourseUsername     string         `json:"discourseUsername"`
	ErrorCode             int            `json:"errorCode"`
}

type UserGrade struct {
	Total12MonthsViolations int                `json:"total12MonthsViolations"`
	GradeDetails            GradeConfiguration `json:"gradeDetails"`
	TotalXP                 int                `json:"totalXP"`
	AtcOperations           int                `json:"atcOperations"`
	AtcRank                 int                `json:"atcRank"`
	LastLevel1ViolationDate time.Time          `json:"lastLevel1ViolationDate"`
	LastLevel2ViolationDate time.Time          `json:"lastLevel2ViolationDate"`
	LastLevel3ViolationDate time.Time          `json:"lastLevel3ViolationDate"`
	LastReportViolationDate time.Time          `json:"lastReportViolationDate"`
	ViolationCountByLevel   ViolationCount     `json:"violationCountByLevel"`
	Roles                   []int              `json:"roles"`
	UserId                  string             `json:"userId"`
	VirtualOrganization     string             `json:"virtualOrganization"`
	DiscourseUsername       string             `json:"discourseUsername"`
	Groups                  []string           `json:"groups"`
	ErrorCode               int                `json:"errorCode"`
}

type ViolationCount struct {
	Level1 int `json:"level1"`
	Level2 int `json:"level2"`
	Level3 int `json:"level3"`
}

type GradeConfiguration struct {
	Grades          []Grade               `json:"grades"`
	GradeIndex      int                   `json:"gradeIndex"`
	RuleDefinitions []GradeRuleDefinition `json:"ruleDefinitions"`
}

type Grade struct {
	Rules []GradeRule `json:"rules"`
	Index int         `json:"index"`
	Name  string      `json:"name"`
	State int         `json:"state"`
}

type GradeRule struct {
	RuleIndex            int                 `json:"ruleIndex"`
	ReferenceValue       float64             `json:"referenceValue"`
	UserValue            float64             `json:"userValue"`
	State                int                 `json:"state"`
	UserValueString      string              `json:"userValueString"`
	ReferenceValueString string              `json:"referenceValueString"`
	Definition           GradeRuleDefinition `json:"definition"`
}

type GradeRuleDefinition struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Property    string  `json:"property"`
	Operator    int     `json:"operator"`
	Period      float64 `json:"period"`
	Order       int     `json:"order"`
	Group       int     `json:"group"`
}

type AirportStatus struct {
	AirportIcao          string              `json:"airportIcao"`
	InboundFlightsCount  int                 `json:"inboundFlightsCount"`
	InboundFlights       []string            `json:"inboundFlights"`
	OutboundFlightsCount int                 `json:"outboundFlightsCount"`
	OutboundFlights      []string            `json:"outboundFlights"`
	AtcFacilities        []ActiveAtcFacility `json:"atcFacilities"`
}

type Track struct {
	Name       string    `json:"name"`
	Path       []string  `json:"path"`
	EastLevels []int     `json:"eastLevels"`
	WestLevels []int     `json:"westLevels"`
	Type       string    `json:"type"`
	LastSeen   time.Time `json:"lastSeen"`
}

type LogbookPage[T any] struct {
	PageIndex       int  `json:"pageIndex"`
	TotalPages      int  `json:"totalPages"`
	TotalCount      int  `json:"totalCount"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	HasNextPage     bool `json:"hasNextPage"`
	Data            []T  `json:"data"`
}

type FlightLogbookPage LogbookPage[LoggedFlight]
type LoggedFlight struct {
	Id                 string  `json:"id"`
	Created            string  `json:"created"`
	UserId             string  `json:"userId"`
	AircraftId         string  `json:"aircraftId"`
	LiveryId           string  `json:"liveryId"`
	Callsign           string  `json:"callsign"`
	Server             string  `json:"server"`
	DayTime            float64 `json:"dayTime"`
	NightTime          int     `json:"nightTime"`
	TotalTime          float64 `json:"totalTime"`
	LandingCount       int     `json:"landingCount"`
	OriginAirport      string  `json:"originAirport"`
	DestinationAirport string  `json:"destinationAirport"`
	Xp                 int     `json:"xp"`
}

type AtcLogbookPage LogbookPage[LoggedAtcSession]
type LoggedAtcSession struct {
	Id             string      `json:"id"`
	SessionGroupId string      `json:"atcSessionGroupId"`
	Facility       AtcFacility `json:"facility"`
	Created        string      `json:"created"`
	Updated        string      `json:"updated"`
	Operations     int         `json:"operations"`
	TotalTime      float64     `json:"totalTime"`
}

type AtcFacility struct {
	Id        string  `json:"id"`
	Icao      string  `json:"airportIcao"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      int     `json:"frequencyType"`
}

type Notam struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Type      int     `json:"type"`
	SessionId string  `json:"sessionId"`
	Radius    int     `json:"radius"`
	Message   string  `json:"message"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Icao      string  `json:"icao"`
	Floor     int     `json:"floor"`
	Ceiling   int     `json:"ceiling"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
}

type Aircraft struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Livery struct {
	Id           string `json:"id"`
	AircraftID   string `json:"aircraftID"`
	AircraftName string `json:"aircraftName"`
	LiveryName   string `json:"liveryName"`
}
