// Package ballclock is a rolling ball clock simulator
package ballclock

import (
	"testing"
)

func TestNew(t *testing.T) {
	clock, err := New(27)

	if err != nil {
		t.Errorf("Failed to create clock")
	}

	if clock.ballCount != 27 {
		t.Errorf("Ball count wasn't assigned properly")
	}

	if len(clock.hourTrack) != cap(clock.hourTrack) {
		t.Errorf("Hour track wasn't loaded properly")
	}

	if len(clock.ballQueue) != 27-cap(clock.hourTrack) {
		t.Errorf("Ball queue wasn't loaded properly")
	}

	if len(clock.minTrack) != 0 && len(clock.fiveMinTrack) != 0 {
		t.Errorf("min or fiveMin tracks not loaded properly")
	}

	clock, err = New(-10)
	if clock != nil {
		t.Errorf("Minimum Balls not followed")
	}

	clock, err = New(300)
	if clock != nil {
		t.Errorf("Maxium Balls not followed")
	}

	if err == nil {
		t.Errorf("Did not set set error properly")
	}
}

func TestIsDefaultState(t *testing.T) {
	clock, err := New(27)
	if err != nil {
		//already  tested
	}
	if !clock.IsDefaultState() {
		t.Errorf("Clock didn't load into  default state properly")
	}

	clock.Tick()
	if clock.IsDefaultState() {
		t.Errorf("Clock should not be in default state after ticking")
	}
}

func TestToJSONString(t *testing.T) {
	clock, err := New(27)
	if err != nil {
		//already  tested
	}

	jsonString, err := clock.ToJSONString()
	if err != nil {
		t.Errorf(err.Error())
	}
	if jsonString != "{\"Min\":[],\"FiveMin\":[],\"Hour\":[1,2,3,4,5,6,7,8,9,10,11],\"Main\":[12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27]}" {
		t.Errorf("JSON not marshalled properly")
	}
}

func TestCalculateDaysUntilDefault(t *testing.T) {
	clock, err := New(30)
	if err != nil {
		//already  tested
	}

	expected := 15
	actual := clock.CalculateDaysUntilDefault()

	if actual != expected {
		t.Errorf("Expected: %v Actual:  %v", expected, actual)
	}

	clock, err = New(45)
	if err != nil {
		//already  tested
	}

	expected = 378
	actual = clock.CalculateDaysUntilDefault()

	if actual != expected {
		t.Errorf("Expected: %v Actual:  %v", expected, actual)
	}
}
