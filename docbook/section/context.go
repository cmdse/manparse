package section

import (
	"encoding/xml"
	"fmt"

	"github.com/oleiade/lane"
)

type Context struct {
	fifo *lane.Queue
}

func NewContext() *Context {
	return &Context{
		lane.NewQueue(),
	}
}

func (context *Context) Put(element xml.StartElement) {
	context.fifo.Append(element)
}

func (context *Context) HandleLast(endElement xml.EndElement) (err error) {
	last := context.fifo.Last()
	if lastElement, ok := last.(xml.StartElement); ok {
		if endElement.Name.Local != lastElement.Name.Local {
			err = fmt.Errorf("XML tags are not matching, opening tag is %v, closing tag is %v", lastElement.Name.Local, endElement.Name.Local)
		}
	} else {
		err = fmt.Errorf("fifo should only contain xml.StartElement but found %t", last)
	}
	if err == nil {
		context.fifo.Pop()
	}
	return err
}
