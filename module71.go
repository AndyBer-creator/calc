package main

import (
	"bufio"
	"fmt"
	"os"
	"skill/module71/calc"
	"strconv"
)

func isValidOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Вводи число")
	firstString, _ := reader.ReadString('\n')
	firstNum, err := strconv.ParseFloat(firstString[:len(firstString)-1], 64)
	if err != nil {
		fmt.Println("Муть ввел")
		return
	}
	fmt.Println("Вводи знак арифметический")
	operator, _ := reader.ReadString('\n')
	operator = operator[:len(operator)-1]
	if !isValidOperator(operator) {
		//log.Fatal(fmt.Errorf("сам то знаешь, что за знак?"))
		fmt.Println("Знак не тот")
		return
	}

	fmt.Println("Второе число введи")
	nextString, _ := reader.ReadString('\n')
	nextNum, err := strconv.ParseFloat(nextString[:len(nextString)-1], 64)
	if err != nil {
		fmt.Println("Муть ввел")
		return
	}
	calсInstance := calc.NewCalculator()
	total := calсInstance.Calculate(firstNum, nextNum, operator)
	fmt.Printf("Итого %.2f\n", total)
}
