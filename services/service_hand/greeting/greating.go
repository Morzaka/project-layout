package greeting

type handLeve string

const (
	High handLeve = "high"
	Low  handLeve = "low"
)

type Hand struct {
	WhichOne  string
	Action    string
	IsHolding bool
	Item      string
	HandLeve  handLeve
}

func NewHand(which string) *Hand {
	return &Hand{
		WhichOne: which,
	}
}

func (h *Hand) Greet() {
	h.HandLeve = High
	h.Action = "waving"
}
