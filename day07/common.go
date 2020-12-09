package main

import (
	"regexp"
	"strconv"
)

type Bag struct {
	name     string
	capacity int
}

type Storage struct {
	bags map[string]map[Bag]bool
}

func (s *Storage) addNeighbour(bagName string, neighbour Bag) {
	s.bags[bagName][neighbour] = true
}

func (s *Storage) addNode(bagName string) {
	_, found := s.bags[bagName]

	if !found {
		s.bags[bagName] = map[Bag]bool{}
	}
}

func parseData(rawStorages []string) (storage Storage) {
	storage = Storage{make(map[string]map[Bag]bool)}

	for _, rawStorage := range rawStorages {
		bags := bagMatch.FindAllString(rawStorage, 20)

		parentNameBag := bagNameMatch.FindString(bags[0])

		storage.addNode(parentNameBag)

		for i := 1; i < len(bags); i++ {
			storage.addNode(bagNameMatch.FindString(bags[i]))

			storage.addNeighbour(parentNameBag, convertToBag(bags[i]))
		}
	}

	return storage
}

func convertToBag(bag string) Bag {
	name := bagNameMatch.FindString(bag)
	capacity, _ := strconv.Atoi(bagCapacityMatch.FindString(bag))

	return Bag{name, capacity}
}

var bagMatch = regexp.MustCompile(`(\d+\s)?([a-z]+\s[a-z]+\sbag)`)
var bagNameMatch = regexp.MustCompile(`[a-z]+\s[a-z]+\sbag`)
var bagCapacityMatch = regexp.MustCompile(`\d+`)
