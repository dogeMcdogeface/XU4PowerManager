package HWReader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

var enabled = false
var UpdateFrequency = 10.
var UpdateTime = time.Second / time.Duration(UpdateFrequency)

var Sensors = map[string]string{
	"Thermal0": "/sys/devices/virtual/thermal/thermal_zone0/temp",
}

var LastRead = map[string]interface{}{}

func Start() {
	enabled = true

	for enabled == true {

		for key, value := range Sensors {
			//fmt.Println("Key:", key, "Value:", value)
			LastRead[key] = readSensor(value)
		}
		fmt.Println(LastRead)

		/*fmt.Println(string(content))
		Server.Temp = string(content)*/

		fmt.Println("Update")
		time.Sleep(UpdateTime)
	}
}

func Stop() {
	enabled = false
}

func readSensor(path string) int {
	content, _ := ioutil.ReadFile(path)
	value, _ := strconv.Atoi(string(content))

	return value
}
