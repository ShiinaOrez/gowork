// 五十音训练
package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Pair struct {
	Key   string
	Value string
}

var m []Pair = []Pair {
	{"あ", "a"},
	{"い", "i"},
	{"う", "u"},
	{"え", "e"},
	{"お", "o"},
}

func init() {
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃        Welcome        ┃")
	fmt.Println("┃***********************┃")
	fmt.Println("┣━━━━━━━┳━━━━━━━┳━━━━━━━┫")
	fmt.Println("┃ COUNT ┃ BINGO ┃ WRONG ┃")
	fmt.Println("┃     0 ┃     0 ┃     \033[0;31;40m0\033[0m ┃")
	fmt.Println("┣━━━━━━━┻━━━━━━━┻━━━━━━━┛")
}

func printDashboard(count, bingo, wrong int) {
	fmt.Printf("┣━━━━━━━┳━━━━━━━┳━━━━━━━┓\n")
	fmt.Printf("┃ COUNT ┃ BINGO ┃ WRONG ┃\n")
	fmt.Printf("┃ %5d ┃ %5d ┃ \033[0;31;40m%5d\033[0m ┃\n", count, bingo, wrong)
	fmt.Printf("┣━━━━━━━┻━━━━━━━┻━━━━━━━┛\n")
}

func main() {
	tot := len(m)
	count := 0
	bingo := 0
	wrong := 0
	for {
		rand.Seed(time.Now().Unix())
		randIndex := rand.Intn(tot)
		p := m[randIndex]
		count += 1

		var v string
		fmt.Printf("┃ %s: ", p.Key)
		fmt.Scanf("%s\n", &v)
		if v == p.Value {
			bingo += 1
		} else {
			wrong += 1
		}
		printDashboard(count, bingo, wrong)
	}
}