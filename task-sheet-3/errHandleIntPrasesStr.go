package main

import (
	"fmt"
	"strconv"
)

type errs struct {
	msg   string
	input string
}

func (e *errs) Error() string {
	return fmt.Sprint("Error while prasing ", "'", e.input, "'", " : ", e.msg)
}

func intPrasesStr(nums int, input string) (int, error) {

	newNum, err := strconv.Atoi(input)

	if err != nil {
		return 0, &errs{input: input, msg: err.Error()}
	} else {
		return newNum, nil
	}

}

func main() {
	num := 123456
	// input := "12345623"
	input := "abcdef"
	newNum, err := intPrasesStr(num, input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newNum)
	}

}
