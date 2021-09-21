package main

import (
	"XU4PowerManager/internal"
	"XU4PowerManager/pkg/HWReader"
	"XU4PowerManager/pkg/Server"
	"fmt"
	_ "github.com/spf13/viper"
)

/*func check(e error) {
	if e != nil {
		panic(e)
	}
}*/

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)

	go Server.Start()
	HWReader.Start()

}
