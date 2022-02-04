package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanner_Scan_QuestionMark(t *testing.T) {
	a := assert.New(t)
	s := "?"
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(QUERY, tok, "token should be QUERY")
	a.Equal("?", lit, "literal should be '?'")
}

func TestScanner_Scan_Operator(t *testing.T) {
	a := assert.New(t)
	s := "/ "
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(OPERATOR, tok, "token should be OPERATOR")
	a.Equal("/", lit, "literal should be '/'")
}

func TestScanner_Scan_Fraction(t *testing.T) {
	a := assert.New(t)
	s := "32/45"
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(NUMERATOR, tok, "token should be NUMERATOR")
	a.Equal("32", lit, "literal should be '32'")

	tok, lit = scanner.Scan()
	a.Equal(DIVIDER, tok, "token should be DIVIDER")
	a.Equal("/", lit, "literal should be '/'")

	tok, lit = scanner.Scan()
	a.Equal(DENOMINATOR, tok, "token should be DENOMINATOR")
	a.Equal("45", lit, "literal should be '45'")
}

func TestScanner_Scan_NegativeFraction(t *testing.T) {
	a := assert.New(t)
	s := "-32/45"
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(NUMERATOR, tok, "token should be NUMERATOR")
	a.Equal("-32", lit, "literal should be '-32'")

	tok, lit = scanner.Scan()
	a.Equal(DIVIDER, tok, "token should be DIVIDER")
	a.Equal("/", lit, "literal should be '/'")

	tok, lit = scanner.Scan()
	a.Equal(DENOMINATOR, tok, "token should be DENOMINATOR")
	a.Equal("45", lit, "literal should be '45'")
}

func TestScanner_Scan_MixedFraction(t *testing.T) {
	a := assert.New(t)
	s := "27_32/45"
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(WHOLENUMBER, tok, "token should be WHOLENUMBER")
	a.Equal("27", lit, "literal should be '27'")

	tok, lit = scanner.Scan()
	a.Equal(SEPARATOR, tok, "token should be SEPARATOR")
	a.Equal("_", lit, "literal should be '_'")

	tok, lit = scanner.Scan()
	a.Equal(NUMERATOR, tok, "token should be NUMERATOR")
	a.Equal("32", lit, "literal should be '32'")

	tok, lit = scanner.Scan()
	a.Equal(DIVIDER, tok, "token should be DIVIDER")
	a.Equal("/", lit, "literal should be '/'")

	tok, lit = scanner.Scan()
	a.Equal(DENOMINATOR, tok, "token should be DENOMINATOR")
	a.Equal("45", lit, "literal should be '45'")
}

func TestScanner_Scan_NegativeMixedFraction(t *testing.T) {
	a := assert.New(t)
	s := "-27_32/45"
	scanner := NewScanner(s)

	tok, lit := scanner.Scan()
	a.Equal(WHOLENUMBER, tok, "token should be WHOLENUMBER")
	a.Equal("-27", lit, "literal should be '-27'")

	tok, lit = scanner.Scan()
	a.Equal(SEPARATOR, tok, "token should be SEPARATOR")
	a.Equal("_", lit, "literal should be '_'")

	tok, lit = scanner.Scan()
	a.Equal(NUMERATOR, tok, "token should be NUMERATOR")
	a.Equal("32", lit, "literal should be '32'")

	tok, lit = scanner.Scan()
	a.Equal(DIVIDER, tok, "token should be DIVIDER")
	a.Equal("/", lit, "literal should be '/'")

	tok, lit = scanner.Scan()
	a.Equal(DENOMINATOR, tok, "token should be DENOMINATOR")
	a.Equal("45", lit, "literal should be '45'")
}
