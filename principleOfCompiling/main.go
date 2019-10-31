package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Tuple struct {
	Syn int
	Val interface{}
}

type SymbolMap map[string]int

func (m *SymbolMap) Insert(syms []string, codes []int) {
	length := len(syms)
	for i := 0; i < length; i++ {
		(*m)[syms[i]] = codes[i]
	}
	return
}

var (
	symbolMap  SymbolMap
	sourceCode string
	result     []Tuple
)

func init() {
	fmt.Println("[log]: Initing...")

	fmt.Printf("[log]: Start to init symbols...")
	symbolMap = make(map[string]int)
	symbolMap.Insert([]string{
		"begin", "if", "then", "while", "do", "end",
		"+", "-", "*", "/", ":", ":=", "<", "<>", "<=", ">", ">=", "=", ";",
		"(", ")", "#",
	}, []int{
		1, 2, 3, 4, 5, 6,
		13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 24, 25, 26,
		27, 28, 0,
	})
	fmt.Println("ok")

	fmt.Printf("[log]: Start to read the source code from in.source...")
	bytes, err := ioutil.ReadFile("./in.source")
	if err != nil {
		fmt.Println("[Error]: Read Source File Failed!\n\tReason: ", err.Error())
	}
	sourceCode = string(bytes)
	fmt.Println("ok")

	_ = os.Remove("./out.log")
	f, _ := os.OpenFile("./out.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	os.Stdout = f
}

func main() {
	sourceCode = strings.Replace(sourceCode, "\t", " ", -1)
	sourceCode = strings.Replace(sourceCode, "\n", "", -1)

	words := strings.Split(sourceCode, " ")
	for _, word := range words {
		if word == "" {
			continue
		} else {
			for word != "" {
				syn, res := FindSymbol(word)
				if res {
					tuple := Tuple{
						Syn: symbolMap[syn],
						Val: syn,
					}
					result = append(result, tuple)
					word = strings.Replace(word, syn, "", 1)
				} else {
					isNumber := false
					if IsNumber(word[0]) { // start with digit
						isNumber = true
					}
					if isNumber { // number
						index := 0
						for IsNumber(word[index]) {
							index++
							if index == len(word) {
								break
							}
						}
						number, _ := strconv.Atoi(word[:index])
						tuple := Tuple{
							Syn: 11,
							Val: number,
						}
						result = append(result, tuple)
						word = word[index:]
					} else if IsNotNumberOrLetter(word[0]) {
						fmt.Println("Syntex Error At: ", string(word[0]))
						return
					} else { // identifier
						index := 0
						for IsNumberOrLetter(word[index]) {
							index++
							if index == len(word) {
								break
							}
						}
						tuple := Tuple{
							Syn: 10,
							Val: word[:index],
						}
						result = append(result, tuple)
						word = word[index:]
					}
				}
			}
		}
	}

	for _, tuple := range result {
		fmt.Println("(", tuple.Syn, ", ", tuple.Val, ")")
	}
	return
}

func FindSymbol(word string) (string, bool) {
	if strings.HasPrefix(word, ":=") {
		return ":=", true
	}
	if strings.HasPrefix(word, "<=") {
		return "<=", true
	}
	if strings.HasPrefix(word, "<>") {
		return "<>", true
	}
	if strings.HasPrefix(word, ">=") {
		return ">=", true
	}

	for key, _ := range symbolMap {
		if strings.HasPrefix(word, key) {
			return key, true
		}
	}
	return "", false
}

func IsNumber(b byte) bool {
	return int(b) <= int('9') && int(b) >= int('0')
}

func IsLetter(b byte) bool {
	return int(b) <= int('z') && int(b) >= int('a')
}

func IsNumberOrLetter(b byte) bool {
	return IsNumber(b) || IsLetter(b)
}

func IsNotNumberOrLetter(b byte) bool {
	return !IsNumberOrLetter(b)
}
