package beetlemania

import (
	"math"

	slot "github.com/slotopol/server/game/slot"
	"github.com/slotopol/server/util"
)

var ReelsReg88 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 6, 8, 2, 8, 6, 8, 1, 9, 6, 2, 8, 4, 8, 6, 3, 7, 5},
	{9, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 9, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 8, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 10, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 3, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 4, 8, 9, 6, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg90 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 7, 4, 6, 8, 2, 9, 6, 8, 1, 9, 6, 2, 8, 4, 8, 6, 3, 7, 5},
	{9, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 8, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 10, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 3, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 4, 8, 9, 6, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg92 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 6, 8, 2, 8, 6, 8, 1, 9, 6, 2, 8, 4, 8, 6, 3, 7, 5},
	{9, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 9, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 10, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 9, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 9, 5, 3, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 4, 8, 9, 6, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg94 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 6, 8, 2, 8, 6, 8, 1, 9, 6, 2, 8, 4, 9, 6, 3, 7, 5},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 10, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 9, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 3, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 4, 8, 9, 6, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg95 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 9, 4, 6, 8, 2, 7, 6, 8, 1, 9, 6, 2, 8, 4, 9, 6, 3, 7, 5},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 10, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 9, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 3, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 4, 8, 9, 6, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg96 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 6, 8, 2, 8, 6, 8, 1, 9, 6, 2, 8, 4, 9, 6, 3, 7, 5},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 10, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 9, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 10, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 8, 4, 8, 9, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg97 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 6, 9, 2, 7, 6, 8, 1, 9, 6, 2, 8, 4, 9, 6, 3, 7, 5},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 3, 5, 9, 1, 6, 8, 9, 7, 10, 9, 4},
	{4, 5, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 4, 5, 3, 6, 7, 11, 8, 6, 9, 5, 7, 2, 9, 7, 4, 6, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 5, 10, 9, 8, 6, 4, 8, 6, 5, 3, 8, 6, 4, 5, 6, 7, 4, 9, 5, 2, 8},
	{5, 7, 10, 9, 5, 10, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 8, 4, 8, 9, 2, 8, 6, 4, 8, 3, 9},
}

var ReelsReg100 = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 1, 7, 5, 3, 8, 4, 5, 9, 2, 7, 6, 8, 1, 9, 6, 2},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 1, 8, 7, 6, 5, 9, 1},
	{4, 9, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 7, 6, 3, 6, 7, 11, 8, 6, 5, 2, 7},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 7, 6, 2, 9, 8, 6, 9, 8, 6, 5, 3, 8, 7, 4, 6},
	{5, 7, 10, 9, 5, 2, 9, 7, 6, 1, 8, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 8, 4, 8, 9},
}

var ReelsBon = slot.Reels5x{
	{8, 6, 10, 8, 9, 1, 7, 5, 6, 9, 7, 5, 3, 8, 4, 7, 9, 2, 7, 9, 5, 7, 9, 5, 2, 8, 4, 9, 8, 7, 8, 5, 9, 7, 8, 9, 8, 6, 10, 8, 9, 7, 3, 5, 7},
	{8, 7, 10, 9, 5, 7, 1, 5, 6, 3, 9, 4, 7, 8, 2, 9, 6, 5, 4, 8, 7, 3, 5, 9, 7, 6, 8, 9, 7, 10, 9, 4, 7, 5, 6, 8, 5, 7, 10, 9, 5, 6, 7, 9, 6},
	{4, 8, 10, 7, 9, 3, 8, 11, 5, 9, 6, 1, 5, 8, 7, 5, 3, 6, 7, 11, 8, 6, 9, 8, 7, 2, 9, 7, 6, 9, 8, 6},
	{7, 5, 10, 9, 8, 5, 1, 7, 9, 2, 8, 5, 1, 7, 9, 5, 8, 9, 6, 5, 3, 8, 6, 5, 4, 6, 5, 7, 4, 5, 9, 2},
	{5, 7, 10, 9, 5, 2, 9, 7, 6, 1, 5, 5, 3, 9, 4, 7, 6, 1, 8, 5, 9, 7, 8, 9, 7, 2, 6, 7, 4, 5, 3, 9},
}

var ReelsReg88u = slot.Reels5x{
	{7, 2, 5, 7, 3, 6, 8, 9, 4, 6, 8, 1, 8, 7, 2, 6, 5, 8, 4, 8, 7, 3, 5, 7, 8, 1, 6, 9, 10, 9},
	{10, 5, 7, 4, 7, 6, 2, 7, 6, 9, 1, 5, 9, 8, 3, 8, 9, 3, 8, 5, 7, 4, 9, 9, 1, 6, 5, 8, 7, 9},
	{6, 10, 7, 8, 4, 6, 1, 5, 2, 9, 4, 9, 5, 7, 11, 6, 8, 2, 9, 3, 5, 4, 6, 7, 1, 8, 9, 11, 7, 9, 3, 5},
	{10, 5, 7, 4, 8, 7, 2, 8, 9, 3, 9, 7, 8, 4, 5, 6, 10, 5, 7, 2, 9, 8, 1, 8, 9, 4, 9, 5, 8, 7, 6, 9},
	{10, 7, 6, 4, 5, 8, 2, 5, 8, 3, 6, 5, 4, 5, 2, 6, 7, 10, 8, 9, 2, 6, 7, 1, 6, 7, 1, 8, 6, 4, 7, 9},
}

var ReelsReg90u = slot.Reels5x{
	{10, 7, 9, 2, 6, 8, 3, 7, 5, 4, 9, 8, 1, 7, 8, 3, 5, 6, 1, 8, 6, 2, 7, 6, 4, 8, 6, 7, 8, 9},
	{10, 7, 5, 3, 6, 7, 1, 8, 9, 2, 7, 8, 4, 6, 7, 1, 5, 9, 3, 5, 8, 9, 4, 5, 9, 5, 8, 9, 7, 9},
	{10, 6, 9, 2, 5, 7, 11, 6, 8, 1, 6, 9, 4, 5, 6, 10, 8, 9, 11, 7, 6, 2, 5, 9, 3, 5, 4, 9, 3, 1, 5, 6},
	{10, 5, 7, 2, 9, 8, 3, 8, 9, 4, 9, 8, 3, 5, 6, 1, 7, 5, 2, 9, 7, 3, 8, 7, 2, 8, 4, 7, 2, 8, 6, 9},
	{10, 7, 6, 3, 5, 8, 4, 6, 7, 10, 8, 9, 2, 6, 1, 7, 4, 7, 3, 5, 8, 4, 6, 2, 5, 3, 8, 2, 1, 5, 7, 9},
}

var ReelsReg92u = slot.Reels5x{
	{10, 7, 9, 2, 9, 6, 1, 9, 8, 3, 6, 5, 1, 8, 9, 4, 6, 8, 5, 6, 3, 8, 7, 4, 7, 4, 5, 2, 5, 9},
	{10, 7, 5, 2, 7, 8, 1, 6, 7, 3, 5, 8, 4, 6, 7, 1, 8, 9, 4, 5, 6, 8, 5, 3, 6, 8, 2, 9, 8, 9},
	{7, 5, 2, 6, 3, 9, 1, 5, 2, 6, 4, 5, 7, 11, 6, 9, 4, 9, 3, 7, 4, 5, 3, 9, 7, 1, 8, 6, 10, 8, 9, 11},
	{10, 5, 7, 3, 8, 6, 1, 7, 6, 10, 9, 8, 7, 2, 9, 8, 3, 7, 8, 4, 8, 9, 4, 5, 7, 8, 5, 7, 8, 7, 9, 6},
	{10, 8, 9, 2, 6, 7, 1, 9, 7, 3, 8, 4, 5, 8, 3, 7, 9, 10, 7, 6, 4, 5, 8, 1, 9, 5, 3, 7, 8, 2, 6, 7},
}

var ReelsReg94u = slot.Reels5x{
	{10, 7, 9, 3, 6, 7, 1, 8, 9, 2, 7, 5, 1, 7, 8, 4, 7, 3, 6, 4, 8, 9, 2, 5, 3, 9, 4, 9, 8, 9},
	{10, 5, 7, 2, 6, 7, 1, 8, 9, 4, 8, 5, 3, 7, 5, 1, 7, 6, 2, 5, 6, 3, 6, 5, 4, 6, 8, 6, 9, 7},
	{10, 8, 7, 4, 8, 9, 11, 7, 6, 3, 6, 8, 1, 5, 9, 2, 5, 7, 11, 6, 8, 10, 9, 5, 2, 5, 4, 6, 5, 1, 5, 9},
	{10, 5, 7, 8, 5, 6, 1, 7, 5, 3, 8, 6, 4, 8, 6, 7, 4, 9, 8, 7, 9, 2, 8, 9, 5, 4, 8, 9, 8, 9, 6, 9},
	{10, 7, 6, 2, 6, 7, 1, 9, 7, 3, 9, 8, 3, 6, 7, 1, 9, 5, 4, 8, 6, 2, 8, 9, 6, 4, 8, 9, 5, 8, 7, 9},
}

var ReelsReg96u = slot.Reels5x{
	{10, 7, 8, 2, 8, 7, 1, 8, 7, 3, 5, 6, 4, 8, 7, 2, 6, 8, 4, 6, 5, 1, 9, 5, 6, 4, 9, 5, 8, 6},
	{10, 5, 7, 3, 9, 8, 1, 6, 9, 2, 8, 9, 10, 9, 8, 4, 8, 6, 3, 9, 6, 4, 9, 8, 1, 5, 9, 3, 9, 7},
	{10, 9, 7, 2, 8, 9, 11, 7, 6, 3, 5, 6, 4, 5, 7, 11, 6, 8, 2, 7, 5, 3, 6, 5, 1, 7, 5, 3, 7, 1, 5, 9},
	{10, 5, 7, 3, 8, 6, 1, 7, 8, 2, 8, 9, 4, 6, 9, 10, 8, 9, 4, 8, 7, 3, 5, 9, 2, 8, 7, 4, 9, 8, 9, 7},
	{10, 5, 7, 9, 6, 7, 1, 9, 7, 2, 5, 7, 3, 8, 6, 7, 1, 9, 7, 9, 6, 5, 8, 3, 8, 5, 4, 9, 5, 6, 7, 5},
}

var ReelsBonu = slot.Reels5x{
	{10, 9, 5, 7, 6, 5, 7, 8, 9, 2, 7, 8, 6, 7, 10, 8, 9, 7, 3, 9, 8, 2, 9, 8, 3, 6, 7, 8, 9, 8, 10, 7, 9, 4, 5, 6, 9, 7, 4, 8, 9, 1, 7, 9, 6},
	{7, 5, 9, 4, 5, 6, 7, 10, 8, 9, 2, 7, 6, 7, 5, 10, 8, 7, 4, 6, 7, 3, 8, 9, 5, 4, 5, 6, 8, 7, 10, 6, 7, 8, 3, 9, 7, 8, 4, 9, 5, 6, 7, 9, 1},
	{8, 5, 3, 9, 5, 8, 6, 10, 9, 7, 2, 8, 9, 11, 7, 6, 5, 8, 3, 9, 8, 4, 5, 7, 11, 6, 8, 9, 5, 1, 8, 7},
	{8, 6, 9, 4, 5, 6, 10, 7, 5, 2, 8, 9, 3, 5, 6, 1, 5, 7, 8, 4, 5, 6, 2, 9, 5, 3, 6, 9, 1, 5, 7, 6, 3},
	{10, 9, 7, 4, 5, 6, 7, 1, 8, 9, 2, 6, 3, 7, 8, 3, 9, 8, 4, 5, 6, 2, 9, 8, 6, 3, 7, 9, 1, 7, 6, 7},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	/*"88":   &ReelsReg88,
	"90":   &ReelsReg90,
	"92":   &ReelsReg92,
	"94":   &ReelsReg94,
	"95":   &ReelsReg95,
	"96":   &ReelsReg96,
	"97":   &ReelsReg97,
	"100":  &ReelsReg100,
	"bon":  &ReelsBon,*/
	88: &ReelsReg88u,
	90: &ReelsReg90u,
	92: &ReelsReg92u,
	94: &ReelsReg94u,
	96: &ReelsReg96u,
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
	{0, 10, 80, 1000, 5000}, //  1 bee
	{0, 5, 30, 200, 1000},   //  2 snail
	{0, 5, 25, 100, 500},    //  3 fly
	{0, 5, 15, 65, 250},     //  4 worm
	{0, 0, 10, 40, 200},     //  5 ace
	{0, 0, 10, 40, 200},     //  6 king
	{0, 0, 5, 20, 100},      //  7 queen
	{0, 0, 5, 20, 100},      //  8 jack
	{0, 0, 5, 20, 100},      //  9 ten
	{0, 0, 0, 0, 0},         // 10 note
	{0, 0, 0, 0, 0},         // 11 jazzbee
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 2, 15, 50} // 10 note

const (
	jbonus = 1 // jazzbee freespins bonus
)

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [11][5]int{
	{0, 0, 0, 0, 0}, //  1 bee
	{0, 0, 0, 0, 0}, //  2 snail
	{0, 0, 0, 0, 0}, //  3 fly
	{0, 0, 0, 0, 0}, //  4 worm
	{0, 0, 0, 0, 0}, //  5 ace
	{0, 0, 0, 0, 0}, //  6 king
	{0, 0, 0, 0, 0}, //  7 queen
	{0, 0, 0, 0, 0}, //  8 jack
	{0, 0, 0, 0, 0}, //  9 ten
	{0, 0, 0, 0, 0}, // 10 note
	{0, 0, 0, 0, 0}, // 11 jazzbee
}

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: util.MakeBitNum(5, 1),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild, scat = 1, 10
const jazz = 11

var bl = slot.BetLinesNvm10

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := range g.Sel.Bits() {
		var line = bl.Line(li)

		/*var numw, numl int
		var syml slot.Sym
		for x := 1; x <= 5; x++ {
			var symx = screen.Pos(x, line)
			if symx == wild {
				if syml == 0 {
					numw = x
				} else {
					numl = x
				}
			} else if symx == scat || symx == jazz {
				break
			} else if syml == 0 {
				syml = symx
				numl = x
			} else if symx == syml {
				numl = x
			} else {
				break
			}
		}*/

		var numw, numl int
		var syml slot.Sym
		for x := 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx != wild {
				if sx != scat && sx != jazz {
					syml = sx
					numl = numw + 1
					for x := numl + 1; x <= 5; x++ {
						var sx = screen.Pos(x, line)
						if sx == syml || sx == wild {
							numl++
						} else {
							break
						}
					}
				}
				break
			}
			numw++
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl > payw {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payl,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payw,
				Mult: 1,
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
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if g.FS > 0 {
		var y int
		if screen.At(3, 1) == jazz {
			y = 1
		} else if screen.At(3, 1) == jazz {
			y = 2
		} else if screen.At(3, 3) == jazz {
			y = 3
		} else {
			return // ignore scatters on freespins
		}
		var xy = slot.NewLine5x()
		xy.Set(3, y)
		*wins = append(*wins, slot.WinItem{
			Mult: 1,
			Sym:  jazz,
			Num:  1,
			XY:   xy,
			BID:  jbonus,
		})
		return
	}

	if count := screen.ScatNum(scat); count >= 3 {
		var pay = ScatPay[count-1]
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel.Num()) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPosCont(scat),
			Free: 10,
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	if g.FS == 0 {
		var _, reels = FindReels(mrtp)
		screen.Spin(reels)
	} else {
		screen.Spin(&ReelsBonu)
	}
}

func (g *Game) Spawn(screen slot.Screen, wins slot.Wins) {
	for i, wi := range wins {
		switch wi.BID {
		case jbonus:
			wins[i].Pay = min(g.Gain, 100_000*g.Bet)
		}
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
