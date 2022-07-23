package module

import (

	"github.com/BurntSushi/toml"
)

var GlobalConfig SetConfig
var GlobalRegister RegisterConfig
func InitGlobalConf(fileph string) {
	_, err := toml.DecodeFile(fileph, &GlobalConfig)
	if err != nil {
		panic("error init GlobalConfig")
	}
}

func InitRegister(fileph string){
	_,err := toml.DecodeFile(fileph, &GlobalRegister)
	if err != nil {
		panic("error init GlobalRegister")
	}
}
