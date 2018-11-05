package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

/**
要做的事情非常简单:
1.编写配置文件信息;
2.根据配置文件编写相应的结构体;
3.根据toml的方式读取配置文件,然后赋予这个结构体;
4.把获取的信息设置给指定的服务即可.
*/
type configStruct struct {
	Env string `toml:"env"`

	Http struct {
		Enabled bool   `toml:"enabled"`
		Listen  string `toml:"listen"`
	}

	Mysql struct {
		DB struct {
			Datasource string `toml:"datasource"`
			Timeout    int    `toml:"timeout"`
		} `toml:"db"`
	} `toml:"mysql"`

	Redis struct {
		Default struct {
			Addr     string `toml:"addr"`
			Password string `toml:"password"`
			DB       int    `toml:"db"`
			Timeout  int    `toml:"timeout"`
		} `toml:"default"`
	} `toml:"redis"`

	Log struct {
		LogLevel      string `toml:"log_level"`
		LogDir        string `toml:"log_dir"`
		LogName       string `toml:"log_name"`
		LogRotateMode string `toml:"log_rotate_mode"`
		Maxsize       uint64 `toml:"max_size"`
	} `toml:"log"`

	Map struct {
		ShareChannel map[string]int `toml:"shareChannel"`
	} `toml:"map"`

	Weixin struct {
		ApiUrl string            `toml:"api_url"`
		Secret string            `toml:"secret"`
		Midas  map[string]string `toml:"midas"`
	} `toml:"weixin"`
}

var C *configStruct

// go run main.go  -config=/XXX/config/dev.toml
var filename = flag.String("config", "config/dev.toml", "Location of the config file.")

func init() {
	flag.Parse()
	if _, err := toml.DecodeFile(*filename, &C); err != nil {
		log.Fatalln(err)
	}
}
