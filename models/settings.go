package models

type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"postgresParams"`
}

type Params struct {
	ServerName string `json:"serverName"`
	PortRun    string `json:"portRun"`
	LogFile    string `json:"logFile"`
	ServerURL  string `json:"serverURL"`
}

type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}
