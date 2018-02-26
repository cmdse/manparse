package reporter

import (
	"bytes"

	"github.com/fatih/color"
)

var bold = color.New(color.FgBlack).Add(color.Bold)

// ParseContext is the scope in which reported events occurred.
type ParseContext struct {
	value         string
	childReports  Reports
	childContexts []*ParseContext
	parentContext *ParseContext
	level         int
}

func (context *ParseContext) hasReports() bool {
	return len(context.childReports) > 0
}

func NewParseContext(message string, parent *ParseContext, level int) *ParseContext {
	return &ParseContext{
		message,
		make(Reports, 0, 5),
		make([]*ParseContext, 0, 2),
		parent,
		level,
	}
}

func (context *ParseContext) AddChildReport(child *Report) {
	context.childReports = append(context.childReports, child)
	child.offset = context.level + 1
}

func (context *ParseContext) AddChildContext(child *ParseContext) {
	context.childContexts = append(context.childContexts, child)
}

func (context *ParseContext) hasChildren() bool {
	return len(context.childReports) > 0 || len(context.childContexts) > 0
}

func (context *ParseContext) sprintContextValue() string {
	return writeOffset(context.level, bold.Sprintf(context.value))
}

func (context *ParseContext) PrettyPrint() string {
	if context.hasChildren() {
		var buffer bytes.Buffer
		buffer.Grow(500)
		buffer.WriteString(context.sprintContextValue())
		for _, context := range context.childContexts {
			buffer.WriteString(context.PrettyPrint())
		}
		for _, report := range context.childReports {
			buffer.WriteString(writeOffset(report.offset, report.PrettyPrint()))
		}
		return buffer.String()
	}
	return ""
}
