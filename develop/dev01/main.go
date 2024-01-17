package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetTime returns time from ntp lib
func GetTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, fmterr := fmt.Fprint(os.Stderr, err.Error())
		if fmterr != nil {
			log.Fatal(fmterr)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
