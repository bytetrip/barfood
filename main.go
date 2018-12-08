package main

import (
	"github.com/bytetrip/barfood/bar"
	"time"
)

func main() {

	b := bar.Start()

	b.Scroll("Good evening to you sir! I hope you are doing well today!")
	time.Sleep(3 * time.Second)

	for {
		b.Update("TADAAAAAAAAAAAA!!!")
		time.Sleep(50 * time.Millisecond)
	}
}
