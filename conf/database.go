package conf

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int64 `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
