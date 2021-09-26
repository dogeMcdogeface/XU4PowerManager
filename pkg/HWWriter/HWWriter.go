package HWWriter

import (
	"fmt"
	"os"
	"os/exec"
)

var systemFiles = map[string]string{
	"cpu_policy0":     "/sys/devices/system/cpu/cpufreq/policy0/scaling_governor",
	"cpu_policy4":     "/sys/devices/system/cpu/cpufreq/policy4/scaling_governor",
	"cpu_policies":    "/sys/devices/system/cpu/cpufreq/policy0/scaling_available_governors",
	"cpu_frequencies": "/sys/devices/system/cpu/cpufreq/policy0/scaling_available_frequencies",
	"gpu_policy":      "/sys/devices/platform/soc/11800000.gpu/devfreq/11800000.gpu/governor",
	"gpu_policies":    "/sys/devices/platform/soc/11800000.gpu/devfreq/11800000.gpu/available_governors",
	"fan_speed":       "",
	"led_mode":        "/sys/class/leds/blue\\:heartbeat/trigger",
	"led_brightness":  "/sys/class/leds/blue\\:heartbeat/brightness",
}

func Echo(path, txt string) {

	call("echo none > /sys/class/leds/blue\\:heartbeat/trigger")

}

func call(command string) {
	cmd := exec.Command("sudo", "bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err1 := cmd.Wait()
	if err1 != nil {
		fmt.Println(err1)
	}
}
