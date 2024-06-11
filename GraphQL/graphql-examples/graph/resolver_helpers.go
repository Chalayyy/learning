package graph

import (
	"fmt"
	"strconv"
)

var count int
func Increment() int {
	count++
	return count
}

func Age(asOf string, birthYear int) (int, error) {
	yearStr := asOf[0:4]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, fmt.Errorf("asOf argument must begin with a four-digit year")
	}
	return year - birthYear, nil
}

type user struct {
	name      string
	birthYear int
	friendIDs []int
	dogIDs    []int
}

type dog struct {
	name      string
	friendIDs []int
}

var users = []user{
	{name: "Elias", birthYear: 1997, friendIDs: []int{2, 1, 0}},
	{name: "Alicia", birthYear: 1960, friendIDs: []int{2}},
	{name: "Evan", birthYear: 1986, friendIDs: []int{0, 1}},
	{name: "Nat", birthYear: 1992, friendIDs: []int{1}, dogIDs: []int{0}},
}

var dogs = []dog{
	{name: "Spoonkin", friendIDs: []int{3}},
}