package main

import (
	"math"
	"regexp"
	"strconv"
)

type Position struct {
	x, y int
}

func FindManhattanDistance(instructions []string, relativeMovement bool) int {
	waypoint := createWaypoint(relativeMovement)
	ship := &Position{}

	for _, instruction := range instructions {
		action, value := extractActionAndValue(instruction)

		if actionIsCardinalDirection(action) {
			if relativeMovement {
				waypoint.moveForCardinalDirection(action, value)
			} else {
				ship.moveForCardinalDirection(action, value)
			}
			continue
		}

		if actionIsTurningDirection(action) {
			waypoint.rotate(action, value)
			continue
		}

		ship.moveForward(value, waypoint)
	}

	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}

func (p *Position) moveForCardinalDirection(direction string, value int) {
	switch direction {
	case "E":
		p.x += value
		return
	case "W":
		p.x -= value
		return
	case "N":
		p.y += value
		return
	case "S":
		p.y -= value
		return
	}
}

func (p *Position) rotate(direction string, angle int) {
	counterClockAngle := getCounterClockAngle(angle, direction)

	x := p.x
	y := p.y

	p.x = int(math.Round(float64(x)*cos(counterClockAngle) - float64(y)*sin(counterClockAngle)))
	p.y = int(math.Round(float64(x)*sin(counterClockAngle) + float64(y)*cos(counterClockAngle)))
}

func (p *Position) moveForward(value int, waypoint *Position) {
	p.x += waypoint.x * value
	p.y += waypoint.y * value
}

func createWaypoint(relativeMovement bool) *Position {
	if relativeMovement {
		return &Position{10, 1}
	} else {
		return &Position{1, 0}
	}
}

func getCounterClockAngle(angle int, direction string) int {
	if direction == "R" {
		return 360 - angle
	} else {
		return angle
	}
}

func sin(value int) float64 {
	return math.Sin(float64(value) * math.Pi / 180)
}

func cos(value int) float64 {
	return math.Cos(float64(value) * math.Pi / 180)
}

func extractActionAndValue(instruction string) (action string, value int) {
	match := actionAndValueMatch.FindAllStringSubmatch(instruction, 10)

	action = match[0][1]
	value, _ = strconv.Atoi(match[0][2])

	return action, value
}

func actionIsTurningDirection(action string) bool {
	return action == "L" || action == "R"
}

func actionIsCardinalDirection(action string) bool {
	return action == "E" ||
		action == "W" ||
		action == "S" ||
		action == "N"
}

var actionAndValueMatch = regexp.MustCompile(`([A-Z])(\d+)`)
