package kociemba

type Tables struct {
	initLevel int
	UDSliceMove [N_SLICE][N_MOVES]int
	TwistMove [N_TWIST_SYM][N_MOVES]int
    FlipMove [N_FLIP_SYM][N_MOVES]int
    UDSliceConj [N_SLICE][8]int
    UDSliceTwistPrun [N_SLICE * N_TWIST_SYM / 8 + 1]int
    UDSliceFlipPrun [N_SLICE * N_FLIP_SYM / 8 + 1]int
    TwistFlipPrun [N_FLIP * N_TWIST_SYM / 8 + 1]int
}

func initTables(tables *Tables, fullInit bool) {
	if (tables.initLevel == 2 || tables.initLevel == 1 && !fullInit) {
		return
	}

	if (tables.initLevel == 0) {
		// CubieCube.initPermSym2Raw();
		// initCPermMove(tables)
		// initEPermMove(tables)
		// initMPermMoveConj(tables)
		// initCombPMoveConj(tables)

		// CubieCube.initFlipSym2Raw()
		// CubieCube.initTwistSym2Raw()
		// initFlipMove(tables)
		// initTwistMove(tables)
		initUDSliceMoveConj(tables)
	}

	// initMCPermPrun(fullInit);
	// initPermCombPPrun(fullInit);
	// initSliceTwistPrun(fullInit);
	// initSliceFlipPrun(fullInit);
	// initTwistFlipPrun(fullInit);

	if (fullInit) {
		tables.initLevel = 2
	} else {
		tables.initLevel = 1
	}
}

func initUDSliceMoveConj(tables *Tables) {
	c := newCubieCube();
    d := newCubieCube();
	for i := 0; i < N_SLICE; i++ {
		c.setUDSlice(i);
		for j := 0; j < N_MOVES; j += 3 {
			CubieCube.EdgeMult(c, CubieCube.moveCube[j], d);
			UDSliceMove[i][j] = d.getUDSlice();
		}
		for (int j = 0; j < 16; j += 2) {
			CubieCube.EdgeConjugate(c, CubieCube.SymMultInv[0][j], d);
			UDSliceConj[i][j >> 1] = d.getUDSlice();
		}
	}
	for i := 0; i < N_SLICE; i++ {
		for j := 0; j < N_MOVES; j += 3 {
			udslice := UDSliceMove[i][j];
			for (int k = 1; k < 3; k++) {
				udslice = UDSliceMove[udslice][j];
				UDSliceMove[i][j + k] = udslice;
			}
		}
	}
}