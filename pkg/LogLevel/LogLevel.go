package LogLevel

import (
	"fmt"
	"time"
)

var level = -1

/*********** FORMAT OPTIONS *******************/
var separator = " "
var prefix = ""
var suffix = ""
var dateFormat = "15:04"
var date = formatDate()

func formatDate() string {
	return (time.Now()).Format(dateFormat) + separator
}

func Println(a ...interface{}) {
	var txt = prefix
	txt += date
	txt += fmt.Sprint(a...)
	txt += suffix

	fmt.Println(txt)
}

func Lprintln(l int, a ...interface{}) {
	if l > level {
		var txt = prefix
		txt += date
		txt += fmt.Sprint(a...)
		txt += suffix

		fmt.Println(txt)
	}
}
