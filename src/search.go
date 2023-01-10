package kociemba

type Search struct {
	nodeUD []int
	nodeRL []int
	nodeFB []int
	phase1Cubie CubieCube

	solLen int
	probe int
	probeMax int
	probeMin int
	verbose int
	solution string
	isRec bool

	length1 int
	maxDep2 int

	urfIdx int
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

func solution(search Search, facelets string, maxDepth int, probeMax int, probeMin int, verbose int) (string, error) {
	check := verify(facelets)
	if check != 0 {
		return "", errors.new("solution error: invalid facelets: " + check)
	}
	search.solLen = maxDepth + 1
	search.probe = 0
	search.probeMax = probeMax
	search.probeMin = min(probeMin, probeMax)
	search.verbose = verbose
	search.solution = ""
	search.isRec = false

	// CoordCube.init(false)
	// initSearch()

	return performSearch(search)
}

func performSearch(search Search) (string, error) {
	for search.length1 = search.isRec ? search.length1 : 0; search.length1 < search.solLen; search.length1++ {
		search.maxDep2 = min(MAX_DEPTH2, search.solLen - length1 - 1)
		for search.urfIdx = search.isRec ? search.urfIdx : 0; search.urfIdx < 6; search.urfIdx++ {
			if (conjMask & 1 << urfIdx) != 0 {
				continue
			}
			if (phase1PreMoves(maxPreMoves, -30, urfCubieCube[urfIdx], selfSym & 0xffff)) == 0 {
				if solution == null {
					return "Error 8"
				}
				return toString(solution)
			}
		}
	}
	if solution == null {
		return "Error 7"
	}
	return toString(solution)
}