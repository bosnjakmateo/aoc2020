package main

func CalculateExpenseOfTwoEntries(expenses []int) int {
	for _, e1 := range expenses {
		for _, e2 := range expenses {
			if e1+e2 == 2020 {
				return e1 * e2
			}
		}
	}

	return -1
}

func CalculateExpenseOfThreeEntries(expenses []int) int {
	for _, e1 := range expenses {
		for _, e2 := range expenses {
			for _, e3 := range expenses {
				if e1+e2+e3 == 2020 {
					return e1 * e2 * e3
				}
			}
		}
	}

	return -1
}
