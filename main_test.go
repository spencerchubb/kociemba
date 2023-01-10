package kociemba

import (
	"testing"
)

func TestSolve(t *testing.T) {
	facelets := "DUUBULDBFRBFRRULLLBRDFFFBLURDBFDFDRFRULBLUFDURRBLBDUDL"
	solution, err := Solve(facelets)
	if err != nil {
		t.Error(err)
	}
	expected := "R2 U2 B2 L2 F2 U' L2 R2 B2 R2 D B2 F L' F U2 F' R' D' L2 R'"
	if solution != expected {
		t.Errorf("Got %s, expected %s", solution, expected)
	}
}
