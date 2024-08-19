package dolphinspearl

import (
	"math"

	"github.com/slotopol/server/game"
)

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 39.644(lined) + 11.288(scatter) = 50.932120%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 50.932(sym) + 0.057217*618.14(fg) = 86.300723%
var ReelsReg86 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 5, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 11, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 5, 12, 7, 3, 11, 9, 5, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 41.407(lined) + 11.288(scatter) = 52.694878%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 52.695(sym) + 0.057217*618.14(fg) = 88.063480%
var ReelsReg88 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 11, 12, 4, 8, 10, 5, 7, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 11, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 5, 12, 7, 3, 11, 10, 5, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 43.709(lined) + 11.288(scatter) = 54.996591%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 54.997(sym) + 0.057217*618.14(fg) = 90.365193%
var ReelsReg90 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 5, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 10, 8, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 7, 10, 6, 8, 7, 5, 12, 4, 9, 11},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 10, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 10, 11, 2, 7, 6, 12, 8, 3, 11, 9, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 45.535(lined) + 11.288(scatter) = 56.822392%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 56.822(sym) + 0.057217*618.14(fg) = 92.190994%
var ReelsReg92 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 10, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 8, 9, 2, 8, 6, 10, 4, 11, 12, 2, 10, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 6, 12, 7, 3, 11, 9, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 47.58(lined) + 11.288(scatter) = 58.867909%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 58.868(sym) + 0.057217*618.14(fg) = 94.236512%
var ReelsReg94 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 8, 9, 2, 8, 6, 10, 4, 7, 12, 2, 10, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 6, 12, 7, 3, 11, 10, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 48.477(lined) + 11.288(scatter) = 59.764586%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 59.765(sym) + 0.057217*618.14(fg) = 95.133189%
var ReelsReg95 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 12, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 49.522(lined) + 11.288(scatter) = 60.809506%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 60.81(sym) + 0.057217*618.14(fg) = 96.178109%
var ReelsReg96 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 50.448(lined) + 11.288(scatter) = 61.735386%
// free games 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 61.735(sym) + 0.057217*618.14(fg) = 97.103989%
var ReelsReg97 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 6, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 9, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 32, 32, 32, 32], total reshuffles 31457280
// symbols: 54.251(lined) + 18.01(scatter) = 72.261626%
// free games 3489480, q = 0.11093, sq = 1/(1-q) = 1.124768
// free games frequency: 1/135.22
// RTP = 72.262(sym) + 0.11093*618.14(fg) = 140.830831%
var ReelsReg141 = game.Reels5x{
	{6, 11, 1, 9, 12, 8, 6, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 10, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 7, 4},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 9, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 11, 9},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 8, 6, 7, 2, 9, 12, 10},
}

// reels lengths [33, 36, 36, 36, 36], total reshuffles 55427328
// symbols: 226.19(lined) + 144.15(scatter) = 370.339985%
// free games 22219920, q = 0.40088, sq = 1/(1-q) = 1.66913
// free games frequency: 1/37.417
// RTP = sq*rtp(sym) = 1.6691*370.34 = 618.143873%
var ReelsBon = game.Reels5x{
	{6, 11, 1, 9, 12, 5, 11, 8, 2, 7, 6, 10, 12, 4, 7, 9, 13, 10, 9, 3, 7, 5, 10, 9, 2, 8, 11, 10, 3, 8, 11, 4, 7},
	{5, 8, 1, 7, 8, 10, 2, 9, 4, 11, 5, 8, 1, 7, 8, 11, 13, 12, 4, 10, 7, 3, 11, 6, 10, 4, 6, 11, 13, 12, 5, 9, 6, 11, 9, 10},
	{11, 9, 1, 11, 4, 7, 8, 2, 6, 11, 9, 1, 11, 4, 8, 12, 13, 9, 8, 2, 12, 10, 5, 12, 3, 6, 7, 12, 13, 9, 10, 5, 3, 7, 5, 10},
	{8, 11, 1, 9, 3, 7, 10, 4, 7, 2, 8, 6, 10, 4, 5, 11, 13, 8, 6, 12, 11, 5, 7, 9, 10, 2, 5, 11, 13, 8, 6, 9, 3, 12, 7, 4},
	{8, 7, 1, 8, 4, 9, 5, 12, 7, 6, 11, 3, 7, 8, 2, 12, 13, 5, 10, 7, 8, 9, 6, 10, 4, 9, 2, 12, 13, 5, 10, 4, 11, 6, 3, 11},
}

// Map with available reels.
var reelsmap = map[float64]*game.Reels5x{
	86.300723:  &ReelsReg86,
	88.063480:  &ReelsReg88,
	90.365193:  &ReelsReg90,
	92.190994:  &ReelsReg92,
	94.236512:  &ReelsReg94,
	95.133189:  &ReelsReg95,
	96.178109:  &ReelsReg96,
	97.103989:  &ReelsReg97,
	140.830831: &ReelsReg141,
}

func FindReels(mrtp float64) (rtp float64, reels game.Reels) {
	for p, r := range reelsmap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}

// Lined payment.
var LinePay = [13][5]float64{
	{0, 10, 250, 2500, 9000}, //  1 dolphin
	{0, 2, 25, 125, 750},     //  2 stingray
	{0, 2, 25, 125, 750},     //  3 crab
	{0, 0, 20, 100, 400},     //  4 seahorse
	{0, 0, 15, 75, 250},      //  5 pterois
	{0, 0, 15, 75, 250},      //  6 angelfish
	{0, 0, 10, 50, 125},      //  7 ace
	{0, 0, 10, 50, 125},      //  8 king
	{0, 0, 5, 25, 100},       //  9 queen
	{0, 0, 5, 25, 100},       // 10 jack
	{0, 0, 5, 25, 100},       // 11 ten
	{0, 2, 5, 25, 100},       // 12 nine
	{0, 0, 0, 0, 0},          // 13 pearl
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 5, 20, 500} // 13 pearl

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 15, 15, 15} // 13 pearl

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [13][5]int{
	{0, 0, 0, 0, 0}, //  1 dolphin
	{0, 0, 0, 0, 0}, //  2 stingray
	{0, 0, 0, 0, 0}, //  3 crab
	{0, 0, 0, 0, 0}, //  4 seahorse
	{0, 0, 0, 0, 0}, //  5 pterois
	{0, 0, 0, 0, 0}, //  6 angelfish
	{0, 0, 0, 0, 0}, //  7 ace
	{0, 0, 0, 0, 0}, //  8 king
	{0, 0, 0, 0, 0}, //  9 queen
	{0, 0, 0, 0, 0}, // 10 jack
	{0, 0, 0, 0, 0}, // 11 ten
	{0, 0, 0, 0, 0}, // 12 nine
	{0, 0, 0, 0, 0}, // 13 pearl
}

type Game struct {
	game.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame(rtp float64) *Game {
	return &Game{
		Slot5x3: game.Slot5x3{
			RTP: rtp,
			SBL: game.MakeBitNum(5),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild, scat = 1, 13

var bl = game.BetLinesNvm10

func (g *Game) Scanner(screen game.Screen, wins *game.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen game.Screen, wins *game.Wins) {
	for li := g.SBL.Next(0); li != 0; li = g.SBL.Next(li) {
		var line = bl.Line(li)

		var numw, numl = 0, 5
		var syml game.Sym
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
				mm = 3
			}
			*wins = append(*wins, game.WinItem{
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
				mm = 3
			}
			*wins = append(*wins, game.WinItem{
				Pay:  g.Bet * payw,
				Mult: mm,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
				Jack: Jackpot[wild-1][numw-1],
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen game.Screen, wins *game.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var mm float64 = 1 // mult mode
		if g.FS > 0 {
			mm = 3
		}
		var pay, fs = ScatPay[count-1], ScatFreespin[count-1]
		*wins = append(*wins, game.WinItem{
			Pay:  g.Bet * float64(g.SBL.Num()) * pay,
			Mult: mm,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}
}

func (g *Game) Spin(screen game.Screen) {
	if g.FS == 0 {
		var _, reels = FindReels(g.RTP)
		screen.Spin(reels)
	} else {
		screen.Spin(&ReelsBon)
	}
}

func (g *Game) Apply(screen game.Screen, wins game.Wins) {
	if g.FS > 0 {
		g.Gain += wins.Gain()
	} else {
		g.Gain = wins.Gain()
	}

	if g.FS > 0 {
		g.FS--
	}
	for _, wi := range wins {
		if wi.Free > 0 {
			g.FS += wi.Free
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetLines(sbl game.Bitset) error {
	var mask game.Bitset = (1<<len(bl) - 1) << 1
	if sbl == 0 {
		return game.ErrNoLineset
	}
	if sbl&^mask != 0 {
		return game.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return game.ErrNoFeature
	}
	g.SBL = sbl
	return nil
}
