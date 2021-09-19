package main

import (
	"XU4PowerManager/internal"
	"fmt"
	_ "github.com/spf13/viper"
	"time"
)

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)

	for true {

		fmt.Println("Update")
		time.Sleep(internal.UpdateTime)
	}
}
