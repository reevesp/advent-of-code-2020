package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*

https://adventofcode.com/2020/day/3

Get a list of coordinates on the "route down" to check for trees.

From the start point at (0,0), iterate over the tobbogan hops until it would be at the end of the course.

Then, for each co-ordinate in the array, look up that position (using Mod for the x co-ordinate) and see whether it is a tree or not,
 incrementing a counter as we go.

Then, return the number of tree collisions.
*/

// Global vars to be used within functions
var courseRows []string
var courseWidth int
var courseLength int

// this feels wrong
type coord struct {
	x int
	y int
}

// Given a pair of x and y increments, plot a course down the slope (and return an array/slice? of co-ordinates)
func plotTobboganCourse(xincrement int, yincrement int) []coord {
	course := make([]coord, 1)
	course[0] = coord{0, 0}
	for true {
		c := course[len(course)-1]
		if c.y+yincrement >= courseLength {
			// fmt.Println("Debug: Adding the following coord", c.x + xincrement, c.y + yincrement, "would take the path over the course length of", courseLength)
			break
		}
		course = append(course, coord{c.x + xincrement, c.y + yincrement})
	}
	return course
}

// Given a co-ordinate, return true of false whether there exists a tree
func treeExistsAtCoord(c coord) bool {
	r := courseRows[c.y]
	x := c.x % courseWidth
	h := string('#')
	// fmt.Println("Debug: lookup tree in row", r, "position", x)
	// why can't I compare this to a '#'
	if h == string(r[x]) {
		return true
	} else {
		return false
	}
}

func countCollisions(course []coord) int {
	numberOfTreeCollisions := 0
	for _, element := range course {
		if treeExistsAtCoord(element) {
			numberOfTreeCollisions = numberOfTreeCollisions + 1
		}
		// fmt.Println("Debug: Number of tree collisions", numberOfTreeCollisions)
	}
	return numberOfTreeCollisions
}

func main() {
	fmt.Println("Hello World") // important

	// read in the input file
	inputfile, _ := ioutil.ReadFile("input")
	// break the file into an array of elements
	courseRows = strings.Split(string(inputfile), "\n")

	courseLength = len(courseRows)
	courseWidth = len(courseRows[0])

	case1 := countCollisions(plotTobboganCourse(1, 1))
	case2 := countCollisions(plotTobboganCourse(3, 1))
	case3 := countCollisions(plotTobboganCourse(5, 1))
	case4 := countCollisions(plotTobboganCourse(7, 1))
	case5 := countCollisions(plotTobboganCourse(1, 2))

	// fmt.Println(case1, case2, case3, case4, case5)
	fmt.Println("Part 1: Number of tree collisions is:", case2)
	fmt.Println("Part 2: 5 courses multiplied together is:", case1*case2*case3*case4*case5)
}
