package kociemba

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	facelets := "DUUBULDBFRBFRRULLLBRDFFFBLURDBFDFDRFRULBLUFDURRBLBDUDL"
	solution := Solve(facelets)
	fmt.Println(solution)
}
