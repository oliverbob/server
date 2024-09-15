package copsnrobbers

// See: https://freeslotshub.com/playngo/cop-the-lot/

import (
	"math"
	"math/rand/v2"

	slot "github.com/slotopol/server/game/slot"
	"github.com/slotopol/server/util"
)

// reels lengths [39, 39, 39, 39, 39], total reshuffles 90224199
// symbols: 64.043(lined) + 10.943(scatter) = 74.985499%
// free spins 6200631, q = 0.068725
// free games frequency: 1/247.36
// RTP = 74.985(sym) + 1.3*0.068725*120.32(fg) = 85.735488%
var ReelsReg86 = slot.Reels5x{
	{6, 9, 11, 7, 8, 10, 9, 2, 10, 9, 11, 6, 5, 8, 10, 1, 11, 9, 7, 10, 4, 6, 7, 8, 5, 6, 8, 9, 11, 4, 8, 7, 3, 11, 10, 5, 11, 3, 10},
	{7, 9, 11, 3, 8, 9, 5, 10, 9, 11, 10, 7, 4, 10, 11, 9, 5, 11, 7, 6, 9, 4, 7, 10, 11, 1, 6, 10, 5, 8, 6, 2, 8, 3, 10, 8, 6, 11, 8},
	{6, 11, 8, 9, 11, 8, 4, 6, 10, 7, 3, 10, 5, 3, 4, 11, 7, 9, 8, 6, 2, 11, 10, 9, 11, 10, 6, 8, 11, 5, 1, 8, 10, 7, 9, 10, 7, 5, 9},
	{10, 6, 2, 11, 9, 7, 10, 8, 11, 6, 4, 9, 8, 10, 7, 9, 8, 11, 10, 5, 1, 10, 9, 8, 11, 3, 7, 8, 6, 10, 7, 3, 9, 5, 11, 6, 4, 11, 5},
	{9, 11, 10, 8, 7, 10, 11, 9, 8, 10, 9, 3, 10, 2, 11, 6, 8, 7, 10, 5, 11, 4, 7, 9, 11, 4, 8, 6, 11, 9, 1, 7, 10, 6, 3, 5, 8, 6, 5},
}

// reels lengths [37, 39, 39, 39, 37], total reshuffles 81207711
// symbols: 64.894(lined) + 11.411(scatter) = 76.304749%
// free spins 5932575, q = 0.073054
// free games frequency: 1/232.7
// RTP = 76.305(sym) + 1.3*0.073054*120.32(fg) = 87.731983%
var ReelsReg88 = slot.Reels5x{
	{4, 9, 11, 5, 6, 8, 10, 3, 6, 9, 10, 5, 11, 10, 8, 6, 7, 11, 2, 3, 10, 9, 6, 8, 7, 9, 4, 5, 11, 9, 10, 7, 8, 11, 7, 1, 8},
	{7, 9, 11, 3, 8, 9, 5, 10, 9, 11, 10, 7, 4, 10, 11, 9, 5, 11, 7, 6, 9, 4, 7, 10, 11, 1, 6, 10, 5, 8, 6, 2, 8, 3, 10, 8, 6, 11, 8},
	{6, 11, 8, 9, 11, 8, 4, 6, 10, 7, 3, 10, 5, 3, 4, 11, 7, 9, 8, 6, 2, 11, 10, 9, 11, 10, 6, 8, 11, 5, 1, 8, 10, 7, 9, 10, 7, 5, 9},
	{10, 6, 2, 11, 9, 7, 10, 8, 11, 6, 4, 9, 8, 10, 7, 9, 8, 11, 10, 5, 1, 10, 9, 8, 11, 3, 7, 8, 6, 10, 7, 3, 9, 5, 11, 6, 4, 11, 5},
	{9, 8, 5, 10, 8, 11, 10, 3, 7, 6, 11, 7, 10, 5, 11, 6, 9, 2, 11, 9, 10, 7, 3, 9, 8, 6, 11, 4, 8, 10, 6, 9, 7, 4, 8, 5, 1},
}

// reels lengths [37, 37, 39, 37, 39], total reshuffles 77043213
// symbols: 66.554(lined) + 11.65(scatter) = 78.203478%
// free spins 5801301, q = 0.075299
// free games frequency: 1/225.77
// RTP = 78.203(sym) + 1.3*0.075299*120.32(fg) = 89.981874%
var ReelsReg90 = slot.Reels5x{
	{4, 9, 11, 5, 6, 8, 10, 3, 6, 9, 10, 5, 11, 10, 8, 6, 7, 11, 2, 3, 10, 9, 6, 8, 7, 9, 4, 5, 11, 9, 10, 7, 8, 11, 7, 1, 8},
	{2, 7, 10, 11, 9, 10, 7, 9, 8, 11, 9, 1, 10, 8, 11, 4, 6, 7, 8, 11, 10, 6, 9, 5, 6, 8, 10, 5, 3, 11, 9, 4, 8, 5, 6, 7, 3},
	{6, 11, 8, 9, 11, 8, 4, 6, 10, 7, 3, 10, 5, 3, 4, 11, 7, 9, 8, 6, 2, 11, 10, 9, 11, 10, 6, 8, 11, 5, 1, 8, 10, 7, 9, 10, 7, 5, 9},
	{3, 10, 8, 7, 6, 3, 7, 9, 8, 1, 6, 9, 7, 6, 11, 9, 4, 10, 11, 8, 9, 11, 5, 6, 8, 11, 10, 5, 9, 10, 8, 7, 10, 5, 4, 11, 2},
	{9, 11, 10, 8, 7, 10, 11, 9, 8, 10, 9, 3, 10, 2, 11, 6, 8, 7, 10, 5, 11, 4, 7, 9, 11, 4, 8, 6, 11, 9, 1, 7, 10, 6, 3, 5, 8, 6, 5},
}

// reels lengths [37, 37, 37, 37, 37], total reshuffles 69343957
// symbols: 67.855(lined) + 12.137(scatter) = 79.991994%
// free spins 5544261, q = 0.079953
// free games frequency: 1/212.62
// RTP = 79.992(sym) + 1.3*0.079953*120.32(fg) = 92.498333%
var ReelsReg92 = slot.Reels5x{
	{4, 9, 11, 5, 6, 8, 10, 3, 6, 9, 10, 5, 11, 10, 8, 6, 7, 11, 2, 3, 10, 9, 6, 8, 7, 9, 4, 5, 11, 9, 10, 7, 8, 11, 7, 1, 8},
	{2, 7, 10, 11, 9, 10, 7, 9, 8, 11, 9, 1, 10, 8, 11, 4, 6, 7, 8, 11, 10, 6, 9, 5, 6, 8, 10, 5, 3, 11, 9, 4, 8, 5, 6, 7, 3},
	{7, 9, 11, 6, 7, 11, 8, 9, 10, 2, 5, 10, 11, 8, 3, 1, 8, 7, 11, 8, 10, 6, 8, 10, 11, 9, 7, 5, 10, 4, 6, 9, 5, 6, 9, 4, 3},
	{3, 10, 8, 7, 6, 3, 7, 9, 8, 1, 6, 9, 7, 6, 11, 9, 4, 10, 11, 8, 9, 11, 5, 6, 8, 11, 10, 5, 9, 10, 8, 7, 10, 5, 4, 11, 2},
	{9, 8, 5, 10, 8, 11, 10, 3, 7, 6, 11, 7, 10, 5, 11, 6, 9, 2, 11, 9, 10, 7, 3, 9, 8, 6, 11, 4, 8, 10, 6, 9, 7, 4, 8, 5, 1},
}

// reels lengths [36, 37, 37, 37, 36], total reshuffles 65646288
// symbols: 67.89(lined) + 12.403(scatter) = 80.293308%
// free spins 5417118, q = 0.082520
// free games frequency: 1/206.01
// RTP = 80.293(sym) + 1.3*0.08252*120.32(fg) = 93.201139%
var ReelsReg93 = slot.Reels5x{
	{11, 4, 7, 8, 6, 5, 8, 9, 4, 10, 7, 1, 9, 11, 5, 2, 8, 11, 6, 8, 9, 11, 10, 9, 3, 7, 8, 10, 9, 3, 10, 5, 6, 11, 7, 10},
	{2, 7, 10, 11, 9, 10, 7, 9, 8, 11, 9, 1, 10, 8, 11, 4, 6, 7, 8, 11, 10, 6, 9, 5, 6, 8, 10, 5, 3, 11, 9, 4, 8, 5, 6, 7, 3},
	{7, 9, 11, 6, 7, 11, 8, 9, 10, 2, 5, 10, 11, 8, 3, 1, 8, 7, 11, 8, 10, 6, 8, 10, 11, 9, 7, 5, 10, 4, 6, 9, 5, 6, 9, 4, 3},
	{3, 10, 8, 7, 6, 3, 7, 9, 8, 1, 6, 9, 7, 6, 11, 9, 4, 10, 11, 8, 9, 11, 5, 6, 8, 11, 10, 5, 9, 10, 8, 7, 10, 5, 4, 11, 2},
	{10, 3, 4, 10, 6, 3, 5, 6, 8, 10, 7, 9, 1, 11, 10, 5, 8, 11, 9, 5, 4, 6, 7, 10, 11, 8, 9, 11, 7, 9, 8, 2, 9, 8, 7, 11},
}

// reels lengths [36, 36, 36, 37, 37], total reshuffles 63872064
// symbols: 68.414(lined) + 12.538(scatter) = 80.951776%
// free spins 5354235, q = 0.083827
// free games frequency: 1/202.8
// RTP = 80.952(sym) + 1.3*0.083827*120.32(fg) = 94.064159%
var ReelsReg94 = slot.Reels5x{
	{11, 4, 7, 8, 6, 5, 8, 9, 4, 10, 7, 1, 9, 11, 5, 2, 8, 11, 6, 8, 9, 11, 10, 9, 3, 7, 8, 10, 9, 3, 10, 5, 6, 11, 7, 10},
	{5, 9, 11, 6, 4, 8, 9, 2, 10, 9, 7, 8, 10, 5, 11, 8, 6, 11, 8, 7, 3, 10, 11, 1, 10, 8, 9, 11, 10, 9, 6, 4, 7, 5, 3, 7},
	{4, 6, 9, 11, 7, 1, 8, 6, 9, 11, 5, 9, 10, 8, 2, 10, 7, 5, 6, 11, 8, 10, 11, 7, 3, 4, 10, 8, 11, 9, 8, 7, 10, 5, 3, 9},
	{3, 10, 8, 7, 6, 3, 7, 9, 8, 1, 6, 9, 7, 6, 11, 9, 4, 10, 11, 8, 9, 11, 5, 6, 8, 11, 10, 5, 9, 10, 8, 7, 10, 5, 4, 11, 2},
	{9, 8, 5, 10, 8, 11, 10, 3, 7, 6, 11, 7, 10, 5, 11, 6, 9, 2, 11, 9, 10, 7, 3, 9, 8, 6, 11, 4, 8, 10, 6, 9, 7, 4, 8, 5, 1},
}

// reels lengths [36, 36, 36, 36, 36], total reshuffles 60466176
// symbols: 69.031(lined) + 12.81(scatter) = 81.841087%
// free spins 5229846, q = 0.086492
// free games frequency: 1/196.55
// RTP = 81.841(sym) + 1.3*0.086492*120.32(fg) = 95.370270%
var ReelsReg95 = slot.Reels5x{
	{11, 4, 7, 8, 6, 5, 8, 9, 4, 10, 7, 1, 9, 11, 5, 2, 8, 11, 6, 8, 9, 11, 10, 9, 3, 7, 8, 10, 9, 3, 10, 5, 6, 11, 7, 10},
	{5, 9, 11, 6, 4, 8, 9, 2, 10, 9, 7, 8, 10, 5, 11, 8, 6, 11, 8, 7, 3, 10, 11, 1, 10, 8, 9, 11, 10, 9, 6, 4, 7, 5, 3, 7},
	{4, 6, 9, 11, 7, 1, 8, 6, 9, 11, 5, 9, 10, 8, 2, 10, 7, 5, 6, 11, 8, 10, 11, 7, 3, 4, 10, 8, 11, 9, 8, 7, 10, 5, 3, 9},
	{5, 7, 6, 11, 8, 4, 11, 9, 10, 3, 9, 11, 8, 1, 7, 9, 8, 10, 6, 11, 5, 8, 9, 11, 10, 8, 9, 2, 7, 10, 5, 3, 7, 10, 4, 6},
	{10, 3, 4, 10, 6, 3, 5, 6, 8, 10, 7, 9, 1, 11, 10, 5, 8, 11, 9, 5, 4, 6, 7, 10, 11, 8, 9, 11, 7, 9, 8, 2, 9, 8, 7, 11},
}

// reels lengths [35, 36, 36, 36, 35], total reshuffles 57153600
// symbols: 69.265(lined) + 13.099(scatter) = 82.364829%
// free spins 5106375, q = 0.089345
// free games frequency: 1/190.27
// RTP = 82.365(sym) + 1.3*0.089345*120.32(fg) = 96.340230%
var ReelsReg96 = slot.Reels5x{
	{8, 10, 7, 11, 9, 7, 3, 10, 9, 7, 5, 4, 10, 11, 8, 9, 6, 11, 5, 1, 8, 6, 11, 2, 8, 6, 10, 4, 5, 9, 11, 10, 8, 9, 3},
	{5, 9, 11, 6, 4, 8, 9, 2, 10, 9, 7, 8, 10, 5, 11, 8, 6, 11, 8, 7, 3, 10, 11, 1, 10, 8, 9, 11, 10, 9, 6, 4, 7, 5, 3, 7},
	{4, 6, 9, 11, 7, 1, 8, 6, 9, 11, 5, 9, 10, 8, 2, 10, 7, 5, 6, 11, 8, 10, 11, 7, 3, 4, 10, 8, 11, 9, 8, 7, 10, 5, 3, 9},
	{5, 7, 6, 11, 8, 4, 11, 9, 10, 3, 9, 11, 8, 1, 7, 9, 8, 10, 6, 11, 5, 8, 9, 11, 10, 8, 9, 2, 7, 10, 5, 3, 7, 10, 4, 6},
	{11, 6, 1, 5, 10, 7, 5, 8, 3, 11, 8, 4, 11, 9, 10, 8, 6, 3, 10, 9, 4, 10, 8, 9, 7, 5, 2, 9, 11, 7, 6, 11, 10, 8, 9},
}

// reels lengths [36, 35, 35, 35, 36], total reshuffles 55566000
// symbols: 69.725(lined) + 13.246(scatter) = 82.970655%
// free spins 5045328, q = 0.090799
// free games frequency: 1/187.23
// RTP = 82.971(sym) + 1.3*0.090799*120.32(fg) = 97.173502%
var ReelsReg97 = slot.Reels5x{
	{11, 4, 7, 8, 6, 5, 8, 9, 4, 10, 7, 1, 9, 11, 5, 2, 8, 11, 6, 8, 9, 11, 10, 9, 3, 7, 8, 10, 9, 3, 10, 5, 6, 11, 7, 10},
	{10, 9, 8, 7, 10, 6, 8, 10, 6, 11, 4, 9, 11, 2, 10, 7, 11, 8, 3, 9, 4, 8, 9, 3, 7, 5, 9, 10, 5, 11, 6, 8, 1, 5, 11},
	{10, 8, 11, 10, 8, 7, 9, 8, 7, 3, 4, 11, 10, 5, 6, 1, 11, 10, 6, 8, 4, 9, 5, 10, 9, 5, 3, 9, 7, 6, 2, 11, 9, 8, 11},
	{8, 11, 9, 10, 6, 9, 8, 4, 6, 11, 8, 9, 5, 1, 6, 3, 10, 11, 5, 10, 2, 9, 8, 7, 11, 10, 4, 11, 7, 10, 8, 3, 7, 5, 9},
	{10, 3, 4, 10, 6, 3, 5, 6, 8, 10, 7, 9, 1, 11, 10, 5, 8, 11, 9, 5, 4, 6, 7, 10, 11, 8, 9, 11, 7, 9, 8, 2, 9, 8, 7, 11},
}

// reels lengths [35, 35, 36, 35, 35], total reshuffles 54022500
// symbols: 70.074(lined) + 13.393(scatter) = 83.467472%
// free spins 4984740, q = 0.092272
// free games frequency: 1/184.24
// RTP = 83.467(sym) + 1.3*0.092272*120.32(fg) = 97.900684%
var ReelsReg98 = slot.Reels5x{
	{8, 10, 7, 11, 9, 7, 3, 10, 9, 7, 5, 4, 10, 11, 8, 9, 6, 11, 5, 1, 8, 6, 11, 2, 8, 6, 10, 4, 5, 9, 11, 10, 8, 9, 3},
	{10, 9, 8, 7, 10, 6, 8, 10, 6, 11, 4, 9, 11, 2, 10, 7, 11, 8, 3, 9, 4, 8, 9, 3, 7, 5, 9, 10, 5, 11, 6, 8, 1, 5, 11},
	{4, 6, 9, 11, 7, 1, 8, 6, 9, 11, 5, 9, 10, 8, 2, 10, 7, 5, 6, 11, 8, 10, 11, 7, 3, 4, 10, 8, 11, 9, 8, 7, 10, 5, 3, 9},
	{8, 11, 9, 10, 6, 9, 8, 4, 6, 11, 8, 9, 5, 1, 6, 3, 10, 11, 5, 10, 2, 9, 8, 7, 11, 10, 4, 11, 7, 10, 8, 3, 7, 5, 9},
	{11, 6, 1, 5, 10, 7, 5, 8, 3, 11, 8, 4, 11, 9, 10, 8, 6, 3, 10, 9, 4, 10, 8, 9, 7, 5, 2, 9, 11, 7, 6, 11, 10, 8, 9},
}

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 70.739(lined) + 13.542(scatter) = 84.280797%
// free spins 4924611, q = 0.093763
// free games frequency: 1/181.31
// RTP = 84.281(sym) + 1.3*0.093763*120.32(fg) = 98.947310%
var ReelsReg99 = slot.Reels5x{
	{8, 10, 7, 11, 9, 7, 3, 10, 9, 7, 5, 4, 10, 11, 8, 9, 6, 11, 5, 1, 8, 6, 11, 2, 8, 6, 10, 4, 5, 9, 11, 10, 8, 9, 3},
	{10, 9, 8, 7, 10, 6, 8, 10, 6, 11, 4, 9, 11, 2, 10, 7, 11, 8, 3, 9, 4, 8, 9, 3, 7, 5, 9, 10, 5, 11, 6, 8, 1, 5, 11},
	{10, 8, 11, 10, 8, 7, 9, 8, 7, 3, 4, 11, 10, 5, 6, 1, 11, 10, 6, 8, 4, 9, 5, 10, 9, 5, 3, 9, 7, 6, 2, 11, 9, 8, 11},
	{8, 11, 9, 10, 6, 9, 8, 4, 6, 11, 8, 9, 5, 1, 6, 3, 10, 11, 5, 10, 2, 9, 8, 7, 11, 10, 4, 11, 7, 10, 8, 3, 7, 5, 9},
	{11, 6, 1, 5, 10, 7, 5, 8, 3, 11, 8, 4, 11, 9, 10, 8, 6, 3, 10, 9, 4, 10, 8, 9, 7, 5, 2, 9, 11, 7, 6, 11, 10, 8, 9},
}

// reels lengths [32, 32, 34, 32, 32], total reshuffles 35651584
// symbols: 77.748(lined) + 15.792(scatter) = 93.540206%
// free spins 4173228, q = 0.117056
// free games frequency: 1/145.23
// RTP = 93.54(sym) + 1.3*0.11706*120.32(fg) = 111.850209%
var ReelsReg112 = slot.Reels5x{
	{5, 7, 9, 2, 10, 7, 3, 6, 5, 11, 9, 4, 11, 7, 8, 1, 6, 10, 4, 8, 9, 10, 7, 5, 3, 8, 6, 11, 10, 9, 11, 8},
	{8, 4, 9, 1, 8, 11, 6, 3, 10, 4, 11, 5, 7, 6, 10, 7, 8, 9, 5, 10, 3, 7, 11, 10, 6, 8, 9, 2, 7, 11, 9, 5},
	{10, 6, 11, 7, 10, 6, 11, 8, 2, 7, 3, 6, 9, 4, 7, 9, 8, 5, 4, 11, 8, 5, 11, 3, 8, 5, 10, 4, 1, 9, 7, 10, 9, 3},
	{2, 10, 8, 3, 10, 9, 11, 1, 7, 11, 4, 7, 3, 11, 5, 8, 9, 6, 8, 5, 6, 11, 8, 6, 10, 5, 9, 10, 7, 9, 4, 7},
	{3, 11, 6, 7, 11, 6, 9, 1, 7, 3, 10, 5, 2, 6, 10, 5, 9, 8, 7, 4, 8, 11, 5, 9, 10, 4, 9, 8, 11, 7, 8, 10},
}

// reels lengths [28, 28, 28, 28, 28], total reshuffles 17210368
// RTP = 120.32(lined) + 0(scatter) = 120.323871%
var ReelsBon = slot.Reels5x{
	{9, 4, 6, 7, 5, 9, 3, 11, 5, 10, 11, 8, 7, 11, 10, 3, 8, 1, 10, 4, 5, 7, 6, 4, 8, 9, 3, 6},
	{10, 3, 5, 8, 3, 6, 10, 7, 11, 9, 5, 6, 10, 11, 3, 7, 5, 11, 4, 8, 1, 7, 4, 9, 8, 4, 6, 9},
	{7, 10, 5, 9, 8, 7, 3, 8, 6, 11, 4, 1, 6, 9, 11, 4, 7, 3, 10, 9, 11, 10, 5, 4, 3, 5, 6, 8},
	{11, 7, 6, 3, 7, 9, 10, 8, 4, 11, 3, 4, 11, 6, 5, 10, 8, 3, 5, 7, 6, 4, 9, 10, 1, 8, 9, 5},
	{11, 3, 5, 7, 8, 5, 10, 4, 11, 10, 9, 8, 4, 6, 1, 9, 11, 3, 9, 7, 3, 8, 5, 6, 7, 4, 6, 10},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	85.735488:  &ReelsReg86,
	87.731983:  &ReelsReg88,
	89.981874:  &ReelsReg90,
	92.498333:  &ReelsReg92,
	93.201139:  &ReelsReg93,
	94.064159:  &ReelsReg94,
	95.370270:  &ReelsReg95,
	96.340230:  &ReelsReg96,
	97.173502:  &ReelsReg97,
	97.900684:  &ReelsReg98,
	98.947310:  &ReelsReg99,
	111.850209: &ReelsReg112,
}

func FindReels(mrtp float64) (rtp float64, reels slot.Reels) {
	for p, r := range ReelsMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}

// Lined payment.
var LinePay = [11][5]float64{
	{0, 3, 30, 300, 3000}, //  1 wild
	{0, 0, 0, 0, 0},       //  2 scatter
	{0, 2, 25, 150, 750},  //  3 money bag
	{0, 2, 20, 100, 500},  //  4 diamonds
	{0, 2, 15, 75, 500},   //  5 robbery
	{0, 0, 15, 75, 250},   //  6 picture
	{0, 0, 10, 75, 250},   //  7 watch
	{0, 0, 5, 50, 150},    //  8 cop
	{0, 0, 5, 50, 125},    //  9 jail
	{0, 0, 5, 25, 100},    // 10 thief
	{0, 0, 5, 25, 100},    // 11 handcuffs
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 3, 25, 250} // 2 scatter

var ScatRand = []int{10, 15, 15, 20, 25}

const (
	Efs = 17  // average free spins for ScatRand set
	Pfs = 0.3 // probability of "got away" at free spins
)

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
	// multiplier for free spins, if them ended by "got away"
	M float64 `json:"m,omitempty" yaml:"m,omitempty" xml:"m,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: util.MakeBitNum(9, 1),
			Bet: 1,
		},
		FS: 0,
		M:  0,
	}
}

const wild, scat = 1, 2

var bl = slot.Lineset5x{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{2, 1, 1, 1, 2}, // 6
	{2, 3, 3, 3, 2}, // 7
	{1, 1, 2, 3, 3}, // 8
	{3, 3, 2, 1, 1}, // 9
}

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	if g.FS == 0 {
		g.ScanScatters(screen, wins)
	}
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := range g.Sel.Bits() {
		var line = bl.Line(li)

		var numw, numl = 0, 5
		var syml slot.Sym
		var mw float64 = 1 // mult wild
		for x := 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
				mw = 2
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl > 0 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl*mw > payw {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = g.M
			}
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payl,
				Mult: mw * mm,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = g.M
			}
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payw,
				Mult: mm,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var pay, fs = ScatPay[count-1], 0
		if count >= 3 {
			fs = ScatRand[rand.N(len(ScatRand))]
		}
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel.Num()) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	if g.FS == 0 {
		var _, reels = FindReels(mrtp)
		screen.Spin(reels)
	} else {
		screen.Spin(&ReelsBon)
	}
}

func (g *Game) Apply(screen slot.Screen, wins slot.Wins) {
	if g.FS > 0 {
		g.Gain += wins.Gain()
	} else {
		g.Gain = wins.Gain()
	}

	if g.FS > 0 {
		g.FS--
		if g.FS == 0 {
			g.M = 0 // no multiplier on regular games
		}
	} else { // free spins can not be nested
		for _, wi := range wins {
			if wi.Free > 0 {
				g.FS = wi.Free
				if rand.Float64() <= Pfs {
					g.M = 2
				} else {
					g.M = 1
				}
			}
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetSel(sel slot.Bitset) error {
	var mask slot.Bitset = (1<<len(bl) - 1) << 1
	if sel == 0 {
		return slot.ErrNoLineset
	}
	if sel&^mask != 0 {
		return slot.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return slot.ErrNoFeature
	}
	g.Sel = sel
	return nil
}
