package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Driver			string `json:"driver"`
	PortPostgre     string `json:"portPostgre"`
	HostPostgre		string `json:"hostPostgre"`
	DBname 			string `json:"dbname"`
	UserNamePostgre string `json:"userNamePostgre"`
	PasswordPostgre string `json:"passwordPostgre"`
	Sslmode			string `json:"sslmode"`
}

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
	}

	Conf = Config{
		Driver: 		 viper.GetString("Postgre.driver"),
		HostPostgre:     viper.GetString("Postgre.host"),
		PortPostgre:     viper.GetString("Postgre.port"),
		DBname: 		 viper.GetString("Postgre.dbname"),
		UserNamePostgre: viper.GetString("Postgre.userName"),
		PasswordPostgre: viper.GetString("Postgre.password"),
		Sslmode:     	 viper.GetString("Postgre.sslmode"),
	}

}
func (config Config) String() string {
	out, err := json.Marshal(config)
	if err != nil {
		return ""
	}
	return (string(out))
}

func GetConfig() Config {
	return Conf
}
