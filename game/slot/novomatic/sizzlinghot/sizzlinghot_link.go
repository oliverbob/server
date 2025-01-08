//go:build !prod || full || novomatic

package sizzlinghot

import (
	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "Novomatic", Name: "Sizzling Hot"},        // see: https://www.slotsmate.com/software/novomatic/novomatic-sizzling-hot
		{Prov: "Novomatic", Name: "Sizzling Hot Deluxe"}, // see: https://www.slotsmate.com/software/novomatic/sizzling-hot-deluxe
		{Prov: "Novomatic", Name: "Age of Heroes"},       // see: https://www.slotsmate.com/software/novomatic/age-of-heroes
		{Prov: "Novomatic", Name: "Hot Cubes"},           // see: https://www.slotsmate.com/software/novomatic/hot-cubes
	},
	GP: game.GPfgno |
		game.GPscat,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	Info.SetupFactory(func() any { return NewGame() }, CalcStat)
}
