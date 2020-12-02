package main

func CalculateExpenseOfTwoEntries(expenses []int) int {
	for i := range expenses {
		for j := range expenses {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}

	return -1
}

func CalculateExpenseOfThreeEntries(expenses []int) int {
	for i := range expenses {
		for j := range expenses {
			for k := range expenses {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}

	return -1
}
