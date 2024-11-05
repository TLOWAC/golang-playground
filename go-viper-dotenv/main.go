package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	LoadEnvConfiguration(".", "config", "env")
	fmt.Println("SERVER_HOST", Config.Server_Port)
	fmt.Println("SERVER_HOST", viper.Get("SERVER_PORT"))
}

type Configuration struct {
	Server_Port string
}

var Config *Configuration

func LoadEnvConfiguration(configPath, configName, configType string) error {
	var config *Configuration

	// config 파일 경로 설정
	viper.AddConfigPath(configPath)
	// config 파일 이름 설정
	viper.SetConfigName(configName)
	// config 파일 타입 설정 ( yaml / toml / env )
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Errorf("could not unmarshal: %v", err)
	}

	Config = config
	return nil
}
