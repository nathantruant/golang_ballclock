package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/nathantruant/golang_ballclock/ballclock"
)

func parseCommandLineArguments() {
	flag.Parse()

	if len(flag.Arg(2)) > 0 {
		fmt.Printf("ballclock %s: unknown command\n.  Run 'ballclock help' for usage", flag.Args())
		return
	}

	if strings.ToLower(flag.Arg(0)) == "-h" || strings.ToLower(flag.Arg(0)) == "help" {
		printHelp()
		return
	}

	handleArguments()
}

func handleArguments() {
	ballCount, duration := 0, 0
	hasCount, hasDuration := false, false

	if len(flag.Arg(0)) > 0 {
		b, err := strconv.Atoi(flag.Arg(0))
		hasCount = err == nil
		if hasCount {
			ballCount = b
		}
	}

	if len(flag.Arg(1)) > 0 {
		d, err := strconv.Atoi(flag.Arg(1))
		hasDuration = err == nil
		if hasDuration {
			duration = d
		}
	}

	validBallCount := ballCount >= ballclock.MinBallCount && ballCount <= ballclock.MaxBallCount
	if hasCount && hasDuration && validBallCount {
		if duration < 0 {
			fmt.Printf("Invalid argument. Duration must not be negative.")
		} else {
			calculateState(ballCount, duration)
		}
	} else if hasCount && validBallCount {
		calculateDuration(ballCount)
	} else {
		fmt.Println("Invalid argument. Ball input must be from 27 to 127.")
	}
}

func calculateState(ballCount, duration int) {
	clock, err := ballclock.New(ballCount)
	if err != nil {
		fmt.Printf("Unable to create ball clock: %s\n\n", err.Error())
		return
	}

	fmt.Printf("Calculating the state of a ball clock with %d balls after %d minutes...\n\n", ballCount, duration)

	for i := 0; i < duration; i++ {
		clock.Tick()
	}

	jsonString, err := clock.ToJSONString()
	if err != nil {
		fmt.Printf("Unable to get clock state: %s\n\n", err.Error())
		return
	}

	fmt.Printf("Clock state after %d minutes:\n", duration)
	fmt.Printf("%s\n", jsonString)
}

func calculateDuration(ballCount int) {
	clock, err := ballclock.New(ballCount)
	if err != nil {
		fmt.Printf("Unable to create ball clock: %s\n\n", err.Error())
		return
	}

	fmt.Printf("Calculating how long it will take %d balls to return  to their default state...\n\n", ballCount)
	fmt.Printf("%d balls cycle after %d days\n\n", ballCount, clock.CalculateDaysUntilDefault())
}

func printHelp() {
	fmt.Printf("This is a ball clock simulator.  There are two available modes:\n\n")
	fmt.Println("    1. Given a specified number of balls from 30 to 127, calculate the number of days")
	fmt.Println("       until the balls return to their initial state")
	fmt.Printf("\n\n        Usage:\n\n")
	fmt.Printf("            ballclock [num of balls]\n\n")
	fmt.Println("    2. Given a specified number of balls from 30 to 127 and a specified number of minutes,")
	fmt.Println("       return the state of the clock once the number of minutes would have passed")
	fmt.Printf("\n\n        Usage:\n\n")
	fmt.Printf("            ballclock [num of balls] [num of minutes]\n\n")
}

func main() {
	parseCommandLineArguments()
}
