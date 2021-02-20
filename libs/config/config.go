package config

import (
	"github.com/spf13/viper"
	"os"
)

var env string
var allConf map[string]*viper.Viper

func loadConfigFile(key string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(key)
	//v.SetConfigType("yaml")

	file := "./config/" + env + "/" + key + ".yml"
	_, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	v.SetConfigFile(file)
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func init() {
	env = os.Getenv("CONFIG_ENV")
	if len(env) == 0 {
		env = "local"
	}
	allConf = make(map[string]*viper.Viper)
}

func GetConf(conf, key string) (map[string]interface{}, error) {
	if _, ok := allConf[conf]; !ok {
		config, err := loadConfigFile(conf)
		if err != nil {
			return nil, err
		} else {
			allConf[conf] = config
		}
	}
	return allConf[conf].GetStringMap(key), nil
}
