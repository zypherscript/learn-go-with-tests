package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// realSleeper := &RealSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}

type RealSleeper struct {
}

func (r *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
