package dag

type dag struct {
	nodes nodes
	links links
	nodeIndex map[int64]linksPair
	linkIindex map[int64]*nodePair
}

type links []*link
type linksPair [2]*links


type link struct {
	no	int64
	source *node
	target *node
	colmap map[string]*column
}
type nodes []*node
type nodePair [2]* node
type node struct {
	no int64
	name string
}

type inputNode struct{
	node
}

type outputNode struct{
	node
}

type operateNode struct{
	node
}

type column struct {
	colType string
	colName string
	calValue interface{}
}


