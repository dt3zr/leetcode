package solveq

import (
	"strconv"
)

const (
	PlusOp = iota
	MinusOp
)

const (
	StateStart = 10 + iota
	StateDigit
)

func solveEquation(equation string) string {
	eq := equation + "\n"

	i := 0
	s := 0
	state := StateStart
	op := PlusOp
	xsum := 0
	ksum := 0

	for i < len(eq) {
		switch eq[i] {
		case 'x':
			if state == StateDigit {
				xval, _ := strconv.Atoi(eq[s:i])
				if op == MinusOp {
					xval = -xval
				}
				xsum += xval
			} else if state == StateStart {
				xval := 1
				if op == MinusOp {
					xval = -xval
				}
				xsum += xval
			}
			i++
			s = i
			state = StateStart
		case '+', '-', '=', '\n':
			if state == StateDigit {
				kval, _ := strconv.Atoi(eq[s:i])
				if op == MinusOp {
					kval = -kval
				}
				ksum += kval
			}
			switch eq[i] {
			case '+':
				op = PlusOp
			case '-':
				op = MinusOp
			case '=', '\n':
				xsum = -xsum
				ksum = -ksum
				op = PlusOp
			}

			i++
			s = i
			state = StateStart

		default:
			if eq[i] >= '0' && eq[i] <= '9' {
				state = StateDigit
			}
			i++
		}
	}

	k := ksum
	x := xsum
	if x == 0 {
		if k == x {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	remainder := k % x
	if remainder != 0 {
		return "No solution"
	} else {
		rval := k / x
		return "x=" + strconv.Itoa(-rval)
	}
}
