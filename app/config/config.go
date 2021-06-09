package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	ConfigPath string
	Conf       = &Config{}
)

//TODO 后期所有的配置独立化 作为公共库配置
type Config struct {
	MysqlDB *Mysql      `toml:"mysql_db"`
	MongoDB *Mongo      `toml:"mongo"`
	HTTPSvc *HTTPServer `toml:"http_svc"`
	Grpc    *Grpc       `toml:"grpc"`
	Log     *Log        `toml:"log"`
	Nsq     *Nsq        `toml:"nsq"`
}

type baseDBDefint struct {
	HTTPServer
	User      string `toml:"user"`
	Password  string `toml:"password"`
	Database  string `toml:"database"`
	IsLogOpen bool   `toml:"is_log_open"`
}

type HTTPServer struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}
type Grpc struct {
	Port string `toml:"port"`
}
type Log struct {
	Level int `toml:"level"`
}
type Nsq struct {
	Address string `toml:"address"`
	Topic   string `toml:"topic"`
	Channel string `toml:"channel"`
}

func init() {
	flag.StringVar(&ConfigPath, "conf", "./config.toml", "default config path")
}

func configInit() (conf *Config) {
	if _, err := toml.DecodeFile(ConfigPath, &conf); err != nil {
		panic(err)
	}
	return
}

func Init() *Config {
	return configInit()
}
