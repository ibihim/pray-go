package main

import (
	"fmt"

	"github.com/ibihim/pray-go/pkg/cmd"
)

func main() {
	if err := cmd.PrayerCommand().Execute(); err != nil {
		fmt.Println(err)
	}
}
