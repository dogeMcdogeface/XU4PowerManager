package HWReader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

var enabled = false

var UpdateTime = 100 * time.Millisecond
var HistoryDuration = 5 * time.Second

var Sensors = map[string]string{
	"Thermal0": "/sys/devices/virtual/thermal/thermal_zone0/temp",
}

var HistoryLength = int((5 * time.Second) / UpdateTime)
var HistoryIndex = 0

var lastRead = map[string]interface{}{}
var history = make([]map[string]interface{}, HistoryLength)

var lock sync.Mutex

/*****************************************************************/
func Start() {
	enabled = true

	for enabled == true {
		lock.Lock()

		lastRead["Time"] = time.Now()
		for key, value := range Sensors {
			lastRead[key] = readSensor(value)
		}
		history[HistoryIndex] = lastRead
		HistoryIndex = (HistoryIndex + 1) % HistoryLength

		//fmt.Println(lastRead)

		lock.Unlock()
		time.Sleep(UpdateTime)
	}
}

func Stop() {
	enabled = false
}

func readSensor(path string) int {
	in, _ := ioutil.ReadFile(path)
	inTxt := strings.TrimSpace(string(in))
	value, _ := strconv.Atoi(inTxt)

	fmt.Println(path, value)
	return value
}

func GetLast() map[string]interface{} {
	lock.Lock()
	defer lock.Unlock()
	return lastRead
}
func GetHistory() []map[string]interface{} {
	lock.Lock()
	defer lock.Unlock()
	return history
}
