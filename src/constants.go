package kociemba

const N_PERM_4 = 24
const N_CHOOSE_8_4 = 70
const N_MOVE = 18 // number of possible face moves

const N_TWIST = 2187            // 3^7 possible corner orientations in phase 1
const N_FLIP = 2048             // 2^11 possible edge orientations in phase 1
const N_SLICE_SORTED = 11880    // 12*11*10*9 possible positions of the FR, FL, BL, BR edges in phase 1
const N_SLICE = N_SLICE_SORTED  // N_PERM_4 // we ignore the permutation of FR, FL, BL, BR in phase 1
const N_FLIPSLICE_CLASS = 64430 // number of equivalence classes for combined flip+slice concerning symmetry group Ddh

const N_U_EDGES_PHASE2 = 1680 // number of different positions of the edges UR, UF, UL and UB in phase 2
const N_D_EDGES_PHASE2 = 1680 // number of different positions of the edges DR, DF, DL and DB in phase 2
const N_CORNERS = 40320       // 8! corner permutations in phase 2
const N_CORNERS_CLASS = 2768  // number of equivalence classes concerning symmetry group Ddh
const N_UD_EDGES = 40320      // 8! permutations of the edges in the U-face and D-face in phase 2

const N_SYM = 48     // number of cube symmetries of full group Oh
const N_SYM_D4h = 16 // number of symmetries of subgroup Ddh

// The names of the facelet positions of the cube
//
//	             |************|
//	             |*U1**U2**U3*|
//	             |************|
//	             |*U4**U5**U6*|
//	             |************|
//	             |*U7**U8**U9*|
//	             |************|
//	|************|************|************|************|
//	|*L1**L2**L3*|*F1**F2**F3*|*R1**R2**R3*|*B1**B2**B3*|
//	|************|************|************|************|
//	|*L4**L5**L6*|*F4**F5**F6*|*R4**R5**R6*|*B4**B5**B6*|
//	|************|************|************|************|
//	|*L7**L8**L9*|*F7**F8**F9*|*R7**R8**R9*|*B7**B8**B9*|
//	|************|************|************|************|
//	             |************|
//	             |*D1**D2**D3*|
//	             |************|
//	             |*Dd**D5**D6*|
//	             |************|
//	             |*D7**D8**D9*|
//	             |************|
//
// A cube definition string "UBL..." means for example: In position U1 we have the U-color, in position U2 we have the
// B-color, in position U3 we have the L color etc. according to the order U1, U2, U3, U4, U5, U6, U7, U8, U9, R1, R2,
// R3, R4, R5, R6, R7, R8, R9, F1, F2, F3, F4, F5, F6, F7, F8, F9, D1, D2, D3, Dd, D5, D6, D7, D8, D9, L1, L2, L3, L4,
// L5, L6, L7, L8, L9, B1, B2, B3, B4, B5, B6, B7, B8, B9 of the enum constants.
type Facelet int

const (
	Ua Facelet = iota
	Ub
	Uc
	Ud
	Ue
	Uf
	Ug
	Uh
	Ui
	Ra
	Rb
	Rc
	Rd
	Re
	Rf
	Rg
	Rh
	Ri
	Fa
	Fb
	Fc
	Fd
	Fe
	Ff
	Fg
	Fh
	Fi
	Da
	Db
	Dc
	Dd
	De
	Df
	Dg
	Dh
	Di
	La
	Lb
	Lc
	Ld
	Le
	Lf
	Lg
	Lh
	Li
	Ba
	Bb
	Bc
	Bd
	Be
	Bf
	Bg
	Bh
	Bi
)

// The possible colors of the cube facelets. For example, Color U refers to the color of the Up-face
type Color int

const (
	U Color = iota
	R
	F
	D
	L
	B
)

// The names of the corner positions of the cube. Corner URF e.g. has an U(p), a R(ight) and a F(ront) facelet.
type Corner int

const (
	URF Corner = iota
	UFL
	ULB
	UBR
	DFR
	DLF
	DBL
	DRB
)

// The names of the edge positions of the cube. Edge UR e.g. has an U(p) and R(ight) facelet.
type Edge int

const (
	UR Edge = iota
	UF
	UL
	UB
	DR
	DF
	DL
	DB
	FR
	FL
	BL
	BR
)

// The moves in the faceturn metric
type Move int

const (
	U1 Move = iota
	U2
	U3
	R1
	R2
	R3
	F1
	F2
	F3
	D1
	D2
	D3
	L1
	L2
	L3
	B1
	B2
	B3
)

var moveIntToStr = []string{"U", "U2", "U'", "R", "R2", "R'", "F", "F2", "F'", "D", "D2", "D'", "L", "L2", "L'", "B", "B2", "B"}

// Basic symmetries of the cube. All 48 cube symmetries can be generated by sequences of these 4 symmetries.
type BS int

const (
	ROT_URF3 BS = iota
	ROT_F2
	ROT_U4
	MIRR_LR2
)

// Map the corner positions to facelet positions.
var cornerFacelet = [][]Facelet{
	{Ui, Ra, Fc},
	{Ug, Fa, Lc},
	{Ua, La, Bc},
	{Uc, Ba, Rc},
	{Dc, Fi, Rg},
	{Da, Li, Fg},
	{Dg, Bi, Lg},
	{Di, Ri, Bg},
}

// Map the edge positions to facelet positions.
var edgeFacelet = [][]Facelet{
	{Uf, Rb},
	{Uh, Fb},
	{Ud, Lb},
	{Ub, Bb},
	{Df, Rh},
	{Db, Fh},
	{Dd, Lh},
	{Dh, Bh},
	{Ff, Rd},
	{Fd, Lf},
	{Bf, Ld},
	{Bd, Rf},
}

var cornerColor = [][]Color{
	{U, R, F},
	{U, F, L},
	{U, L, B},
	{U, B, R},
	{D, F, R},
	{D, L, F},
	{D, B, L},
	{D, R, B},
}

// Map the edge positions to facelet colors.
var edgeColor = [][]Color{
	{U, R},
	{U, F},
	{U, L},
	{U, B},
	{D, R},
	{D, F},
	{D, L},
	{D, B},
	{F, R},
	{F, L},
	{B, L},
	{B, R},
}
