package app

import (
	"fmt"
	"strings"

	"github.com/adapterik/goco"
	"github.com/adapterik/goco/components/button"
	"github.com/adapterik/goco/components/input"
	"github.com/adapterik/goco/components/text"

	"honnef.co/go/js/dom"
)

type Component struct {
	// Connector
	// ValueConnector
	parentNode dom.Element
	rootNode   *dom.HTMLDivElement

	input1  *input.Input
	input2  *input.Input
	text1   *text.Component
	button1 *button.Button
}

func New(parent dom.Element) *Component {
	self := &Component{parentNode: parent}
	self.Attach()
	return self
}

// SetStyle gives access to raw styles for components.
// It is likely that we will prefer to establish interfaces
// for common styling models. Components should probably not
// be arbitrarily stylable, and families of components may have
// specific style constraints.
func (i *Component) SetStyle(name string, value string) {
	i.rootNode.Style().SetProperty(name, value, "")
}

// func selectDivElement(from Selectable, selector string) (*dom.HTMLDivElement, error) {
// 	n := from.QuerySelector(selector)

// 	if n == nil {
// 		return nil, errors.New("Sorry could not find dom node")
// 	}
// 	switch n.(type) {
// 	default:
// 		return nil, fmt.Errorf("Not an Element! %s", selector)
// 	case *dom.HTMLDivElement:
// 		return n.(*dom.HTMLDivElement), nil
// 	}
// }

// func selectInputElement(from Selectable, selector string) (*dom.HTMLInputElement, error) {
// 	n := from.QuerySelector(selector)

// 	if n == nil {
// 		return nil, errors.New("Sorry could not find dom node")
// 	}
// 	switch n.(type) {
// 	default:
// 		return nil, fmt.Errorf("Not an Input Element! %s", selector)
// 	case *dom.HTMLInputElement:
// 		return n.(*dom.HTMLInputElement), nil
// 	}
// }

// func selectElement(from Selectable, selector string) (dom.Element, error) {
// 	n := from.QuerySelector(selector)

// 	if n == nil {
// 		return nil, errors.New("Sorry could not find dom node")
// 	}

// 	return n, nil
// }

func findComponentNode(from dom.Element, name string) *dom.HTMLDivElement {
	selector := strings.Join([]string{"[data-name=\"", name, "\"]"}, "")
	n := from.QuerySelector(selector)
	return n.(*dom.HTMLDivElement)
}

func (i *Component) FindComponentNode(name string) *dom.HTMLDivElement {
	selector := strings.Join([]string{"[data-name=\"", name, "\"]"}, "")
	n := i.rootNode.QuerySelector(selector)
	return n.(*dom.HTMLDivElement)
}

func (i *Component) SetComponent(name string, comp goco.Component) {
	i.FindComponentNode(name).AppendChild(comp.Element())
}

// Attach this component to a parent node identified by the given
// selector string.
// TODO: This needs a dom context, or more simply a node.
func (i *Component) Attach() error {
	i.parentNode.SetInnerHTML(T_markup())
	i.rootNode = findComponentNode(i.parentNode, "root")
	// Create input Component
	i.input1 = input.New()
	i.SetComponent("input1", i.input1)

	// Create another input Component
	i.input2 = input.New()
	i.SetComponent("input2", i.input2)

	// Create a text Component
	i.text1 = text.New()
	i.SetComponent("text1", i.text1)

	// create a button component
	i.button1 = button.NewButton()
	fmt.Println("hi")

	i.SetComponent("button1", i.button1)

	// set up the button.
	i.button1.SetBackgroundColor("yellow")
	i.button1.SetColor("red")
	i.button1.SetHoverBackgroundColor("green")
	i.button1.SetLabel("Hi, from a button")
	i.button1.OnClick(func(ev dom.Event) {
		i.input2.SetValue("well, hello")
	})

	// Listen for changes to the value for input1.
	i.input1.OnValue(func(newValue string) {
		i.text1.SetText(newValue)
	})

	// Initialize values
	i.input1.SetValue("well, maybe it will work 1")
	i.input2.SetValue("well, maybe it will work 2")

	return nil
}

func (i *Component) Detach() {

}
