package main

type Form struct {
	answers []string
}

func CalculateSumOfSameAnswersByAnyone(rawForms []string) (sum int) {
	forms := parseRawForms(rawForms)

	for _, form := range forms {
		answers := map[string]bool{}

		for _, answer := range form.answers {
			for _, a := range answer {
				answers[string(a)] = true
			}
		}

		sum += len(answers)
	}

	return sum
}

func CalculateSumOfSameAnswersByEveryone(rawForms []string) (sum int) {
	forms := parseRawForms(rawForms)

	for _, form := range forms {
		answers := map[string]int{}

		for _, answer := range form.answers {
			for _, a := range answer {
				answers[string(a)] = answers[string(a)] + 1
			}
		}

		if form.wasAnsweredByOnePerson() {
			sum += len(answers)
		} else {
			for _, a := range answers {
				if form.answerWasAnsweredByEveryone(a) {
					sum++
				}
			}
		}
	}

	return sum
}

func parseRawForms(rawForms []string) (forms []Form) {
	form := Form{}

	for _, rawForm := range rawForms {
		if rawForm == "" {
			forms = append(forms, form)
			form = Form{}
			continue
		}

		form.answers = append(form.answers, rawForm)
	}

	forms = append(forms, form)

	return forms
}

func (form *Form) wasAnsweredByOnePerson() bool {
	return len(form.answers) == 1
}

func (form *Form) answerWasAnsweredByEveryone(answerCount int) bool {
	return answerCount == len(form.answers)
}
