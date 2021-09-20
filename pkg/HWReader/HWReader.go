package HWReader

import (
	"XU4PowerManager/pkg/Server"
	"fmt"
	"io/ioutil"
	"time"
)

var enabled = false
var UpdateFrequency = 10.
var UpdateTime = time.Second / time.Duration(UpdateFrequency)

var Thermal0 = "/sys/devices/virtual/thermal/thermal_zone0/temp"

func Start() {
	enabled = true

	for enabled == true {
		content, _ := ioutil.ReadFile(Thermal0)
		fmt.Println(string(content))
		Server.Temp = string(content)

		fmt.Println("Update")
		time.Sleep(UpdateTime)
	}
}

func Stop() {
	enabled = false
}
