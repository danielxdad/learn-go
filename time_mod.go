package main

import (
	"fmt"
	"time"
)

func time_module() {
	now := time.Now()

	// fmt.Println(time.RFC3339)
	// fmt.Println(time.RFC3339Nano)

	dur, err := time.ParseDuration("1h10m30s")
	if err != nil {
		panic(err)
	}

	location, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		panic(err)
	}

	fmt.Println(now)
	fmt.Println(now.Add(dur))
	fmt.Println(now.UTC())
	fmt.Println(now.In(location))
	fmt.Println(now.Unix())
	fmt.Println(now.Zone())
}
