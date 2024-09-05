package main

import "strings"

func simplifyPath(path string) string {
	var stack []string

	for _, str := range strings.Split(path, "/") {
		switch {
		case str == "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case str != "" && str != ".":
			stack = append(stack, str)
		}
	}

	return "/" + strings.Join(stack, "/")
}
