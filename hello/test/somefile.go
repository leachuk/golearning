package test

import "strings"

func DebugThis(s string) []string {
	splitStrings := strings.Split(s, "/")
	return splitStrings
}