package nestedsets

import (
	"math"
	"math/big"
	"unicode"

	"github.com/golang-collections/collections/stack"
)

const baseLength = 62 // 0-9 A-Z a-z

type Node struct {
	ID            string
	ParentID      string
	Childs        []*Node
	LeftBowerInt  int
	RightBowerInt int
}

type Set struct {
	ID            string
	LeftBower     string
	RightBower    string
	LeftBowerInt  int
	RightBowerInt int
}

func NodeFromSets(sets []Set) *Node {
	if len(sets) < 1 {
		return nil
	}
	node := &Node{
		ID:            sets[0].ID,
		Childs:        []*Node{},
		LeftBowerInt:  sets[0].LeftBowerInt,
		RightBowerInt: sets[0].RightBowerInt,
	}

	if len(sets) == 1 {
		return nil
	}
	st := stack.New()
	st.Push(node)
	for _, item := range sets[1:] {

		n := &Node{
			ID:            item.ID,
			Childs:        []*Node{},
			LeftBowerInt:  item.LeftBowerInt,
			RightBowerInt: item.RightBowerInt,
		}
		last := &Node{}
		for st.Len() > 0 {
			last = st.Peek().(*Node)
			if item.LeftBowerInt > last.LeftBowerInt && item.LeftBowerInt < last.RightBowerInt {
				break
			}
			last = st.Pop().(*Node)
		}
		if st.Len() == 0 {
			return node
		}
		// this node is a child of previous node
		if item.RightBowerInt-item.LeftBowerInt != 1 { // this node is not a leaf
			st.Push(n)
		}
		n.ParentID = last.ID
		last.Childs = append(last.Childs, n)

	}

	return node
}

// maintaining
func SetsFromNote(root Node) []Set {
	bowerIDs := make([]string, 0)
	lastChildMap := map[string]struct{}{}
	visit(root, bowerIDs, lastChildMap)
	bowerLen := len(bowerIDs)
	if bowerLen == 0 {
		return nil
	}
	stepLen := int(math.Ceil(float64(bowerLen/baseLength))*baseLength) / bowerLen
	sets := make([]Set, 0)
	setMap := make(map[string]Set)

	for idx, item := range bowerIDs {
		if v, ok := setMap[item]; ok {
			v.RightBower = getBowerStr(idx * stepLen)
			setMap[item] = v
		} else {
			setMap[item] = Set{
				ID:        item,
				LeftBower: getBowerStr(idx * stepLen),
			}
		}
	}

	for _, val := range setMap {
		sets = append(sets, val)
	}

	return sets
}

func visit(n Node, bowerIDs []string, lastChildMap map[string]struct{}) {
	bowerIDs = append(bowerIDs, n.ID)
	if n.Childs == nil || len(n.Childs) == 0 { // nếu k có childs, duyệt thêm 1 lần
		bowerIDs = append(bowerIDs, n.ID)
	}
	if _, ok := lastChildMap[n.ID]; ok && len(n.ParentID) > 0 {
		visit(Node{ParentID: n.ParentID}, bowerIDs, lastChildMap)
	}
	for idx, item := range n.Childs {
		if idx == len(n.Childs)-1 {
			lastChildMap[n.ID] = struct{}{}
		}
		visit(*item, bowerIDs, lastChildMap)
	}
}

func getBowerStr(bower int) string {
	s := []rune(big.NewInt(int64(bower)).Text(baseLength))
	for idx, r := range s {
		if unicode.IsLower(r) {
			s[idx] = unicode.ToUpper(r)
			continue
		}
		if unicode.IsUpper(r) {
			s[idx] = unicode.ToLower(r)
			continue
		}
	}
	return string(s)
}
