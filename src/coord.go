package kociemba

type CoordCube struct {
	twist int
    tsym int
    flip int
    fsym int
    slice int
    prun int
    twistc int
    flipc int
}

func newCoordCube() CoordCube {
	return CoordCube{}
}