package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValid(tokens map[string]string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, field := range requiredFields {
		_, ok := tokens[field]
		if !ok {
			// fmt.Printf("%s invalid. missing %s\n", tokens, field)
			return false
		}
	}

	birthYear, err := strconv.Atoi(tokens["byr"])
	if err != nil || !(birthYear >= 1920 && birthYear <= 2002) {
		// fmt.Printf("%s invalid. bad byr %s\n", tokens, tokens["byr"])
		return false
	}

	issueYear, err := strconv.Atoi(tokens["iyr"])
	if err != nil || !(issueYear >= 2010 && issueYear <= 2020) {
		// fmt.Printf("%s invalid. bad iyr %s\n", tokens, tokens["iyr"])
		return false
	}

	expirationYear, err := strconv.Atoi(tokens["eyr"])
	if err != nil || !(expirationYear >= 2020 && expirationYear <= 2030) {
		// fmt.Printf("%s invalid. bad eyr %s\n", tokens, tokens["eyr"])
		return false
	}

	sHeightUnit := tokens["hgt"]
	sHeight, unit := sHeightUnit[:len(sHeightUnit)-2], sHeightUnit[len(sHeightUnit)-2:]
	height, err := strconv.Atoi(sHeight)
	if err != nil {
		// fmt.Printf("%s invalid. bad hgt %s\n", tokens, tokens["hgt"])
		return false
	}

	if unit == "cm" {
		if !(height >= 150 && height <= 193) {
			// fmt.Printf("%s invalid. bad hgt %s\n", tokens, tokens["hgt"])
			return false
		}
	} else if unit == "in" {
		if !(height >= 59 && height <= 76) {
			// fmt.Printf("%s invalid. bad hgt %s\n", tokens, tokens["hgt"])
			return false
		}
	} else {
		return false
	}

	validHairColor := regexp.MustCompile(`^\#[a-f0-9]{6}$`)
	if !validHairColor.MatchString(tokens["hcl"]) {
		// fmt.Printf("%s invalid. bad hcl %s\n", tokens, tokens["hcl"])
		return false
	}

	eyeColor := tokens["ecl"]
	if !(eyeColor == "amb" ||
		eyeColor == "blu" ||
		eyeColor == "brn" ||
		eyeColor == "gry" ||
		eyeColor == "grn" ||
		eyeColor == "hzl" ||
		eyeColor == "oth") {
		// fmt.Printf("%s invalid. bad ecl %s\n", tokens, tokens["ecl"])
		return false
	}

	validPassport := regexp.MustCompile(`^[0-9]{9}$`)
	if !validPassport.MatchString(tokens["pid"]) {
		// fmt.Printf("%s invalid. bad pid %s\n", tokens, tokens["pid"])
		return false
	}

	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokens := make(map[string]string)
	validCount, invalidCount := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if isValid(tokens) {
				validCount++
			} else {
				invalidCount++
			}
			tokens = make(map[string]string)
			continue
		}

		fields := strings.Split(t, " ")
		for _, field := range fields {
			token := strings.Split(field, ":")
			key, value := token[0], token[1]
			tokens[key] = value
		}
	}
	if isValid(tokens) {
		validCount++
	} else {
		invalidCount++
	}
	fmt.Printf("found %d valid passports, %d invalid passports\n", validCount, invalidCount)

}
