package kociemba

type Search struct {
	nodeUD []int
	nodeRL []int
	nodeFB []int
	phase1Cubie CubieCube
}

func newSearch() Search {
	var s Search
	for i := 0; i < 21; i++ {
		s.nodeUD[i] = newCoordCube()
		s.nodeRL[i] = newCoordCube()
		s.nodeFB[i] = newCoordCube()
		s.phase1Cubie[i] = newCubieCube()
	}
	for i := 0; i < 6; i++ {
		s.urfCubieCube[i] = newCubieCube()
		s.urfCoordCube[i] = newCoordCube()
	}
	for i := 0; i < MAX_PRE_MOVES; i++ {
		s.preMoveCubes[i + 1] = newCubieCube()
	}
	return s
}