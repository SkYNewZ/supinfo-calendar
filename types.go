package main

type PlanningResponse struct {
	BoosterId            int    `json:"BoosterId"`
	IcsPlanning          string `json:"PlanningICS"`
	LastUpdatedDate      string `json:"LastUpdateDate"`
	LastUpdatedTimestamp int    `json:"LastUpdateDateUnix"`
	MethodCallIsValid    bool   `json:"MethodCallIsValid"`
	MethodCallMessage    string `json:"MethodCallMessage"`
}

type LoginResponse struct {
	BoosterId          int     `json:"BoosterId"`
	Token              string  `json:"Token"`
	LastName           string  `json:"LastName"`
	FirstName          string  `json:"FirstName"`
	Curriculum         string  `json:"Curriculum"`
	CampusClassId      string  `json:"CampusClassId"`
	TotalECTS          float64 `json:"TotalECTS"`
	TotalSuccessPoints float64 `json:"TotalSuccessPoints"`
	// Photo base64.Encoding `json:"Photo"` // IGNORE BASE64 PICTURE
	MethodCallIsValid bool   `json:"MethodCallIsValid"`
	MethodCallMessage string `json:"MethodCallMessage"`
}

type SupinfoStudent struct {
	BoosterId     int
	Token         string
	LastName      string
	FirstName     string
	CampusClassId string
}

type AppConfig struct {
	CampusId       int    `short:"u" long:"campus-id" description:"Your CampusBooster id" required:"true" env:"CAMPUS_ID"`
	CampusPassword string `short:"p" long:"password" description:"Your CampusBooster password" required:"true" env:"CAMPUS_PASSWORD"`
	SupinfoAPIKey  string `short:"k" long:"key" description:"http://campus-api.supinfo.com API KEY" required:"true" env:"SUPINFO_API_KEY"`
	OutputPath     string `short:"o" long:"output-path" description:".ics downloaded file location" required:"false" env:"OUTPUT_PATH" default:"."`
}
