//go:build !prod || full || agt

package happysanta50

import (
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "50 Happy Santa"},
		{Prov: "AGT", Name: "40 Bigfoot"}, // see: https://demo.agtsoftware.com/games/agt/bigfoot40
	},
	GP: game.GPsel |
		game.GPfgno |
		game.GPscat |
		game.GPwild,
	SX:  5,
	SY:  4,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		var aid = util.ToID(ga.Prov + "/" + ga.Name)
		game.ScanFactory[aid] = CalcStat
		game.GameFactory[aid] = func() any { return NewGame() }
	}
}
