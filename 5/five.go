package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// I am a huge dummy
// I realised that the two sections (row and column) were easy to turn into binary
// I didn't realise that the Seat ID calculation (row multiplied by *8*, plus column) meant that the whole SeatReference could be converted to binary in one jump

func seatRefToId(seatReference string) int64 {
	r := strings.NewReplacer(`F`, `0`, `B`, `1`, `L`, `0`, `R`, `1`)
	s := r.Replace(seatReference)
	seatId, _ := strconv.ParseInt(s, 2, 0)
	return seatId
}

func main() {
	fmt.Println("Hello World") // important

	inputfile, _ := ioutil.ReadFile("input")
	seatReferences := strings.Split(string(inputfile), "\n")

	var highestSeat int64 = 0
	var lowestSeat int64 = 1023
	var runningTotal int64 = 0
	for _, seat := range seatReferences {
		seatId := seatRefToId(seat)
		runningTotal = runningTotal + seatId
		if seatId > highestSeat {
			highestSeat = seatId
		}
		if seatId < lowestSeat {
			lowestSeat = seatId
		}
	}

	fmt.Println("Part 1(Highest Seat ID)", highestSeat)

	// It also took me a while to _understand_ part two.
	// "some of the seats at the front and the back don't exist"
	// I started off statically summing from 0 to 1023 (2^10) and wondering why I was miles off
	// Only them did I start calculating using the lowest and highest number...

	var sum int64 = 0
	for i := lowestSeat; i <= highestSeat; i++ {
		sum += i
	}

	fmt.Println("Part 2(Missing Seat ID)", sum-runningTotal)
}
