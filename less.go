package grunt

import (
	"github.com/caothu159/golibs"
)

// add appends a new child node to n with the specified name.
func (n *node) addLess() *node {
	child := &node{
		identity: golibs.UUID(),
		name:     "less",
		changed:  n.changed,
	}
	child.addOption()
	n.children = append(n.children, child)
	n.changed()
	return child
}
