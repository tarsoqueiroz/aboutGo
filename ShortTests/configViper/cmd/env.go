package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	fmt.Println(viper.Get("URI"))
	fmt.Println(viper.Get("PORT"))
	fmt.Println(viper.Get("SECRET"))

}
