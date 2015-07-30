package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var (
		at   string
		test bool
	)

	flag.StringVar(&at, "at", "", "Time to send trap")
	flag.BoolVar(&test, "test", false, "Play test sound (for testing only)")
	flag.Parse()

	t, err := time.Parse("2006-01-02 15:04:05 MST", at)
	if err != nil {
		log.Fatal(err)
	}

	sound := "trigger-tone.wav"
	if test {
		sound = "test-tone.wav"
	}
	log.Println("wait for it...")
	select {
	case _ = <-time.After(t.Sub(time.Now())):
		log.Println("click")
		err := play(sound)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("done")
	}

}
