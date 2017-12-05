package main

import (
	"github.com/adapterik/goco-example-app/app"
	"honnef.co/go/js/dom"
)

func main() {
	root := dom.GetWindow().Document().QuerySelector("#app")
	_ = app.New(root.(*dom.HTMLDivElement))
}
