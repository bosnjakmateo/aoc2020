package main

import "sort"

type Position struct{ start, end int }

func CalculateHighestSeatId(boardingPasses []string) int {
	seatsIds := calculateSeatsIds(boardingPasses)

	return seatsIds[len(seatsIds)-1]
}

func CalculateYourSeatId(boardingPasses []string) int {
	seatsIds := calculateSeatsIds(boardingPasses)

	for i := 1; i < len(seatsIds)-1; i++ {
		if seatsIds[i+1]-seatsIds[i] == 2 {
			return seatsIds[i] + 1
		}
	}

	return -1
}

func calculateSeatsIds(boardingPasses []string) (seatsIds []int) {
	for _, pass := range boardingPasses {
		rowPosition := Position{0, 127}
		columnPosition := Position{0, 7}

		rows := pass[:7]
		columns := pass[7:]

		for _, row := range rows {
			rowPosition.moveToTheNewHalf(string(row))
		}

		for _, column := range columns {
			columnPosition.moveToTheNewHalf(string(column))
		}

		seatsIds = append(seatsIds, (rowPosition.start*8)+columnPosition.start)
	}

	sort.Ints(seatsIds)

	return seatsIds
}

func (position *Position) moveToTheNewHalf(direction string) {
	middle := position.calculateMiddle()

	switch direction {
	case "F", "L":
		position.end = middle
	case "B", "R":
		position.start = middle + 1
	}
}

func (position *Position) calculateMiddle() int {
	return (position.end + position.start) / 2
}
