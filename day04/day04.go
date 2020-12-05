package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

func CalculateNumberOfValidPassports(rawPassports []string) (count int) {
	passports := parseRawPassportData(rawPassports)

	for _, passport := range passports {
		if passport.allFieldsPresent() {
			count++
		}
	}

	return count
}

func CalculateNumberOfValidPassportsV2(rawPassports []string) (count int) {
	passports := parseRawPassportData(rawPassports)

	for _, passport := range passports {
		if passport.allFieldsPresent() && passport.allFieldsValid() {
			count++
		}
	}

	return count
}

func parseRawPassportData(rawPassports []string) (passports []Passport) {
	passport := Passport{}

	for _, rawPassport := range rawPassports {
		if rawPassport == "" {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		passport.fillData(rawPassport)
	}

	passports = append(passports, passport)

	return passports
}

func (p *Passport) fillData(rawData string) {
	if rawData == "" {
		return
	}

	for _, data := range strings.Split(rawData, " ") {
		name, value := extractNameAndValue(data)

		switch name {
		case "byr":
			p.byr, _ = strconv.Atoi(value)
		case "iyr":
			p.iyr, _ = strconv.Atoi(value)
		case "eyr":
			p.eyr, _ = strconv.Atoi(value)
		case "hgt":
			p.hgt = value
		case "hcl":
			p.hcl = value
		case "ecl":
			p.ecl = value
		case "pid":
			p.pid = value
		}
	}
}

func (p *Passport) allFieldsPresent() bool {
	if p.byr == 0 || p.iyr == 0 || p.eyr == 0 || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
		return false
	}

	return true
}

func (p *Passport) allFieldsValid() bool {
	return isValueInRange(p.byr, 1920, 2002) &&
		isValueInRange(p.iyr, 2010, 2020) &&
		isValueInRange(p.eyr, 2020, 2030) &&
		isHeightValid(p.hgt) &&
		hclCheck.MatchString(p.hcl) &&
		eclCheck.MatchString(p.ecl) &&
		pidCheck.MatchString(p.pid)
}

func extractNameAndValue(data string) (string, string) {
	splitData := strings.Split(data, ":")

	return splitData[0], splitData[1]
}

func isHeightValid(height string) bool {
	heightValue, _ := strconv.Atoi(height[:len(height)-2])

	switch height[len(height)-2:] {
	case "cm":
		if !isValueInRange(heightValue, 150, 193) {
			return false
		}
	case "in":
		if !isValueInRange(heightValue, 59, 76) {
			return false
		}
	default:
		return false
	}

	return true
}

func isValueInRange(value, min, max int) bool {
	return value >= min && value <= max
}

var hclCheck = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eclCheck = regexp.MustCompile(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`)
var pidCheck = regexp.MustCompile(`^\d{9}$`)
