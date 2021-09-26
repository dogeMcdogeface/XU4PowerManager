package main

import (
	"fmt"
	_ "github.com/spf13/viper"

	"XU4PowerManager/internal"
	"XU4PowerManager/pkg/HWReader"
	"XU4PowerManager/pkg/Server"
)

/*func check(e error) {
	if e != nil {
		panic(e)
	}
}*/

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)
	//HWWriter.Echo("/sys/class/leds/blue\\:heartbeat/trigger", "none")

	go Server.Start()
	HWReader.Start()

}
