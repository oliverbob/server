package columbus

import (
	slot "github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [10][5]float64{
	{0, 10, 100, 1000, 5000}, //  1 columbus
	{0, 5, 50, 200, 1000},    //  2 spain
	{0, 5, 25, 100, 500},     //  3 necklace
	{0, 5, 15, 75, 250},      //  4 sextant
	{0, 0, 10, 40, 150},      //  5 ace
	{0, 0, 10, 40, 150},      //  6 king
	{0, 0, 10, 40, 150},      //  7 queen
	{0, 0, 5, 20, 100},       //  8 jack
	{0, 0, 5, 20, 100},       //  9 ten
	{0, 0, 0, 0, 0},          // 10 ship
}

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [10][5]int{
	{0, 0, 0, 0, 0}, //  1 columbus
	{0, 0, 0, 0, 0}, //  2 spain
	{0, 0, 0, 0, 0}, //  3 necklace
	{0, 0, 0, 0, 0}, //  4 sextant
	{0, 0, 0, 0, 0}, //  5 ace
	{0, 0, 0, 0, 0}, //  6 king
	{0, 0, 0, 0, 0}, //  7 queen
	{0, 0, 0, 0, 0}, //  8 jack
	{0, 0, 0, 0, 0}, //  9 ten
	{0, 0, 0, 0, 0}, // 10 ship
}

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: slot.MakeBitNum(5, 1),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild, scat = 1, 10

var bl = slot.BetLinesNvm10

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	var wbon slot.Sym
	if g.FS > 0 {
		wbon = scat
	}

	for li := g.Sel.Next(0); li != -1; li = g.Sel.Next(li) {
		var line = bl[li-1]

		var numw, numl slot.Pos = 0, 5
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild || sx == wbon {
				if syml == 0 {
					numw = x
				}
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
	if count := screen.ScatNumOdd(scat); count >= 3 {
		*wins = append(*wins, slot.WinItem{
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPosOdd(scat),
			Free: 10,
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	if g.FS == 0 {
		var _, reels = slot.FindReels(ReelsMap, mrtp)
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
	if sel.IsZero() {
		return slot.ErrNoLineset
	}
	if bs := sel; !bs.AndNot(slot.MakeBitNum(len(bl), 1)).IsZero() {
		return slot.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return slot.ErrDisabled
	}
	g.Sel = sel
	return nil
}
