package trie

type Node struct {
	isWord   bool
	children map[rune]*Node
}

type Trie struct {
	root *Node
}

func (t *Trie) Init() {
	t.root = new(Node)
	t.root.Init()
}

func (n *Node) Init() {
	n.isWord = false
	n.children = make(map[rune]*Node)
}

func (t *Trie) Insert(word string) bool {
	root := t.root
	for _, c := range word {
		_, ok := root.children[c]
		if !ok {
			n := new(Node)
			n.Init()
			root.children[c] = n
		}
		root = root.children[c]
	}
	root.isWord = true
	return true
}

func (t *Trie) Search(word string) bool {
	root := t.root
	for _, c := range word {
		_, ok := root.children[c]
		if !ok {
			return false
		}
		root = root.children[c]
	}
	return root.isWord
}

func (t *Trie) SearchPrefix(prefix string) bool {
	root := t.root
	for _, c := range prefix {
		_, ok := root.children[c]
		if !ok {
			return false
		}
		root = root.children[c]
	}
	return true
}
