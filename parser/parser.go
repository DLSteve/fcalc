package parser

import (
	"fcalc/calculator"
	"fmt"
	"strconv"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(s string) *Parser {
	return &Parser{
		s: NewScanner(s),
	}
}

func (p *Parser) Parse() (calculator.Problem, error) {
	if tok, _ := p.scan(); tok != QUERY {
		return calculator.Problem{}, fmt.Errorf("expecting opening '?'")
	}

	frac1, err := p.scanMixedFraction()
	if err != nil {
		return calculator.Problem{}, err
	}

	tok, lit := p.scan()
	if tok != OPERATOR {
		return calculator.Problem{}, fmt.Errorf("expecting one of the following operators (+, -, *, /)")
	}

	operator := lit

	frac2, err := p.scanMixedFraction()
	if err != nil {
		return calculator.Problem{}, err
	}

	return calculator.Problem{
		Fraction1: frac1,
		Fraction2: frac2,
		Operator:  operator,
	}, nil
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	// skip over whitespace
	if tok == WS {
		tok, lit = p.scan()
	}

	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanMixedFraction() (calculator.MixedFraction, error) {
	frac := calculator.MixedFraction{}
	tok, lit := p.scan()
	if tok != WHOLENUMBER && tok != NUMERATOR {
		return calculator.MixedFraction{}, fmt.Errorf("expecting a regular or mixed fraction")
	}

	if tok == WHOLENUMBER {
		intVal, err := strconv.Atoi(lit)
		if err != nil {
			return calculator.MixedFraction{}, err
		}
		frac.WholeNumber = intVal

		tok, _ = p.scan()
		if tok != SEPARATOR {
			return calculator.MixedFraction{}, fmt.Errorf("expecting a seperator '_' between whole number and fraction")
		}

		numerator, denominator, err := p.scanFraction()
		if err != nil {
			return calculator.MixedFraction{}, err
		}

		frac.Numerator = numerator
		frac.Denominator = denominator
	} else {
		// move back in buffer so scanFraction can pick up the correct token
		p.unscan()
		numerator, denominator, err := p.scanFraction()
		if err != nil {
			return calculator.MixedFraction{}, err
		}

		frac.Numerator = numerator
		frac.Denominator = denominator
	}

	return frac, nil
}

func (p *Parser) scanFraction() (numerator int, denominator int, err error) {
	tok, lit := p.scan()
	if tok != NUMERATOR {
		return 0, 0, fmt.Errorf("expecting numerator for fractional value")
	}

	numVal, err := strconv.Atoi(lit)
	if err != nil {
		return 0, 0, err
	}

	numerator = numVal

	tok, _ = p.scan()
	if tok != DIVIDER {
		return 0, 0, fmt.Errorf("expecting a divider '/' between numerator and denominator")
	}

	tok, lit = p.scan()
	if tok != DENOMINATOR {
		return 0, 0, fmt.Errorf("expecting denominator for fractional value")
	}

	denomVal, err := strconv.Atoi(lit)
	if err != nil {
		return 0, 0, err
	}

	denominator = denomVal

	return
}
