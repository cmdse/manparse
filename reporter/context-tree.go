package reporter

// ParseContext is the scope in which reported events occurred.
type ContextTree struct {
	root        *ParseContext
	lastContext *ParseContext
}

func NewContextTree(root string) *ContextTree {
	context := NewParseContext(root, nil, 0)
	return &ContextTree{
		context,
		context,
	}
}

func (tree *ContextTree) AddContext(message string) *ParseContext {
	newContext := NewParseContext(message, tree.lastContext, tree.lastContext.level+1)
	tree.lastContext.AddChildContext(newContext)
	tree.lastContext = newContext
	return newContext
}

func (tree *ContextTree) ReleaseContext() {
	tree.lastContext = tree.lastContext.parentContext
}

func (tree *ContextTree) PrettyPrint() string {
	return tree.root.PrettyPrint()
}
