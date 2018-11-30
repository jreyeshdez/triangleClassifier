package classifier

import "C"
import (
	"errors"
	"fmt"
	"math/big"
	"sort"
)

type Triangle struct {
	A *big.Float
	B *big.Float
	C *big.Float
}

func NewTriangle(a, b, c float64) *Triangle {
	sA := new(big.Float).SetFloat64(a)
	sB := new(big.Float).SetFloat64(b)
	sC := new(big.Float).SetFloat64(c)
	triangle := &Triangle{
		A: sA,
		B: sB,
		C: sC,
	}
	return triangle
}

func (t *Triangle) isValid() bool {
	z := new(big.Float).SetFloat64(0)
	if t.A.Cmp(z) <= 0 || t.B.Cmp(z) <= 0 || t.C.Cmp(z) <= 0 {
		return false
	}
	return true
}

// As per triangle inequality theorem,
// the sum of the side lengths of any 2 sides of a triangle
// must exceed the length of the third side
func (t *Triangle) canFormTriangle() bool {
	c := new(big.Float).Add(t.A, t.B)
	b := new(big.Float).Add(t.A, t.C)
	a := new(big.Float).Add(t.B, t.C)
	if c.Cmp(t.C) <= 0 || b.Cmp(t.B) <= 0 || a.Cmp(t.A) <= 0 {
		return false
	}
	return true
}

// Returns the type of the triangle if it can be classified or an error
func (t *Triangle) GetType() (string, error) {
	if !t.isValid() {
		return "", errors.New("given sides must be positive")
	}

	if !t.canFormTriangle() {
		return "", errors.New(fmt.Sprintf("sides with length %f, %f, %f can't form a triangle", t.A, t.B, t.C))
	}

	if rs := t.isEquilateral(); rs != "" {
		return rs, nil
	}

	if rs := t.isScaleneOrRightScalene(); rs != "" {
		return rs, nil
	}

	if rs := t.isIsoscelesOrRightIsosceles(); rs != "" {
		return rs, nil
	}
	return "", nil
}

// Equilateral triangles have all three sides with the same length
func (t *Triangle) isEquilateral() string {
	if t.A.Cmp(t.B) == 0 && t.B.Cmp(t.C) == 0 {
		return "Equilateral"
	}
	return ""
}

// Isosceles triangles have two sides with the same length
func (t *Triangle) isIsosceles() string {
	if t.A.Cmp(t.B) == 0 || t.A.Cmp(t.C) == 0 || t.B.Cmp(t.C) == 0 {
		return "Isosceles"
	}
	return ""
}

// Right triangles have three sides where a^2 + b^2 = c^2
// A Right triangle may be isosceles or scalene
// Since Floating-point numbers are not exact,
// and comparisons using == will often fail I decided to use c^2 - (a^2 + b^2) ~= 0
func (t *Triangle) isRight() bool {
	s := []*big.Float{t.A, t.B, t.C}
	rs := make([]*big.Float, 0, len(s))
	for _, n := range s {
		rs = append(rs, new(big.Float).Mul(n, n))
	}

	sort.Slice(rs, func(i, j int) bool { return rs[i].Cmp(rs[j]) < 0 })

	a := new(big.Float).Add(rs[0], rs[1])
	// Take the absolute difference between the two numbers
	// and check it's very small
	b := new(big.Float).Sub(rs[2], a)
	abs := new(big.Float).Abs(b)

	if abs.Cmp(new(big.Float).SetFloat64(0.013)) <= 0 {
		return true
	}
	return false
}

// Scalene triangles have three sides with different lengths
// A right triangle may be isosceles or scalene
func (t *Triangle) isScaleneOrRightScalene() string {
	if !t.isRight() && t.isIsosceles() == "" {
		return "Scalene"
	}

	if t.isRight() && t.isIsosceles() == "" {
		return "Right Scalene"
	}
	return ""
}

func (t *Triangle) isIsoscelesOrRightIsosceles() string {
	if !t.isRight() && t.isIsosceles() != "" {
		return "Isosceles"
	}

	if t.isRight() && t.isIsosceles() != "" {
		return "Right Isosceles"
	}
	return ""
}
