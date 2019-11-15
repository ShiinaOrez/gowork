package main

import (
	"log"
	_ "fmt"
)

func is_factor(tuples []Tuple) bool {
	if len(tuples) == 1 {
		if tuples[0].Syn == 10 || tuples[0].Syn == 11 {
			return true;
		} else {
			log.Println("Error: Item must be indentifier or number or (expression)!")
			return false
		}
	} else if len(tuples) == 0 {
		log.Println("Error: Item can not be EMPTY string!")
		return false
	} else {
		if tuples[0].Syn == 27 && tuples[len(tuples)-1].Syn == 28 {
			if res := is_expression(tuples[1:len(tuples)-1]); !res {
				log.Println("Error: Error with expression!")
				return false
			}
		} else {
			for _, tuple := range tuples {
				log.Println(tuple)
			}
			log.Println("Error: Expression in item must with `(` and `)` !")
			return false
		}
	}
	return true
}

func is_item(tuples []Tuple) bool {
	// log.Println(tuples)
	last_end := -1
	length := len(tuples)
	if length == 0 {
		log.Println("Error: Item can not be EMPTY string!")
		return false
	}
	if tuples[0].Syn == 27 {
		if res := is_factor(tuples); !res {
			log.Println("Error: Error with factor!")
			return false
		}
		return true
	}
	for index, tuple := range tuples {
		if tuple.Syn == 15 || tuple.Syn == 16 || index == length-1 {
			if index == length-1 {
				index += 1
			}
			if res := is_factor(tuples[last_end+1:index]); !res {
				log.Println("Error: Error with factor!")
				return false
			} else {
				last_end = index
			}
		}
	}
	return true
}

func is_expression(tuples []Tuple) bool {
	last_end := -1
	left_bracket := []int{}
	length := len(tuples)
	if length == 0 {
		log.Println("Error: Expression can not be EMPTY string!")
		return false
	}
	for index, tuple := range tuples {
		if tuple.Syn == 27 {
			left_bracket = append(left_bracket, index)
			// log.Println(left_bracket)
		} else if tuple.Syn == 28 {
			// log.Println(tuples[left_bracket[len(left_bracket)-1]:index+1])
			if len(left_bracket) == 0 {
				log.Println("Error: Lack of `(` !")
				return false
			}
			if res := is_item(tuples[left_bracket[len(left_bracket)-1]:index+1]); !res {
				log.Println("Error: Error with item!")
				return false
			}
			left_bracket = left_bracket[:len(left_bracket)-1]
			last_end = index
		} else if tuple.Syn == 13 || tuple.Syn == 14 || index == length-1 {
			if len(left_bracket) != 0 {
				continue
			}
			if index == length-1 {
				index += 1
				if len(left_bracket) != 0 {
					log.Println("Error: Lack of `)` !")
					return false
				}
			}
			if last_end == index-1 {
				last_end = index
				continue
			}
			if res := is_item(tuples[last_end+1:index]); !res {
				log.Println("Error: Error with item!")
				return false
			} else {
				last_end = index
			}
		}
	}
	return true
}

func is_assignment_statement(tuples []Tuple) bool {
	length := len(tuples)
	if length == 0 {
		log.Println("Error: Assignment statement can not be EMPTY string!")
		return false
	}
	if tuples[0].Syn != 10 {
		log.Println("Error: Assignment statement required indentifier!")
		return false
	}
	if length < 2 || tuples[1].Syn != 18 {
		log.Println("Error: Assignment statement required `:=` !")
		return false;
	}
	if res := is_expression(tuples[2:]); !res {
		log.Println("Error: Error with expression!")
		return false
	}
	return true
}

func is_statement(tuples []Tuple) bool {
	length := len(tuples)
	if length == 0 {
		log.Println("Error: Statement can not be EMPTY string!")
		return false
	}
	if res := is_assignment_statement(tuples); !res {
		log.Println("Error: Error with assignment statement!")
		return false
	}
	return true
}

func is_statement_string(tuples []Tuple) bool {
	length := len(tuples)
	if length == 0 {
		log.Println("Error: Statement string can not be EMPTY string!")
		return false
	}
	last_end := -1
	for index, tuple := range tuples {
		if tuple.Syn == 26 || index == length-1 {
			if index == length-1 {
				index += 1
			}
			if res := is_statement(tuples[last_end+1:index]); !res {
				log.Println("Error: Error with statement!")
				return false
			} else {
				last_end = index
			}
		}
	}
	return true
}

func main() {
	log.Println("Start: Analysising...")
	tuples := lexicalAnalysis()
	length := len(tuples)
	if tuples[0].Syn != 1 || tuples[length-1].Syn != 0 {
		log.Println("Error: Please start with `begin` and end with `end`!")
		return
	}
	if res := is_statement_string(tuples[1:length-1]); !res {
		log.Println("Error: Error with statement string!")
	} else {
		log.Println("Analysis Successful!")
	}
	return
}