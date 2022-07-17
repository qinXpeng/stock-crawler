package module

import (
	"github.com/BurntSushi/toml"
)

var GlobalConfig SetConfig

func InitGlobalConf(fileph string) {
	_, err := toml.DecodeFile(fileph, &GlobalConfig)
	if err != nil {
		panic("error init GlobalConfig")
	}
}
