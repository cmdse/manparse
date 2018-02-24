package reporter

import "github.com/emirpasic/gods/stacks/arraystack"

type ContextQueue struct {
	*arraystack.Stack
}

func NewContextQueue() *ContextQueue {
	return &ContextQueue{
		arraystack.New(),
	}
}

func interfaceToParseContext(context interface{}, ok bool) (ParseContext, bool) {
	if !ok {
		return "", false
	}
	if context, ok := context.(ParseContext); ok {
		return context, true
	} else {
		panic("unexpected type in ContextQueue#Pop")
	}
}

func (cqueue *ContextQueue) Pop() (ParseContext, bool) {
	context, ok := cqueue.Stack.Pop()
	return interfaceToParseContext(context, ok)
}

func (cqueue *ContextQueue) Peek() (ParseContext, bool) {
	context, ok := cqueue.Stack.Peek()
	return interfaceToParseContext(context, ok)
}

func (cqueue *ContextQueue) Push(context ParseContext) {
	cqueue.Stack.Push(context)
}

func (cqueue *ContextQueue) Values() []ParseContext {
	values := cqueue.Stack.Values()
	contexts := make([]ParseContext, len(values))
	for i, val := range values {
		if context, ok := val.(ParseContext); ok {
			contexts[len(values)-1-i] = context
		}
	}
	return contexts
}
