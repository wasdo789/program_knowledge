package structure

type TrieNode struct {
	Pass  int
	End   int
	Child map[rune]*TrieNode
}

// func (t TrieNode) Print() {
// 	res := fmt.Sprint("CurNode Pass:%v,End:%v\n", t.Pass, t.End)
// 	for k,v:=range t.Child {
// 		fmt.Sprint("\tChildNode Pass:%v,End:%v\n", t.Pass, t.End)
// 	}

// }

func NewTrieNode() *TrieNode {
	n := &TrieNode{}
	n.Child = make(map[rune]*TrieNode)
	return n
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = NewTrieNode()
	return t
}

func (r *Trie) Insert(s string) {
	curNode := r.Root
	curNode.Pass++
	for _, x := range s {
		//fmt.Printf("%c\n", x)
		if v, ok := curNode.Child[x]; ok {
			v.Pass++
		} else {
			n := NewTrieNode()
			n.Pass = 1
			curNode.Child[x] = n
		}
		curNode = curNode.Child[x]
	}
	curNode.End++
	return
}

func (r *Trie) Delete(s string) {
	if r.Search(s) != 0 {
		curNode := r.Root
		curNode.Pass--
		for _, x := range s {
			//fmt.Printf("%c\n", x)
			curNode.Child[x].Pass--
			if curNode.Child[x].Pass == 0 {
				delete(curNode.Child, x)
				break
			}
			curNode = curNode.Child[x]
		}
		curNode.End--
	}

}

func (r Trie) Search(s string) int {
	curNode := r.Root
	for _, c := range s {
		if _, ok := curNode.Child[c]; ok {
			curNode = curNode.Child[c]
		} else {
			return 0
		}
	}
	return curNode.End
}

func (r Trie) PrefixSearch(s string) int {
	curNode := r.Root
	for _, c := range s {
		if _, ok := curNode.Child[c]; ok {
			curNode = curNode.Child[c]
		} else {
			return 0
		}
	}
	return curNode.Pass
}
