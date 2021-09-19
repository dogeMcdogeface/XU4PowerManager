package main

import (
	"XU4PowerManager/internal"
	"bufio"
	"fmt"
	_ "github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting XU4 Power Manager ver." + internal.Ver)

	for true {
		file, err := os.Open(internal.Thermal0)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		fmt.Println(scanner.Text())

		fmt.Println("Update")
		time.Sleep(internal.UpdateTime)
	}
}
