// fibonacci.go
// description: Get the nth Fibonacci Number
// details:
// In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. [Fibonacci number](https://en.wikipedia.org/wiki/Fibonacci_number)
// author(s) [red_byte](https://github.com/i-redbyte)
// see fibonacci_test.go

package fibonacci

// Matrix This function calculates the n-th fibonacci number using the matrix method. [See](https://en.wikipedia.org/wiki/Fibonacci_number#Matrix_form)
func Matrix(n uint) uint {
	a, ta, b, tb := 1, 1, 1, 1
	c, rc, tc := 1, 0, 0
	d, rd := 0, 1

	for n != 0 {
		if n&1 == 1 {
			tc = rc
			rc = rc*a + rd*c
			rd = tc*b + rd*d
		}

		ta = a
		tb = b
		tc = c
		a = a*a + b*c
		b = ta*b + b*d
		c = c*ta + d*c
		d = tc*tb + d*d

		n >>= 1
	}
	return uint(rc)
}
