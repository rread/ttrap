package main

import (
	"flag"
	"log"
	"time"
)

type Trigger struct {
	sound string
}

func (trigger *Trigger) play() {
	err := play(trigger.sound)
	if err != nil {
		log.Fatal(err)
	}
}

func (trigger *Trigger) repeat(v time.Duration, count int, stop *time.Time) {
	i := 0
	baseTime := time.Now()
	for {
		baseTime = baseTime.Add(v)
		i++
		if stop != nil && time.Now().After(*stop) {
			break
		}
		log.Printf("click %d", i)
		before := time.Now()
		trigger.play()
		if count > 0 && i >= count {
			break
		}
		time.Sleep(baseTime.Sub(time.Now()))
		since := time.Since(before)
		if since > v {
			log.Println("Interval duration is too short for shooting time.")
		}
	}
}

func main() {
	var (
		test     bool
		at       string
		delay    time.Duration
		interval time.Duration
		max      time.Duration
		end      string
		count    int
	)

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	trigger := &Trigger{}

	flag.StringVar(&at, "at", "", "Time to send trap")
	flag.DurationVar(&delay, "delay", 0, "Initial delay before starting to shoot.")
	flag.DurationVar(&interval, "interval", 0, "Repeat trigger every interval.")
	flag.DurationVar(&max, "for", 0, "For how long.")
	flag.StringVar(&end, "stop", "", "Time to stop interval")
	flag.IntVar(&count, "count", 0, "Number of images to take")
	flag.BoolVar(&test, "test", false, "Play test sound (for testing only)")
	flag.Parse()

	trigger.sound = "trigger-tone.wav"
	if test {
		trigger.sound = "test-tone.wav"
	}

	if at != "" {
		t, err := time.Parse("2006-01-02 15:04:05 MST", at)
		if err != nil {
			log.Fatal(err)
		}
		delay = t.Sub(time.Now())
	}

	var stop *time.Time
	if end != "" {
		var err error
		t, err := time.Parse("2006-01-02 15:04:05 MST", end)
		if err != nil {
			log.Fatal(err)
		}
		stop = &t
	}

	if max > 0 {
		if end != "" {
			log.Fatal("Can't set 'end' and 'max' at the same time.")
		}
		t := time.Now().Add(max)
		stop = &t
	}

	if interval == 0 {
		count = 1
	}

	if delay > 0 {
		log.Printf("delay start by %v\n", delay)
		time.Sleep(delay)
	}

	trigger.repeat(interval, count, stop)
}
