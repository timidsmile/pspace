package setting

import (
	"fmt"
	"path/filepath"
	"sync"
	"github.com/BurntSushi/toml"
)


var (
	Cfg * Config
	once sync.Once
)


type DbConfig struct {
	Host string
	Port int
	User string
	Pwd  string
	Name string
}

type Config struct {
	Debug    bool
	PspaceDb       DbConfig
}


func init() {
	ReadConfig()
}

func ReadConfig() *Config {
	once.Do(func() {
		filePath, err := filepath.Abs("./config/dev.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _ , err := toml.DecodeFile(filePath, &Cfg); err != nil {
			panic(err)
		}
	})

	return nil;
}

