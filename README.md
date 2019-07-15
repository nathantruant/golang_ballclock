SUMMARY
=======

This is a simulation of a rolling ball clock simulator, built for a job interview assignment.  
This is not an optimal solution using LCM or permutations, but rather simulates the function of a real rolling ball clock with the following description: 

Every minute, the least recently used ball is removed from the queue of balls at the bottom of the clock, elevated, then deposited on the minute indicator track, which is able to hold four balls. When a fifth ball rolls on to the minute indicator track, its weight causes the track to tilt. The four balls already on the track run back down to join the queue of balls waiting at the bottom in reverse order of their original addition to the minutes track. The fifth ball, which caused the tilt,rolls on down to the five-minute indicator track. This track holds eleven balls. The twelfth ball carried over from the minutes causes the five-minute track to tilt, returning the eleven balls to the queue, again in reverse order of their addition. The twelfth ball rolls down to the hour indicator. The hour indicator also holds eleven balls, but has one extra fixed ball which is always present so that counting the balls in the hour indicator will yield an hour in the range one to twelve. The twelfth ball carried over from the five-minute indicator causes the hour indicator to tilt, returning the eleven free balls to the queue, in reverse order, before the twelfth ball itself also returns to the queue.

RUNNING BALL CLOCK SIMULATOR
=================

The clock runs with command line arguments, supporting two possible modes: 

1. Takes a single argument specifying the number of balls and reports the number of balls given (between  27 and  127) in the input and the number of days (24-hour periods) which elapse before the clock returns to its initial ordering.

Usage:

    go run main.go 35

Output:

    35 balls cycle after 12 days

2. Takes two parameters, the number of balls and the number of minutes to run for.  If the number of minutes is specified, the clock will run to the number of minutes and report the state of the tracks at that point in a JSON format. The clock starts from 12:00 (AKA: 11 balls in the hour slot with an assumed  static slot  representing  1 for a standard twelve hour clock).

Usage:

    go run main.go 27 125

Output:

    Clock state after 125 minutes:
    {"Min":[],"FiveMin":[20],"Hour":[3],"Main":[6,5,19,24,23,1,21,26,12,25,18,17,4,8,27,15,2,7,16,14,13,22,9,10,11]}


RUNNING UNIT TESTS
=================

If you'd like to run the unit tests, you can run the following from
src/ballclock:

	go test ./...
