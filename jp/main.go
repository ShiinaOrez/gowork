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
	{"か", "ka"},
	{"き", "ki"},
	{"く", "ku"},
	{"け", "ke"},
	{"こ", "ko"},
	{"さ", "sa"},
	{"し", "shi"},
	{"す", "su"},
	{"せ", "se"},
	{"そ", "so"},
	{"た", "ta"},
	{"ち", "chi"},
	{"つ", "tsu"},
	{"て", "te"},
	{"と", "to"},
	{"な", "na"},
	{"に", "ni"},
	{"ぬ", "nu"},
	{"ね", "ne"},
	{"の", "no"},
	{"は", "ha"},
	{"ひ", "hi"},
	{"ふ", "hu"},
	{"へ", "he"},
	{"ほ", "ho"},
	{"ま", "ma"},
	{"み", "mi"},
	{"む", "mu"},
	{"め", "me"},
	{"も", "mo"},
	{"ら", "ra"},
	{"り", "ri"},
	{"る", "ru"},
	{"れ", "re"},
	{"ろ", "ro"},
}

func init() {
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃        Welcome        ┃")
	fmt.Println("┃***********************┃")
	fmt.Println("┣━━━━━━━┳━━━━━━━┳━━━━━━━┫")
	fmt.Println("┃ COUNT ┃ BINGO ┃ WRONG ┃")
	fmt.Println("┃     0 ┃     0 ┃     \033[1;31m0\033[0m ┃")
	fmt.Println("┣━━━━━━━┻━━━━━━━┻━━━━━━━┛")
}

func printDashboard(count, bingo, wrong int) {
	fmt.Printf("┣━━━━━━━┳━━━━━━━┳━━━━━━━┓\n")
	fmt.Printf("┃ COUNT ┃ BINGO ┃ WRONG ┃\n")
	fmt.Printf("┃ %5d ┃ %5d ┃ \033[1;31m%5d\033[0m ┃\n", count, bingo, wrong)
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
		fmt.Printf("┃ (%03d) %s: ", count, p.Key)
		fmt.Scanf("%s\n", &v)
		if v == p.Value {
			bingo += 1
		} else {
			wrong += 1
			fmt.Printf("\033[1;31m┣━━━━━━WRONG: %3s ━━━━━━┫\033[0m\n", p.Value)
		}
		printDashboard(count, bingo, wrong)
	}
}
