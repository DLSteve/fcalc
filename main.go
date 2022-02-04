package main

import (
	"bufio"
	"fcalc/calculator"
	"fcalc/parser"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a fraction problem:")
	reader := bufio.NewReader(os.Stdin)
	problem, _ := reader.ReadString('\n')

	solution, err := SolveProblem(problem)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(solution)
}

func SolveProblem(problem string) (string, error) {
	prsr := parser.NewParser(problem)
	prbm, err := prsr.Parse()
	if err != nil {
		return "", err
	}

	switch prbm.Operator {
	case "+":
		return fmt.Sprintf("= %v", calculator.Add(prbm.Fraction1, prbm.Fraction2).ToString()), nil
	case "-":
		return fmt.Sprintf("= %v", calculator.Subtract(prbm.Fraction1, prbm.Fraction2).ToString()), nil
	case "*":
		return fmt.Sprintf("= %v", calculator.Multiply(prbm.Fraction1, prbm.Fraction2).ToString()), nil
	case "/":
		return fmt.Sprintf("= %v", calculator.Divide(prbm.Fraction1, prbm.Fraction2).ToString()), nil
	}

	return "", fmt.Errorf("invalid operator")
}
