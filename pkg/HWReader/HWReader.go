package HWReader

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

/***************** CONFIGURATION VARIABLES *************************/

var UpdateTime = 100 * time.Millisecond
var HistoryDuration = 5 * time.Second

var Sensors = map[string]string{
	"Thermal0": "/sys/devices/virtual/thermal/thermal_zone0/temp",
	"Freq0":    "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq",
}

/***************** RUNTIME VARIABLES *******************************/
var HistoryLength = int((5 * time.Second) / UpdateTime)
var HistoryIndex = 0

var lastRead = map[string]interface{}{}
var history = make([]map[string]interface{}, HistoryLength)

var enabled = false
var lock sync.Mutex

/***************** MAIN METHOD *************************************/
func Start() {
	enabled = true

	for enabled == true {
		lock.Lock()

		lastRead = map[string]interface{}{}
		lastRead["Time"] = time.Now()
		for key, value := range Sensors {
			lastRead[key] = readSensor(value)
		}
		history[HistoryIndex] = lastRead
		HistoryIndex = (HistoryIndex + 1) % HistoryLength

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
	return value
}

/***************** GETTERS *****************************************/

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
