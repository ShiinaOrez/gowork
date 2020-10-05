package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"math/rand"
	"os"
	"time"
)

type Data struct {
	GrammaList grammaList `yaml:"GrammaList"`
}

type grammaList []Gramma

type Gramma struct {
	Text  string  `yaml:"Text"`
	Desc  string  `yaml:"Desc"`
	Tips  string  `yaml:"Tips"`
	From  int     `yaml:"From"`
}

var (
	src = Data{}
	occurMap = make(map[int]int)
)

// get gramma data from yaml
func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	yamlFile, err := os.Open(pwd+"/data.yaml")
	if err != nil {
		panic(err)
	}
    yamlDecoder := yaml.NewDecoder(yamlFile)
    err = yamlDecoder.Decode(&src)
    if err != nil {
    	panic(err)
	}
	rand.Seed(time.Now().Unix())
}

// main function
func main() {
	fmt.Println(len(src.GrammaList))

	round := 1
	randSeed := rand.Int63()

	// main endless for
	for {
		var (
			index int
		)
		{ // pre work
			rand.Seed(randSeed)
			randSeed = rand.Int63()
			index = int(randSeed) % len(src.GrammaList)
		}

		{ // deal
			fmt.Printf("==>> [%d] [%d] [Gramma]: %s\n", round, occurMap[index], src.GrammaList[index].Text)
		}

		{ // callback
			occurMap[index] += 1
			round += 1
		}
		inputSomething()
		{
			fmt.Printf("\tDesc: %s\n\tTips: %s\n", src.GrammaList[index].Desc, src.GrammaList[index].Tips)
		}
		inputSomething()
	}
}

func inputSomething() {
	var a string
	fmt.Scanf("%s", &a)
}
