package editor

import (
	"github.com/blue-bird1/ConfusedPHP/nodeProcess"
	"github.com/blue-bird1/ConfusedPHP/util"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

//  EditWalker check have  which node need change
type EditWalker struct {
	process        []nodeProcess.NodePrecess
	beforeNode     util.EnterNode
	beforeNodeList []util.EnterNode
	modifyNode     map[node.Node]node.Node
	addNode        []node.Node
}

func (e *EditWalker) ModifyNode() map[node.Node]node.Node {
	return e.modifyNode
}

func (e *EditWalker) SetModifyNode(modifyNode map[node.Node]node.Node) {
	e.modifyNode = modifyNode
}

func (e *EditWalker) AddNode() []node.Node {
	return e.addNode
}

func (e *EditWalker) SetAddNode(addNode []node.Node) {
	e.addNode = addNode
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
		if p.Check(n, e.beforeNode) {
			add, rep := p.Precess(n)
			if rep != nil {
				e.modifyNode[n] = rep
				e.addNode = append(e.addNode, add...)
				break
			}
		}
	}

	return true
}

func (e *EditWalker) LeaveNode(w walker.Walkable) {
	// do nothing
}

func (e *EditWalker) EnterChildNode(key string, w walker.Walkable) {
	e.beforeNode.Key = key
	e.beforeNode.Node = w.(node.Node)
	e.beforeNodeList = append(e.beforeNodeList, e.beforeNode)

}

func (e *EditWalker) LeaveChildNode(key string, w walker.Walkable) {
	e.beforeNode = util.EnterNode{}
	if len(e.beforeNodeList) != 0 {
		e.beforeNodeList = e.beforeNodeList[:len(e.beforeNodeList)-1]
	}
	if len(e.beforeNodeList) != 0 {
		e.beforeNode = e.beforeNodeList[len(e.beforeNodeList)-1]
	}

}

func (e *EditWalker) EnterChildList(key string, w walker.Walkable) {
	e.beforeNode.Key = key
	e.beforeNode.Node = w.(node.Node)
	e.beforeNodeList = append(e.beforeNodeList, e.beforeNode)
}

func (e *EditWalker) LeaveChildList(key string, w walker.Walkable) {
	e.beforeNode = util.EnterNode{}
	if len(e.beforeNodeList) != 0 {
		e.beforeNodeList = e.beforeNodeList[:len(e.beforeNodeList)-1]
	}
	if len(e.beforeNodeList) != 0 {
		e.beforeNode = e.beforeNodeList[len(e.beforeNodeList)-1]
	}
}
