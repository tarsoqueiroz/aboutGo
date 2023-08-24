package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type MysqlCfg struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func init() {

	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Erro fatal no arquivo de configuração: \n\n%w\n", err))
	}

}

func main() {

	fmt.Println(viper.GetString("environment"))
	fmt.Println(viper.GetString("databases.mysql.host"))

	b, err := json.Marshal(viper.Get("databases.mysql"))
	if err != nil {
		panic(err)
	}

	var mysql MysqlCfg

	err = json.Unmarshal(b, &mysql)
	if err != nil {
		panic(err)
	}

	fmt.Println(mysql)
	fmt.Println(mysql.Host)

}
