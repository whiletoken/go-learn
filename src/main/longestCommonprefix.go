package main

import (
	"fmt"
	"strings"
)

// no := 14
func longestCommon(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs {
		for strings.Index(s, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func longestCommonTest() {

	strs := []string{"flower", "flow", "flows"}
	fmt.Println(longestCommon(strs))
}
