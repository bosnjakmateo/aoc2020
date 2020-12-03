package main

type Point struct {
	x int
	y int
}

func CalculateProductForNumberOfTreesEncounteredForGivenSlopes(landscape []string, slopes []Point) int {
	product := 1

	for _, slope := range slopes {
		product *= CalculateNumberOfTreesForGivenSlope(landscape, slope)
	}

	return product
}

func CalculateNumberOfTreesForGivenSlope(landscape []string, slope Point) (count int) {
	width, height := getLandscapeSize(landscape)

	currentPosition := Point{0, 0}

	for currentPosition.y < (height - slope.y) {
		xMoves := calculateMove(currentPosition.x, slope.x, width)
		yMoves := currentPosition.y + slope.y

		currentPosition.x = xMoves
		currentPosition.y = yMoves

		if currentPositionIsATree(landscape, currentPosition) {
			count++
		}
	}

	return count
}

func getLandscapeSize(landscape []string) (int, int) {
	return len(landscape[0]), len(landscape)
}

func calculateMove(currentPosition int, placesToMove int, maxAllowedPosition int) int {
	return (currentPosition + placesToMove) % maxAllowedPosition
}

func currentPositionIsATree(landscape []string, currentPosition Point) bool {
	return string(landscape[currentPosition.y][currentPosition.x]) == "#"
}
