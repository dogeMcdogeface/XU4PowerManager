package HWWriter

import (
	"fmt"
	"log"
	"os"
)

func Echo(path, txt string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(txt)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
