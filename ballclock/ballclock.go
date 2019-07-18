// Package ballclock is a rolling ball clock simulator
package ballclock

import (
	"encoding/json"
	"fmt"
)

const (
	minutesInDay         = 1440
	minuteTrackSlots     = 4
	fiveMinuteTrackSlots = 11
	hourTrackSlots       = 11
	// MaxBallCount is the maximum allowed balls
	MaxBallCount = 127
	// MinBallCount is the minimum allowed balls
	MinBallCount = 27
)

var defaultState string

// BallClock object
type BallClock struct {
	minTrack     []int
	fiveMinTrack []int
	hourTrack    []int
	ballQueue    []int
	ballCount    int
}

// New creates a new instance of the ball clock simulator object with the given number of balls
// And sets it to the default time of 12:00
func New(ballCount int) (*BallClock, error) {
	if ballCount < MinBallCount || ballCount > MaxBallCount {
		return nil, fmt.Errorf("Invalid argument: %d. Ball count must be from %d to %d", ballCount, MinBallCount, MaxBallCount)
	}

	ballClock := &BallClock{
		minTrack:     make([]int, 0, minuteTrackSlots),
		fiveMinTrack: make([]int, 0, fiveMinuteTrackSlots),
		hourTrack:    make([]int, 0, hourTrackSlots),
		ballQueue:    make([]int, 0, ballCount),
		ballCount:    ballCount,
	}

	ballClock.load()

	jsonString, err := ballClock.ToJSONString()
	defaultState = jsonString

	return ballClock, err
}

func (clock *BallClock) load() {
	for i := 0; i < clock.ballCount; i++ {
		clock.ballQueue = append(clock.ballQueue, i+1)
	}
}

// IsDefaultState checks the hour and queue to see if they are in the default 12:00 state
func (clock *BallClock) IsDefaultState() bool {
	if len(clock.ballQueue) == clock.ballCount {
		for i := 0; i < clock.ballCount; i++ {
			if clock.ballQueue[i] != i+1 {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// ToJSONString returns the clock state in a JSON format
func (clock *BallClock) ToJSONString() (string, error) {
	state := struct {
		Min     []int
		FiveMin []int
		Hour    []int
		Main    []int
	}{
		clock.minTrack,
		clock.fiveMinTrack,
		clock.hourTrack,
		clock.ballQueue,
	}

	jsonState, err := json.Marshal(state)

	return string(jsonState), err
}

// CalculateDaysUntilDefault gets the number of days until the balls return to the default state
func (clock *BallClock) CalculateDaysUntilDefault() int {
	days := 1
	for {
		for m := 0; m < minutesInDay; m++ {
			clock.Tick()

			if m == minutesInDay-1 && days >= 12 {
				if clock.IsDefaultState() {
					return days
				}
			}
		}
		days++
	}
}

// Tick progressess the clock by  1 minute, handling dumping of tracks when they are filled
func (clock *BallClock) Tick() {
	ball := clock.ballQueue[0]
	clock.ballQueue = clock.ballQueue[1:]

	if !runBall(&clock.minTrack, &clock.ballQueue, ball, minuteTrackSlots) {
		if !runBall(&clock.fiveMinTrack, &clock.ballQueue, ball, fiveMinuteTrackSlots) {
			if !runBall(&clock.hourTrack, &clock.ballQueue, ball, hourTrackSlots) {
				clock.ballQueue = append(clock.ballQueue, ball)
			}
		}
	}
}

func runBall(track *[]int, ballQueue *[]int, ball int, slots int) bool {
	roomOnTrack := len(*track) < cap(*track)
	if roomOnTrack {
		*track = append(*track, ball)
	} else {
		reverse(*track)
		*ballQueue = append(*ballQueue, *track...)
		*track = nil
		*track = make([]int, 0, slots)
	}

	return roomOnTrack
}

func reverse(slice []int) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}
