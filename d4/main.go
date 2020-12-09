package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("values")
	pps := fileToPassports(string(file))
	requiredAttributes := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	validators := map[string]func(string) bool{
		"byr": func(s string) bool {
			val, _ := strconv.Atoi(s)
			return val >= 1920 && val <= 2002
		},
		"iyr": func(s string) bool {
			val, _ := strconv.Atoi(s)
			return val >= 2010 && val <= 2020
		},
		"eyr": func(s string) bool {
			val, _ := strconv.Atoi(s)
			return val >= 2020 && val <= 2030
		},
		"hgt": func(s string) bool {
			if strings.Contains(s, "in") {
				s = strings.TrimSuffix(s, "in")
				val, _ := strconv.Atoi(s)
				return val >= 59 && val <= 76
			} else if strings.Contains(s, "cm") {
				s = strings.TrimSuffix(s, "cm")
				val, _ := strconv.Atoi(s)
				return val >= 150 && val <= 193
			}
			return false
		},
		"hcl": func(s string) bool {
			matched, _ := regexp.Match("#[0-9a-f]{6}", []byte(s))
			return matched
		},
		"ecl": func(s string) bool {
			allowed := "amb,blu,brn,gry,grn,hzl,oth,"
			return strings.Contains(allowed, s+",")
		},
		"pid": func(s string) bool {
			_, err := strconv.Atoi(s)
			return err == nil && len(s) == 9
		},
	}

	counter := 0
	for _, pp := range pps {
		if isPPValid(pp, requiredAttributes, validators) {
			counter++
		}
	}
	fmt.Printf("%v\n", counter)
}

func isPPValid(passport Passport, requiredAttrs []string, validators map[string]func(string) bool) bool {
	for _, attr := range requiredAttrs {
		if _, ok := passport.attributes[attr]; !ok {
			return false
		}
		if validator, ok := validators[attr]; ok {
			attr := passport.attributes[attr]
			if !validator(attr) {
				return false
			}
		}
	}
	return true
}

func fileToPassports(file string) []Passport {
	ppStrings := strings.Split(file, "\n\n")
	passports := make([]Passport, len(ppStrings))
	for ppIndex, ppString := range ppStrings {
		pp := Passport{attributes: make(map[string]string)}
		ppString = strings.ReplaceAll(ppString, "\n", " ")

		ppAttributes := strings.Split(ppString, " ")
		for _, ppAttribute := range ppAttributes {
			splitted := strings.Split(strings.TrimSpace(ppAttribute), ":")
			pp.attributes[splitted[0]] = splitted[1]
		}
		passports[ppIndex] = pp
	}
	return passports
}

type Passport struct {
	attributes map[string]string
}
