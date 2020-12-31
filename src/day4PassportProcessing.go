package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"./files"
)

var requiredKeys = [...]string{
	"byr", // (Birth Year)
	"iyr", // (Issue Year)
	"eyr", // (Expiration Year)
	"hgt", // (Height)
	"hcl", // (Hair Color)
	"ecl", // (Eye Color)
	"pid"} // (Passport ID)

var fieldValidation = map[string]func(string) bool{
	"byr": func(v string) bool {
		return checkValidRange(v, 1920, 2002)
	}, // (Birth Year)
	"iyr": func(v string) bool {
		return checkValidRange(v, 2010, 2020)
	}, // (Issue Year)
	"eyr": func(v string) bool {
		return checkValidRange(v, 2020, 2030)
	}, // (Expiration Year)
	"hgt": func(v string) bool {
		if strings.HasSuffix(v, "in") {
			heightStr := strings.ReplaceAll(v, "in", "")
			return checkValidRange(heightStr, 59, 76)
		} else if strings.HasSuffix(v, "cm") {
			heightStr := strings.ReplaceAll(v, "cm", "")
			return checkValidRange(heightStr, 150, 193)
		}
		return false
	}, // (Height)
	"hcl": func(v string) bool {
		return checkRegex("^#[0-9a-f]{6}$", v)
	}, // (Hair Color)
	"ecl": func(v string) bool {
		return checkRegex("^(amb|blu|brn|gry|grn|hzl|oth)$", v)
	}, // (Eye Color)
	"pid": func(v string) bool {
		return checkRegex("^[0-9]{9}$", v)
	}} // (Passport ID)

func day4() {
	fmt.Println("Day 4.")
	lines := files.ReadLines("inputs/day4.txt")
	passports := parsePassports(lines)
	fmt.Println("Total Passports", len(passports))

	day4Sol1(passports)
	day4Sol2(passports)

	// day3Sol2(board, steps)
}

func day4Sol1(passports []map[string]string) int {
	fmt.Println("Solution 1")
	sum := 0
	for _, pass := range passports {
		if isValid(pass) {
			sum += 1
		}
	}
	fmt.Println(sum)
	return sum

}

func day4Sol2(passports []map[string]string) int {
	fmt.Println("Solution 2")
	sum := 0
	for _, pass := range passports {
		if isValidEnhanced(pass) {
			sum += 1
		}
	}
	fmt.Println(sum)
	return sum

}

func isValid(pass map[string]string) bool {
	for _, key := range requiredKeys {
		_, ok := pass[key]
		if !ok {
			return false
		}
	}
	return true
}

func isValidEnhanced(pass map[string]string) bool {
	for _, key := range requiredKeys {
		val, ok := pass[key]
		if !ok {
			return false
		} else if !fieldValidation[key](val) {
			fmt.Println(key, val, "invalid")
			return false
		}
	}
	return true
}

func checkValidRange(value string, lLimit int, hLimit int) bool {
	n, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return lLimit <= n && n <= hLimit
}
func checkRegex(patt string, v string) bool {
	match, err := regexp.MatchString(patt, v)
	if err != nil {
		return false
	}
	return match
}

func parsePassports(lines []string) []map[string]string {
	var passports []map[string]string
	var pass = make(map[string]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			passports = append(passports, pass)
			pass = make(map[string]string)
		} else {
			lineToks := strings.Split(line, " ")
			fmt.Println(lineToks)
			for _, tok := range lineToks {
				entry := strings.Split(tok, ":")
				pass[entry[0]] = entry[1]
			}
		}
	}
	passports = append(passports, pass)

	return passports

}
