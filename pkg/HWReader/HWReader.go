package HWReader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

/***************** CONFIGURATION VARIABLES *************************/
var UpdateTime = 500 * time.Millisecond
var LogTime = 2 * time.Second
var cacheLifetime = 10 * time.Millisecond

var Sensors = SystemStatus{
	Therm: map[string]interface{}{
		"Thermal0":   "/sys/devices/virtual/thermal/thermal_zone0/temp",
		"Thermal1":   "/sys/devices/virtual/thermal/thermal_zone1/temp",
		"Thermal2":   "/sys/devices/virtual/thermal/thermal_zone2/temp",
		"Thermal3":   "/sys/devices/virtual/thermal/thermal_zone3/temp",
		"ThermalGpu": "/sys/devices/virtual/thermal/thermal_zone4/temp",
	}, Freq: map[string]interface{}{
		"Freq0":   "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq",
		"Freq1":   "/sys/devices/system/cpu/cpu1/cpufreq/scaling_cur_freq",
		"Freq2":   "/sys/devices/system/cpu/cpu2/cpufreq/scaling_cur_freq",
		"Freq3":   "/sys/devices/system/cpu/cpu3/cpufreq/scaling_cur_freq",
		"Freq4":   "/sys/devices/system/cpu/cpu4/cpufreq/scaling_cur_freq",
		"Freq5":   "/sys/devices/system/cpu/cpu5/cpufreq/scaling_cur_freq",
		"Freq6":   "/sys/devices/system/cpu/cpu6/cpufreq/scaling_cur_freq",
		"Freq7":   "/sys/devices/system/cpu/cpu7/cpufreq/scaling_cur_freq",
		"FreqGpu": "/sys/devices/platform/soc/11800000.gpu/devfreq/11800000.gpu/cur_freq",
	}, Fans: map[string]interface{}{
		"FanSpeed": "/sys/devices/platform/pwm-fan/hwmon/hwmon0/pwm1",
	}, Stats: map[string]interface{}{},
}

/***************** RUNTIME VARIABLES *******************************/
var enabled = false
var cachedSystemStatus = SystemStatus{}
var lock sync.Mutex

/***************** MAIN METHOD *************************************/
func Start() {
	enabled = true

	for enabled == true {
		var s = GetSystemStatus() //poll hardware

		//calculate averages

		fmt.Println(s)
		time.Sleep(UpdateTime)
	}
}

func Stop() {
	enabled = false
}

func readFile(path string) int {
	in, _ := ioutil.ReadFile(path)
	inTxt := strings.TrimSpace(string(in))
	value, _ := strconv.Atoi(inTxt)
	return value
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
		s.Freq[key] = readFile(value.(string))
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
