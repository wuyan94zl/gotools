package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

func MustConfig(path string, config interface{}) {
	v := viper.New()
	file := filepath.Base(path)
	dir := filepath.Dir(path)
	v.AddConfigPath(dir)
	v.SetConfigFile(file)

	err := v.ReadInConfig()
	if err != nil {
		panic("config file unload")
	}
	err = v.Unmarshal(config)
	if err != nil {
		panic(fmt.Sprintf("config file load err：%v", err))
	}
}
