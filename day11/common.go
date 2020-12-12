package main

import "strings"

func getSeatValue(seats [][]string, x, y int) string {
	maxX := len(seats[0])
	maxY := len(seats)

	if x < 0 || x >= maxX {
		return "0"
	}

	if y < 0 || y >= maxY {
		return "0"
	}

	return seats[y][x]
}

func parseSeats(rawSeats []string) [][]string {
	seats := make([][]string, len(rawSeats))
	for i := range seats {
		seats[i] = make([]string, len(rawSeats[0]))
	}

	for i, rawSeat := range rawSeats {
		split := strings.Split(rawSeat, "")

		for j, s := range split {
			seats[i][j] = s
		}
	}

	return seats
}

type Coordinates struct {
	x, y int
}

var adjacentSeatsCoordinates = []Coordinates{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}
