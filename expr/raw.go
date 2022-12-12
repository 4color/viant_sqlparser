package expr

import "github.com/viant/sqlparser/node"

//Raw represetns a raw expr
type Raw struct {
	Raw string
	X   node.Node
}

func NewRaw(raw string) *Raw {
	return &Raw{Raw: raw}
}
