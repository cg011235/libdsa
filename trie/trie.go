package trie

import "sort"

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

func (t *Trie) SearchPrefixNode(prefix string) *Node {
	node := t.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return nil
		}
		node = node.children[char]
	}
	return node
}

// Suggest collects up to 3 suggestions lexicographically for a given prefix node.
func (t *Trie) Suggest(prefix string) []string {
	node := t.SearchPrefixNode(prefix)
	var suggestions []string
	var collect func(n *Node, prefix string)
	collect = func(n *Node, prefix string) {
		if n == nil {
			return
		}
		if n.isWord {
			suggestions = append(suggestions, prefix)
			if len(suggestions) == 3 {
				return
			}
		}
		for char, childNode := range n.children {
			collect(childNode, prefix+string(char))
			if len(suggestions) == 3 {
				break
			}
		}
	}

	collect(node, prefix)
	return suggestions
}

func SuggestedProducts(products []string, searchWord string) [][]string {
	trie := new(Trie)
	trie.Init()
	for _, product := range products {
		trie.Insert(product)
	}

	var result [][]string
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		suggestions := trie.Suggest(prefix)
		sort.Strings(suggestions) // Ensure lexicographical order
		result = append(result, suggestions)
	}

	return result
}
