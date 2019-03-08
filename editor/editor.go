package editor

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

//  EditWalker check have  which node need change
type EditWalker struct {
	process     []nodeProcess.NodePrecess
	currentNode node.Node
	beforeNode  node.Node
	modifyNode  map[node.Node]node.Node
	addNode     []node.Node
}

func NewEditWalker(process []nodeProcess.NodePrecess) *EditWalker {
	return &EditWalker{process: process, modifyNode: map[node.Node]node.Node{}}
}

func (e *EditWalker) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		panic("error node")
	}
	for _, p := range e.process {
		if p.Check(n, e.currentNode) {
			add, rep := p.Precess(n)
			if rep != nil {
				e.modifyNode[n] = rep
				e.addNode = append(e.addNode, add...)
				break
			}
		}
	}
	//e.currentNode = append(e.currentNode, n)
	return true
}

func (e *EditWalker) LeaveNode(w walker.Walkable) {
	//e.currentNode = e.currentNode[:len(e.currentNode)-1]
	// do nothing
}

func (e *EditWalker) EnterChildNode(key string, w walker.Walkable) {
	e.beforeNode = e.currentNode
	e.currentNode = w.(node.Node)

}

func (e *EditWalker) LeaveChildNode(key string, w walker.Walkable) {
	e.currentNode = e.beforeNode
}

func (e *EditWalker) EnterChildList(key string, w walker.Walkable) {
	e.beforeNode = e.currentNode
	e.currentNode = w.(node.Node)
}

func (e *EditWalker) LeaveChildList(key string, w walker.Walkable) {
	e.currentNode = e.beforeNode
}
