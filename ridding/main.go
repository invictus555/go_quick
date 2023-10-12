package ridding

type None struct{}

var none None

type Set struct {
	item map[string]None
}

func (s *Set) Add(str string) {
	s.item[str] = none
}

func TestifyRidding() {
}
