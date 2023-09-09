package utils

import "regexp"

func RemoveFromArrayByRegex(array []string, pattern *regexp.Regexp) []string {
	for i, msg := range array {
		array[i] = pattern.ReplaceAllString(msg, "")
	}
	return array
}
