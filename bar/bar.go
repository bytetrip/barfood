package bar

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

const columnCount = 240

// Bar : Wrapper for a Bar program
type Bar struct {
	stderr io.ReadCloser
	stdout io.ReadCloser
	stdin  io.WriteCloser
}

// Start Creates the bar process and returns a Bar to operate on
func Start() *Bar {

	b := &Bar{}
	cmd := exec.Command("lemonbar", "-f", "-*-terminus-*-*-*-*-32-*-*-*-*-*-*-*", "-F#0576BF", "-B#040536")

	var err error
	if b.stdin, err = cmd.StdinPipe(); err != nil {
		panic(err)
	}
	if b.stdout, err = cmd.StdoutPipe(); err != nil {
		panic(err)
	}
	if b.stderr, err = cmd.StderrPipe(); err != nil {
		panic(err)
	}

	cmd.Start()

	return b
}

// Update sends a new line to the bar
func (b Bar) Update(s string) {
	trim := len(s)
	if trim > columnCount {
		trim = columnCount
	}

	b.stdin.Write([]byte(fmt.Sprintf("%v\n", s[:trim])))
}

// Scroll Sends text scrolling right to left across the bar
func (b Bar) Scroll(s string) {

	// Fade in and scroll left
	sleepInterval := 50 * time.Millisecond
	counter := columnCount
	length := len(s)

	for counter > 0 {
		text := strings.Repeat(" ", counter)

		trim := columnCount - counter

		if length > trim {
			text += s[:trim]
		} else {
			text += s
		}

		b.Update(fmt.Sprintf("%*v", counter, text))
		time.Sleep(sleepInterval)
		counter--
	}

	// Fade out from far left side
	counter = 0
	for counter <= length {
		b.Update(s[counter:])
		time.Sleep(sleepInterval)
		counter++
	}
}
