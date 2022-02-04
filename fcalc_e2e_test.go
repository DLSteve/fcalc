package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFractionCalc_Addition_MixedFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 + 3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 4_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Addition_PureFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 + 3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 1_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Addition_ImproperFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 + 34/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 9", solution, "the answer is not correct")
}

func TestFractionCalc_Addition_NegativeFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 + -3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -3_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Subtraction_MixedFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 - 3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -3_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Subtraction_PureFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 - 3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Subtraction_ImproperFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 - 34/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -8", solution, "the answer is not correct")
}

func TestFractionCalc_Subtraction_NegativeFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 - -3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 4_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Multiplication_MixedFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 * 3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 1_7/8", solution, "the answer is not correct")
}

func TestFractionCalc_Multiplication_PureFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 * 3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 3/8", solution, "the answer is not correct")
}

func TestFractionCalc_Multiplication_ImproperFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 * 34/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 4_1/4", solution, "the answer is not correct")
}

func TestFractionCalc_Multiplication_NegativeFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 * -3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -1_7/8", solution, "the answer is not correct")
}

func TestFractionCalc_Division_MixedFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 / 3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 2/15", solution, "the answer is not correct")
}

func TestFractionCalc_Division_PureFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 / 3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 2/3", solution, "the answer is not correct")
}

func TestFractionCalc_Division_ImproperFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 / 34/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= 1/17", solution, "the answer is not correct")
}

func TestFractionCalc_Division_NegativeFractions(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 / -3_3/4"
	solution, err := SolveProblem(problem)

	a.Nil(err, "error while parsing problem")
	a.Equal("= -2/15", solution, "the answer is not correct")
}

func TestFractionCalc_ParserError_NoQuery(t *testing.T) {
	a := assert.New(t)

	problem := "1/2 / -3_3/4"
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting opening '?'", err.Error(), "incorrect error thrown")
}

func TestFractionCalc_ParserError_NoOperator(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2  -3_3/4"
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting one of the following operators (+, -, *, /)", err.Error(), "incorrect error thrown")
}

func TestFractionCalc_ParserError_InvalidOperator(t *testing.T) {
	a := assert.New(t)

	problem := "? 1/2 = -3_3/4"
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting one of the following operators (+, -, *, /)", err.Error(), "incorrect error thrown")
}

func TestFractionCalc_ParserError_InvalidFraction(t *testing.T) {
	a := assert.New(t)

	problem := "? 34 1/2 + 3_3/4"
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting a regular or mixed fraction", err.Error(), "incorrect error thrown")
}

func TestFractionCalc_ParserError_InvalidInput(t *testing.T) {
	a := assert.New(t)

	problem := "? 34g1/2 + 3_3/4"
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting a regular or mixed fraction", err.Error(), "incorrect error thrown")
}

func TestFractionCalc_ParserError_MissingFraction(t *testing.T) {
	a := assert.New(t)

	problem := "? 34_1/2 + "
	_, err := SolveProblem(problem)

	a.NotNil(err, "no error thrown")
	a.Equal("expecting a regular or mixed fraction", err.Error(), "incorrect error thrown")
}
