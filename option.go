package grunt

import (
	"github.com/caothu159/golibs"
)

// add appends a new child node to n with the specified name.
func (n *node) addOption() *node {
	child := &node{
		identity: golibs.UUID(),
		name:     "options",
		changed:  n.changed,
	}
	n.changed()
	n.children = append(n.children, child)
	return child
}
