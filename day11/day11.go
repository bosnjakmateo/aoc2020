package main

import . "aoc2019/util"

func FindNumberOfOccupiedSeatsPartOne(rawSeats []string) (count int) {
	return findNumberOfOccupiedSeats(rawSeats, 4, countAdjacentTakenSeatsRightNextToSeat)
}

func FindNumberOfOccupiedSeatsPartTwo(rawSeats []string) (count int) {
	return findNumberOfOccupiedSeats(rawSeats, 5, countAdjacentTakenSeatsWithExpandedVision)
}

func findNumberOfOccupiedSeats(rawSeats []string, occupiedSeatTolerance int, countAdjacentTakenSeats func(seats [][]string, x, y int) (count int)) (count int) {
	seats := parseSeats(rawSeats)

	for {
		newSeats, changed := generateNewSeatsState(seats, occupiedSeatTolerance, countAdjacentTakenSeats)
		seats = newSeats

		if !changed {
			break
		}
	}

	for i := range seats {
		for _, seat := range seats[i] {
			if seat == "#" {
				count++
			}
		}
	}

	return count
}

func generateNewSeatsState(seats [][]string, occupiedSeatTolerance int, countAdjacentTakenSeats func(seats [][]string, x, y int) (count int)) (newSeats [][]string, changed bool) {
	newSeats = Create2DStringArray(len(seats), ".")

	for y := range seats {
		for x := range seats[y] {
			if seats[y][x] == "." {
				continue
			}

			takenSeats := countAdjacentTakenSeats(seats, x, y)

			if seats[y][x] == "L" && takenSeats == 0 {
				newSeats[y][x] = "#"
				changed = true
				continue
			}

			if seats[y][x] == "#" && takenSeats >= occupiedSeatTolerance {
				newSeats[y][x] = "L"
				changed = true
				continue
			}

			newSeats[y][x] = seats[y][x]
		}
	}

	return newSeats, changed
}

func countAdjacentTakenSeatsWithExpandedVision(seats [][]string, x, y int) (count int) {
	for _, coordinate := range adjacentSeatsCoordinates {
		newX := x
		newY := y

		for {
			newX += coordinate.x
			newY += coordinate.y

			seatValue := getSeatValue(seats, newX, newY)

			if seatValue == "0" || seatValue == "L" {
				break
			}

			if seatValue == "." {
				continue
			}

			if seatValue == "#" {
				count++
				break
			}
		}
	}

	return count
}

func countAdjacentTakenSeatsRightNextToSeat(seats [][]string, x, y int) (count int) {
	for _, coordinate := range adjacentSeatsCoordinates {
		if getSeatValue(seats, x+coordinate.x, y+coordinate.y) == "#" {
			count++
		}
	}

	return count
}
