package HWWriter

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Echo(path, txt string) {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}
