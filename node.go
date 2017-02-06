package grunt

import (
	"github.com/caothu159/golibs"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

// node is an implementation of gxui.TreeNode.
type node struct {
	identity string
	name     string
	changed  func()
	children []*node
}

// add appends a new child node to n with the specified name.
func (n *node) add(name string) *node {
	child := &node{
		identity: golibs.UUID(),
		name:     name,
		changed:  n.changed,
	}
	n.children = append(n.children, child)
	n.changed()
	return child
}

// Count implements gxui.TreeNodeContainer.
func (n *node) Count() int {
	return len(n.children)
}

// Create implements gxui.TreeNode.
func (n *node) Create(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	label := theme.CreateLabel()
	label.SetText(n.name)

	textbox := theme.CreateTextBox()
	textbox.SetText(n.name)
	textbox.SetPadding(math.ZeroSpacing)
	textbox.SetMargin(math.ZeroSpacing)

	addButton := theme.CreateButton()
	addButton.SetText("+")
	addButton.OnClick(func(gxui.MouseEvent) {
		n.add("new")
	})

	edit := func() {
		layout.RemoveAll()
		layout.AddChild(textbox)
		layout.AddChild(addButton)
		gxui.SetFocus(textbox)
	}

	commit := func() {
		n.name = textbox.Text()
		label.SetText(n.name)
		layout.RemoveAll()
		layout.AddChild(label)
		layout.AddChild(addButton)
	}

	// When the user clicks the label, replace it with an editable text-box
	label.OnClick(func(gxui.MouseEvent) { edit() })

	// When the text-box loses focus, replace it with a label again.
	textbox.OnLostFocus(commit)

	layout.AddChild(label)
	layout.AddChild(addButton)
	return layout
}

// Item implements gxui.TreeNode.
func (n *node) Item() gxui.AdapterItem {
	return n.identity
}

// ItemIndex implements gxui.TreeNodeContainer.
func (n *node) ItemIndex(identity gxui.AdapterItem) int {
	for i, c := range n.children {
		if c.identity == identity || c.ItemIndex(identity) >= 0 {
			return i
		}
	}
	return -1
}

// NodeAt implements gxui.TreeNodeContainer.
func (n *node) NodeAt(index int) gxui.TreeNode {
	return n.children[index]
}
