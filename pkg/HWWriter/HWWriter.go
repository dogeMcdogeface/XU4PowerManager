package HWWriter

import (
	"fmt"
	"os"
	"os/exec"
)

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
