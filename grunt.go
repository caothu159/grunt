package grunt

import (
	"github.com/caothu159/grunt/mixins"
)

type Grunt interface {
	GetLess() Less
}

func CreateGrunt() Grunt {
	return mixins.CreateGrunt()
}
