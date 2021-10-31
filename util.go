package main

import (
	"os"
	"regexp"
	"strings"
)

func fileExists(clippingsFilePath string) bool {
	info, err := os.Stat(clippingsFilePath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func findPatternMatchesInString(regexString string, stringToSearch string) []string {
	pattern := regexp.MustCompile(regexString)
	matches := pattern.FindAllStringSubmatch(stringToSearch, -1)
	matchGroups := make([]string, 0)

	for _, match := range matches {
		for i := 1; i < len(match); i++ {
			if !contains(matchGroups, match[i]) {
				matchGroups = append(matchGroups, strings.TrimSpace(match[i]))
			}
		}
	}
	return matchGroups
}

func contains(aSlice []string, elementToSearch string) bool {
	for _, element := range aSlice {
		if elementToSearch == element {
			return true
		}
	}
	return false
}

func split(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '-' || r == ','
}