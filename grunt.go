package grunt

import (
	"flag"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/themes/light"
)

var DefaultScaleFactor float32
var FlagTheme string

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

func appFlags() {
	flagTheme := flag.String("theme", "dark", "Theme to use {dark|light}.")
	defaultScaleFactor := flag.Float64("scaling", 1.0, "Adjusts the scaling of UI rendering")
	flag.Parse()

	DefaultScaleFactor = float32(*defaultScaleFactor)
	FlagTheme = *flagTheme
}

func appCreateTheme(driver gxui.Driver) gxui.Theme {
	if FlagTheme == "light" {
		return light.CreateTheme(driver)
	}
	return dark.CreateTheme(driver)
}

func InitTree(driver gxui.Driver) {
	appFlags()
	theme := appCreateTheme(driver)

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

	window := theme.CreateWindow(800, 600, "Code tools")
	window.SetScale(DefaultScaleFactor)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func CreateWindow() {
	gl.StartDriver(InitTree)
}
