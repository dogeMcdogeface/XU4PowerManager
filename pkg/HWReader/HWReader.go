package HWReader

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/***************** CONFIGURATION VARIABLES *************************/
var UpdateTime = 500 * time.Millisecond
var LogTime = 1 * time.Second
var cacheLifetime = 10 * time.Millisecond

var logPath = "log.csv"

var Sensors = SystemStatus{
	Therm: map[string]interface{}{
		"Thermal0":   "/sys/devices/virtual/thermal/thermal_zone0/temp",
		"Thermal1":   "/sys/devices/virtual/thermal/thermal_zone1/temp",
		"Thermal2":   "/sys/devices/virtual/thermal/thermal_zone2/temp",
		"Thermal3":   "/sys/devices/virtual/thermal/thermal_zone3/temp",
		"ThermalGpu": "/sys/devices/virtual/thermal/thermal_zone4/temp",
	}, Freq: map[string]interface{}{
		"Freq A15": "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq",
		"Freq A7":  "/sys/devices/system/cpu/cpu4/cpufreq/scaling_cur_freq",
		"FreqGpu":  "/sys/devices/platform/soc/11800000.gpu/devfreq/11800000.gpu/cur_freq",
	}, Fans: map[string]interface{}{
		"FanSpeed": "/sys/devices/platform/pwm-fan/hwmon/hwmon0/pwm1",
	}, Stats: map[string]interface{}{
		"LedBlue": "/sys/class/leds/blue\\:heartbeat/brightness",
	},
}

/***************** RUNTIME VARIABLES *******************************/
var enabled = false
var systemMonitoring = true
var systemLogging = true
var systemHandling = true

var cachedSystemStatus = SystemStatus{}
var averageSystemStatus = SystemAverage{}
var lock sync.Mutex

/***************** MAIN METHOD *************************************/
func Start() {
	enabled = true

	for enabled == true {
		var s = GetSystemStatus() //poll hardware

		manageLogging(s) //calculate and log averages
		//manageHardware(s)

		//fmt.Println(s)
		time.Sleep(UpdateTime)
	}
}
func Stop() {
	enabled = false
}

func readSystemStatus() SystemStatus {
	var s = SystemStatus{}
	s.Time = time.Now()
	//Read sensors

	/**** READ TEMPS ****/
	s.Therm = map[string]interface{}{}
	for key, value := range Sensors.Therm {
		s.Therm[key] = byte(readFile(value.(string)) / 1000)
	}
	/**** READ FREQS ****/
	s.Freq = map[string]interface{}{}
	for key, value := range Sensors.Freq {
		var val = readFile(value.(string))
		if !strings.Contains(key, "Gpu") {
			val *= 1000
		}
		s.Freq[key] = val
	}
	/**** READ FANS *****/
	s.Fans = map[string]interface{}{}
	for key, value := range Sensors.Fans {
		s.Fans[key] = byte(readFile(value.(string)) * 100 / 255)
	}
	/**** READ STATS ****/
	s.Stats = map[string]interface{}{}
	for key, value := range Sensors.Stats {
		s.Stats[key] = readFile(value.(string))
	}

	cachedSystemStatus = s
	return s
}
func readFile(path string) int {
	in, _ := ioutil.ReadFile(path)
	inTxt := strings.TrimSpace(string(in))
	value, _ := strconv.Atoi(inTxt)
	return value
}

func manageLogging(s SystemStatus) {
	if time.Since(averageSystemStatus.Time) > LogTime {
		if averageSystemStatus.Time.Second() > 0 {
			averageSystemStatus.LogAverage(logPath) //save average
			//fmt.Println(readLogFile())              //
		}
		averageSystemStatus = SystemAverage{Time: time.Now()}

	}
	for _, value := range s.Therm { //calculate averages
		averageSystemStatus.Therm.Add(int(value.(byte)))
	}
}

func readLogFile() (list []SystemAverage) {
	list = []SystemAverage{}

	f, err := os.Open(logPath)
	if err != nil {
		fmt.Println(err)
		return list
	}
	defer f.Close()

	for true {
		dateByte := make([]byte, 10)
		_, err := f.Read(dateByte)
		if err != nil {
			break
		}
		date := int64(binary.LittleEndian.Uint64(dateByte[0:8]))
		temp := int(binary.LittleEndian.Uint16(dateByte[8:10]))
		//fmt.Println(dateByte, dateByte[0:8], dateByte[8:10], date, temp)
		list = append(list, SystemAverage{
			Time:  time.Unix(date, 0),
			Therm: average{Value: temp},
		})
	}
	return list
}

/***************** GETTERS *****************************************/
func GetSystemStatus() SystemStatus {
	lock.Lock()
	defer lock.Unlock()

	if time.Since(cachedSystemStatus.Time) < cacheLifetime {
		return cachedSystemStatus
	}
	return readSystemStatus()
}

func GetHistory() []map[string]interface{} {
	return nil

}
