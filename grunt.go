package grunt

import (
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type Grunt struct {
	gxui.AdapterBase
	node
}

func (a *Grunt) Size(t gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 18}
}

func (root *node) Init() *node {
	root.addLess()

	autoprefixer := root.add("autoprefixer")
	autoprefixer.add("options")

	cssmin := root.add("cssmin")
	cssmin.add("options")

	concat := root.add("concat")
	concat.add("options")

	ngAnnotate := root.add("ngAnnotate")
	ngAnnotate.add("options")

	uglify := root.add("uglify")
	uglify.add("options")

	root.add("watch")
	return root
}

func CreateGrunt(theme gxui.Theme) gxui.LinearLayout {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	grunt := &Grunt{}

	// hook up node changed function to the adapter OnDataChanged event.
	grunt.changed = func() { grunt.DataChanged(false) }

	nodeRoot := grunt.add("grunt").Init()

	tree := theme.CreateTree()
	tree.SetAdapter(grunt)
	tree.Select(nodeRoot)
	tree.Show(tree.Selected())

	layout.AddChild(tree)

	row := theme.CreateLinearLayout()
	row.SetDirection(gxui.LeftToRight)
	layout.AddChild(row)

	expandAll := theme.CreateButton()
	expandAll.SetText("Expand All")
	expandAll.OnClick(func(gxui.MouseEvent) { tree.ExpandAll() })
	row.AddChild(expandAll)

	collapseAll := theme.CreateButton()
	collapseAll.SetText("Collapse All")
	collapseAll.OnClick(func(gxui.MouseEvent) { tree.CollapseAll() })
	row.AddChild(collapseAll)

	return layout
}
