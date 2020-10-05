package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &n)
	a := make([]string, n)
	for i:=0; i<n; i++ {
		var ai string
		fmt.Scanf("%s", &ai)
		a[i] = ai
	}

	for {
		find := false
		for _, ai := range a {
			if strings.Contains(s, ai) {
				find = true
				s = strings.Replace(s, ai, "", 1)
				break
			}
		}
		if !find {
			break
		}
	}
	fmt.Printf("%s", s)
}