package kociemba

type CubieCube struct {
	ca []int
    ea []int
    temps CubieCube
}

func newCubieCube() CubieCube {
	var cc CubieCube
	cc.ca = {0, 1, 2, 3, 4, 5, 6, 7}
	cc.ea = {0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}
	return cc
}