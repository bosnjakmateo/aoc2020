package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type Bus struct {
	id, index int
}

func FindEarliestBusIdMultipliedByWaitingTime(data []string) int {
	departure, busIds := extractDepartureAndBusIds(data)

	waitingTime := math.MaxInt64
	idealBusId := -1

	for _, busId := range busIds {
		nearestDeparture := calculateNearestDeparture(departure, busId)

		if nearestDeparture < departure {
			nearestDeparture += busId
		}

		if nearestDeparture-departure < waitingTime {
			waitingTime = nearestDeparture - departure
			idealBusId = busId
		}
	}

	return waitingTime * idealBusId
}

func FindPerfectBusDepartures(data []string) int {
	var n []*big.Int
	var a []*big.Int

	buses := extractBusIdsWithIdealDepartureTime(data)

	for _, bus := range buses {
		n = append(n, big.NewInt(int64(bus.id-bus.index)))
		a = append(a, big.NewInt(int64(bus.id)))
	}

	solution, _ := crt(n, a)

	return int(solution.Int64())
}

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(big.NewInt(1)) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func calculateNearestDeparture(departure, busId int) int {
	departure = departure + busId/2
	departure = departure - (departure % busId)

	return departure
}

func extractDepartureAndBusIds(data []string) (departure int, busIds []int) {
	departure, _ = strconv.Atoi(data[0])

	ids := strings.Split(data[1], ",")

	for _, id := range ids {
		if id != "x" {
			intId, _ := strconv.Atoi(id)
			busIds = append(busIds, intId)
		}
	}

	return departure, busIds
}

func extractBusIdsWithIdealDepartureTime(data []string) (busIds []Bus) {
	busIds = []Bus{}
	ids := strings.Split(data[1], ",")

	for i, id := range ids {
		if id != "x" {
			intId, _ := strconv.Atoi(id)
			busIds = append(busIds, Bus{id: intId, index: i})
		}
	}

	return busIds
}
