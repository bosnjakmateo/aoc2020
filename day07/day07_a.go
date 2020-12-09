package main

func CalculateNumberOfBagsThatCanHoldAShinyGoldBag(rawStorages []string) (sum int) {
	storage := parseData(rawStorages)

	return calculateLongestPath(storage)
}

func calculateLongestPath(storage Storage) (max int) {
	for bagName := range storage.bags {
		if bagContainsTheWantedBag(bagName, "shiny gold bag", storage) {
			max++
		}
	}

	return max - 1
}

func bagContainsTheWantedBag(currentBagName string, wantedBagName string, storage Storage) bool {
	if currentBagName == wantedBagName {
		return true
	}

	for bag := range storage.bags[currentBagName] {
		containsBag := bagContainsTheWantedBag(bag.name, wantedBagName, storage)

		if containsBag {
			return true
		}
	}

	return false
}
