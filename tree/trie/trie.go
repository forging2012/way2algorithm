package trie

type TrieNode struct {
	path     string
	Children map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[byte]*TrieNode),
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (trie *Trie) Insert(path string) {
	root := trie.Root
	for i := 0; i < len(path); i++ {
		child, ok := root.Children[path[i]]
		if !ok {
			root.Children[path[i]] = NewTrieNode()
			child = root.Children[path[i]]
		}
		root = child
	}
	root.path = path
}

func (trie *Trie) Has(path string) bool {
	root := trie.Root
	for i := 0; i < len(path); i++ {
		child, ok := root.Children[path[i]]
		if !ok {
			return false
		}
		root = child
	}

	return root.path != ""
}

func (trie *Trie) Delete(path string) {
	trie.deleteFrom(trie.Root, path)
}

func (trie *Trie) deleteFrom(root *TrieNode, path string) bool {
	if path == "" {
		if root.path == "" {
			return false
		}
		root.path = ""
		return len(root.Children) == 0
	}

	child, ok := root.Children[path[0]]
	if !ok {
		return false
	}

	deleted := trie.deleteFrom(child, path[1:])
	if deleted {
		delete(root.Children, path[0])
	}

	return root.path == "" && len(root.Children) == 0
}
