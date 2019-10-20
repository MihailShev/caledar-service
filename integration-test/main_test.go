package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("Waiting all services start")
	wait(10)
	//time.Sleep(10 * time.Second)
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"features"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func wait(duration uint) {
	ticker := time.NewTicker(1 * time.Second)

	for i := duration; i >= 1; i-- {
		if i == 1 {
			fmt.Printf(" %d\n", i)
		} else {
			fmt.Printf(" %d", i)
		}
		<-ticker.C
	}
}
