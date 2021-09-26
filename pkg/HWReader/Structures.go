package HWReader

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type SystemStatus struct {
	Time  time.Time
	Therm map[string]interface{}
	Freq  map[string]interface{}
	Stats map[string]interface{}
	Fans  map[string]interface{}
}

type SystemAverage struct {
	Time  time.Time
	Therm average
}

func (a SystemAverage) Parse(f *os.File) []SystemAverage {
	var list = []SystemAverage{}
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

func (a *SystemAverage) Compress() []byte {
	var temp = averageSystemStatus.Therm.Value
	var time = averageSystemStatus.Time.Unix()

	var data []byte
	data = append(data, intxToByte(time, 8)...)
	data = append(data, intxToByte(temp, 2)...)

	fmt.Println(averageSystemStatus.Time.Unix(), averageSystemStatus.Therm.Value, data)
	return data
}

func intxToByte(val interface{}, size int) []byte {
	var asByte = make([]byte, size)

	switch size {
	case 1:
		asByte[0] = val.(byte)
		break
	case 2:
		binary.LittleEndian.PutUint16(asByte, uint16(val.(int)))
		break
	case 4:
		binary.LittleEndian.PutUint32(asByte, uint32(val.(int32)))
		break
	case 8:
		binary.LittleEndian.PutUint64(asByte, uint64(val.(int64)))
		break
	}
	return asByte
}

func (a *SystemAverage) LogAverage(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err = f.Write(a.Compress()); err != nil {
		fmt.Println(err)
	}
}

type average struct {
	count int
	Value int
}

func (a average) Find(v ...int) int {
	a.count = 0
	a.Value = 0
	return a.Add(v...)
}
func (a *average) Add(v ...int) int {
	for _, value := range v {
		a.count++
		a.Value = ((a.Value * (a.count - 1)) + (value)) / (a.count)
		//a.Value = int((a.Value * float64(a.count-1)) + float64(value)) / float64(a.count)
	}
	return a.Value
}
