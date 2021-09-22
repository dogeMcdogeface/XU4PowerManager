package HWReader

import "time"

type SystemStatus struct {
	Time  time.Time
	Therm map[string]interface{}
	Freq  map[string]interface{}
	Stats map[string]interface{}
	Fans  map[string]interface{}
}
