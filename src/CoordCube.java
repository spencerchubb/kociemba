package src;

class CoordCube {
    // TODO these are redundant in CoordCube and Tables
    static final int N_MOVES = 18;
    static final int N_MOVES2 = 10;
    static final int N_SLICE = 495;
    static final int N_TWIST = 2187;
    static final int N_TWIST_SYM = 324;
    static final int N_FLIP = 2048;
    static final int N_FLIP_SYM = 336;
    static final int N_PERM = 40320;
    static final int N_PERM_SYM = 2768;
    static final int N_MPERM = 24;
    static final int N_COMB = 140;
    static final int P2_PARITY_MOVE = 0xA5;

    static synchronized void init(boolean fullInit) {
        Tables.init(fullInit);
    }

    int twist;
    int tsym;
    int flip;
    int fsym;
    int slice;
    int prun;

    int twistc;
    int flipc;

    CoordCube() {
    }

    void set(CoordCube node) {
        this.twist = node.twist;
        this.tsym = node.tsym;
        this.flip = node.flip;
        this.fsym = node.fsym;
        this.slice = node.slice;
        this.prun = node.prun;
        this.twistc = node.twistc;
        this.flipc = node.flipc;
    }

    void calcPruning(boolean isPhase1) {
        prun = Math.max(
                Math.max(
                        Tables.getPruning(Tables.UDSliceTwistPrun,
                                twist * N_SLICE + Tables.UDSliceConj[slice][tsym]),
                        Tables.getPruning(Tables.UDSliceFlipPrun,
                                flip * N_SLICE + Tables.UDSliceConj[slice][fsym])),
                Math.max(
                        Tables.getPruning(Tables.TwistFlipPrun,
                                (twistc >> 3) << 11 | CubieCube.FlipS2RF[flipc ^ (twistc & 7)]),
                        Tables.getPruning(Tables.TwistFlipPrun,
                                twist << 11 | CubieCube.FlipS2RF[flip << 3 | (fsym ^ tsym)])));
    }

    boolean setWithPrun(CubieCube cc, int depth) {
        twist = cc.getTwistSym();
        flip = cc.getFlipSym();
        tsym = twist & 7;
        twist = twist >> 3;

        prun = Tables.getPruning(Tables.TwistFlipPrun,
                twist << 11 | CubieCube.FlipS2RF[flip ^ tsym]);
        if (prun > depth) {
            return false;
        }

        fsym = flip & 7;
        flip = flip >> 3;

        slice = cc.getUDSlice();
        prun = Math.max(prun, Math.max(
                Tables.getPruning(Tables.UDSliceTwistPrun,
                        twist * N_SLICE + Tables.UDSliceConj[slice][tsym]),
                Tables.getPruning(Tables.UDSliceFlipPrun,
                        flip * N_SLICE + Tables.UDSliceConj[slice][fsym])));
        if (prun > depth) {
            return false;
        }

        CubieCube pc = new CubieCube();
        CubieCube.CornConjugate(cc, 1, pc);
        CubieCube.EdgeConjugate(cc, 1, pc);
        twistc = pc.getTwistSym();
        flipc = pc.getFlipSym();
        prun = Math.max(prun,
                Tables.getPruning(Tables.TwistFlipPrun,
                        (twistc >> 3) << 11 | CubieCube.FlipS2RF[flipc ^ (twistc & 7)]));

        return prun <= depth;
    }

    /**
     * @return pruning value
     */
    int doMovePrun(CoordCube cc, int m, boolean isPhase1) {
        slice = Tables.UDSliceMove[cc.slice][m];

        flip = Tables.FlipMove[cc.flip][CubieCube.Sym8Move[m << 3 | cc.fsym]];
        fsym = (flip & 7) ^ cc.fsym;
        flip >>= 3;

        twist = Tables.TwistMove[cc.twist][CubieCube.Sym8Move[m << 3 | cc.tsym]];
        tsym = (twist & 7) ^ cc.tsym;
        twist >>= 3;

        prun = Math.max(
                Math.max(
                        Tables.getPruning(Tables.UDSliceTwistPrun,
                                twist * N_SLICE + Tables.UDSliceConj[slice][tsym]),
                        Tables.getPruning(Tables.UDSliceFlipPrun,
                                flip * N_SLICE + Tables.UDSliceConj[slice][fsym])),
                Tables.getPruning(Tables.TwistFlipPrun,
                        twist << 11 | CubieCube.FlipS2RF[flip << 3 | (fsym ^ tsym)]));
        return prun;
    }

    int doMovePrunConj(CoordCube cc, int m) {
        m = CubieCube.SymMove[3][m];
        flipc = Tables.FlipMove[cc.flipc >> 3][CubieCube.Sym8Move[m << 3 | cc.flipc & 7]] ^ (cc.flipc & 7);
        twistc = Tables.TwistMove[cc.twistc >> 3][CubieCube.Sym8Move[m << 3 | cc.twistc & 7]] ^ (cc.twistc & 7);
        return Tables.getPruning(Tables.TwistFlipPrun,
                (twistc >> 3) << 11 | CubieCube.FlipS2RF[flipc ^ (twistc & 7)]);
    }
}
