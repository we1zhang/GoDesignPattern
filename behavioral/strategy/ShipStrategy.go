package strategy

type ShipStrategy struct {
}

func (s *ShipStrategy) Travel() string {
	return "i travel by ship"
}
