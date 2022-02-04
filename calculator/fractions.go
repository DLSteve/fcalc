package calculator

import (
	"fmt"
	"math"
	"sort"
)

type MixedFraction struct {
	WholeNumber int
	Numerator   int
	Denominator int
}

func (m MixedFraction) ToString() string {
	if m.WholeNumber != 0 {

		// format negative numerator mixed fraction results
		if m.WholeNumber < 0 && m.Numerator < 0 {
			return fmt.Sprintf("%v_%v/%v", m.WholeNumber, InvertInt(m.Numerator), m.Denominator)
		}

		// format negative denominator mixed fraction results
		if m.WholeNumber < 0 && m.Denominator < 0 {
			return fmt.Sprintf("%v_%v/%v", m.WholeNumber, m.Numerator, InvertInt(m.Denominator))
		}

		// format whole number results
		if m.Numerator == 0 && m.Denominator == 1 {
			return fmt.Sprintf("%v", m.WholeNumber)
		}

		return fmt.Sprintf("%v_%v/%v", m.WholeNumber, m.Numerator, m.Denominator)
	}

	// format negative denominator fraction results
	if m.Denominator < 0 {
		return fmt.Sprintf("%v/%v", InvertInt(m.Numerator), InvertInt(m.Denominator))
	}

	return fmt.Sprintf("%v/%v", m.Numerator, m.Denominator)
}

func (m MixedFraction) ToImproperFraction() Fraction {
	var numerator int

	if m.WholeNumber < 0 {
		numerator = (m.WholeNumber * m.Denominator) + InvertInt(m.Numerator)
	} else {
		numerator = (m.WholeNumber * m.Denominator) + m.Numerator
	}

	return Fraction{
		Numerator:   numerator,
		Denominator: m.Denominator,
	}
}

type Fraction struct {
	Numerator   int
	Denominator int
}

func (i Fraction) ToMixedFraction() MixedFraction {
	var numerator int
	var denominator int

	commonFactor := GreatestCommonFactor(i.Numerator, i.Denominator)
	numerator = i.Numerator / commonFactor
	denominator = i.Denominator / commonFactor

	whole := numerator / denominator
	numerator = numerator % denominator

	return MixedFraction{
		WholeNumber: whole,
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func GreatestCommonFactor(num1 int, num2 int) int {
	factors1 := Factors(num1)
	factors2 := Factors(num2)

	// As we are looking for the highest value we
	// can iterate backwards. This is O(n^2) worst
	// case scenario.
	for i := len(factors1) - 1; i >= 0; i-- {
		for j := len(factors2) - 1; j >= 0; j-- {
			if factors1[i] == factors2[j] {
				return factors1[i]
			}
		}
	}

	return 1
}

func Factors(number int) []int {
	var factors []int
	var sqrt int

	// handle square root for negative numbers by making it positive
	if number < 0 {
		sqrt = int(math.Sqrt(float64(InvertInt(number))))
	} else {
		sqrt = int(math.Sqrt(float64(number)))
	}

	i := 1
	for i <= sqrt {
		if number%i == 0 {
			factors = append(factors, i)
			if number != 1 {
				factors = append(factors, number/i)
			}

			// add inverse of factors
			if number < 0 {
				neg := InvertInt(i)
				factors = append(factors, neg)
				factors = append(factors, number/neg)
			}
		}
		i++
	}
	sort.Ints(factors)
	return factors
}
