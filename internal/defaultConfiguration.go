package internal

import "time"

var Ver = "12.3"

var thermal0 = "/sys/devices/virtual/thermal/thermal_zone0/temp"

var UpdateFrequency = 1.
var UpdateTime = time.Second / time.Duration(UpdateFrequency)
