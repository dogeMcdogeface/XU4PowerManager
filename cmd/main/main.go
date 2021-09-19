package main

import (
	"XU4PowerManager/internal"
	"fmt"
	_ "github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)

	content, err := ioutil.ReadFile(internal.Thermal0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	for true {

		fmt.Println("Update")
		time.Sleep(internal.UpdateTime)
	}
}
