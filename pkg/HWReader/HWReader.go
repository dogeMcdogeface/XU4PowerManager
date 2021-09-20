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
var HistoryLength = int((5 * time.Second) / UpdateTime)
var HistoryIndex = 0

var Sensors = map[string]string{
	"Thermal0": "/sys/devices/virtual/thermal/thermal_zone0/temp",
}

var LastRead = map[string]interface{}{}
var History = make([]map[string]interface{}, 1, 1)

func Start() {
	enabled = true

	for enabled == true {

		LastRead["Time"] = time.Now()
		for key, value := range Sensors {
			LastRead[key] = readSensor(value)
		}
		History[0] = LastRead
		HistoryIndex = (HistoryIndex + 1) % HistoryLength

		fmt.Println(LastRead)

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
