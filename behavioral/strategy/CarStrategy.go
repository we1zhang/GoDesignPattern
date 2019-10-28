package strategy

type CarStrategy struct{}

func (c *CarStrategy) Travel() string {
	return "i travel by car"
}
