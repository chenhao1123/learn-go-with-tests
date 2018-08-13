package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalword = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type CountdownOperationSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

//Countdown countdownStart .. go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalword)

}

func main() {
	sleep := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleep)
}
