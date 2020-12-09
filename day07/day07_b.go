package main

func CalculateNumberOfIndividualBagsInsideTheShinyGoldBag(rawStorages []string) (sum int) {
	storage := parseData(rawStorages)

	return calculateNumberOfBags("shiny gold bag", storage)
}

func calculateNumberOfBags(name string, storage Storage) int {
	sum := 0

	for bag := range storage.bags[name] {
		sum += bag.capacity*calculateNumberOfBags(bag.name, storage) + bag.capacity
	}

	return sum
}
