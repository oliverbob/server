//go:build !prod || full || keno

package kenofast

import (
	"context"

	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Keno Fast"},
	},
	GP:  0,
	SX:  80,
	SY:  0,
	SN:  0,
	LN:  0,
	BN:  0,
	RTP: []float64{95.616967},
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		var aid = util.ToID(ga.Prov + "/" + ga.Name)
		game.ScanFactory[aid] = func(ctx context.Context, mrtp float64) float64 {
			return Paytable.CalcStat(ctx)
		}
		game.GameFactory[aid] = func() any { return NewGame() }
	}
}
