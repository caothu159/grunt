package mixins

type Less struct {
	option []Option
}

func GetLess() Less {
	var less Less
	return less
}
