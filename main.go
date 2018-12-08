package main

import (
	"fmt"
	"github.com/bytetrip/barfood/bar"
	"os"
	"time"
)

func main() {

	b := bar.Start()

	for {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		batt := getBatteryText()
		text := fmt.Sprintf("%s %s", batt, timestamp)
		b.Update(fmt.Sprintf("%239s", text))
		time.Sleep(50 * time.Millisecond)
	}
}

func getBatteryText() string {

	file, err := os.Open("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		panic(err)
	}

	data := make([]byte, 3)
	file.Read(data)
	return fmt.Sprintf("%s", string(data[:len(data)-1]))

}
