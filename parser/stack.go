package parser

import (
	"bufio"
	"strings"
)

type RuneStack struct {
	runes []rune
}

func NewRuneStack(s string) RuneStack {
	var rs []rune
	r := bufio.NewReader(strings.NewReader(s))
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			break
		}
		rs = append([]rune{ch}, rs...)
	}

	return RuneStack{
		runes: rs,
	}
}

func (r *RuneStack) Push(rns ...rune) {
	for _, rn := range rns {
		r.runes = append(r.runes, rn)
	}
}

func (r *RuneStack) Pop() rune {
	n := len(r.runes) - 1
	if n < 0 {
		return rune(0)
	}

	rn := r.runes[n]
	r.runes = r.runes[:n]
	return rn
}

func (r *RuneStack) Peek() rune {
	n := len(r.runes) - 1
	if n < 0 {
		return rune(0)
	}

	return r.runes[n]
}
