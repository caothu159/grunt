package mixins

type Grunt struct {
	GetLess mixins.GetLess
}

func CreateGrunt() Grunt {
	return Grunt
}
