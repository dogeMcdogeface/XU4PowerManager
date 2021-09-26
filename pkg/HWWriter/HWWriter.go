package HWWriter

import (
	"log"
	"os/exec"
)

func Echo(path, txt string) {

	cmd := exec.Command("echo ass > t.txt")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
