package main

import (
	"XU4PowerManager/internal"
	"XU4PowerManager/pkg/Server"
	"fmt"
	_ "github.com/spf13/viper"
	"io/ioutil"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)
	fmt.Println("Server " + Server.Status)

	Server.Start()

	for true {

		content, _ := ioutil.ReadFile(internal.Thermal0)
		fmt.Println(string(content))
		Server.Temp = string(content)

		fmt.Println("Update")
		time.Sleep(internal.UpdateTime)
	}
}
