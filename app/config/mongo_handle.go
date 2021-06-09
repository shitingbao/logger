package conf

type Mongo struct {
	Driver    string `toml:"driver"`
	Database  string `toml:"database"`
	IsLogOpen bool   `toml:"is_log_open"`
}
