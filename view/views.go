package view

import "github.com/ungerik/go-start/utils"

///////////////////////////////////////////////////////////////////////////////
// Views

// Views implements the View interface for a slice of views.
// The views of the slice are the child views.
type Views []View

func (self Views) Init(thisView View) {
	self.IterateChildren(func(parent View, child View) (next bool) {
		if child != nil {
			child.Init(child)
		}
		return true
	})
}

func (self Views) OnRemove() {
	self.IterateChildren(func(parent View, child View) (next bool) {
		if child != nil {
			child.OnRemove()
		}
		return true
	})
}

func (self Views) ID() string {
	return ""
}

// Does not iterate nil children
func (self Views) IterateChildren(callback IterateChildrenCallback) {
	for _, view := range self {
		if view != nil && !callback(self, view) {
			return
		}
	}
}

func (self Views) Render(context *Context, writer *utils.XMLWriter) (err error) {
	for _, view := range self {
		if view != nil {
			err = view.Render(context, writer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}