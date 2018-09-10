package websock_v2

var C = NewClient(DefaultConfig)

type Client struct {
	*Melody
}

func NewClient(cfg Config) *Client {
	c := &Client{
		Melody: NewMelody(cfg),
	}
	return c
}
