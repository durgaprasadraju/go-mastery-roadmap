// Package trie implements a prefix tree (Trie) from scratch.
package trie

// Node represents a trie node.
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

// Trie is a prefix tree for string keys.
type Trie struct {
	root *Node
	size int
}

// New creates an empty trie.
func New() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

// Insert adds a word. O(m) where m is word length.
func (t *Trie) Insert(word string) {
	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = &Node{children: make(map[rune]*Node)}
		}
		cur = cur.children[ch]
	}
	if !cur.isEnd {
		t.size++
	}
	cur.isEnd = true
}

// Search reports if the exact word exists. O(m).
func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, ch := range word {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return cur.isEnd
}

// StartsWith reports if any word has the given prefix. O(m).
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, ch := range prefix {
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return true
}

// Size returns the number of words.
func (t *Trie) Size() int { return t.size }

// Autocomplete returns words with the given prefix. O(m + k).
func (t *Trie) Autocomplete(prefix string) []string {
	cur := t.root
	for _, ch := range prefix {
		if cur.children[ch] == nil {
			return nil
		}
		cur = cur.children[ch]
	}
	var results []string
	var dfs func(*Node, string)
	dfs = func(n *Node, path string) {
		if n.isEnd {
			results = append(results, path)
		}
		for ch, child := range n.children {
			dfs(child, path+string(ch))
		}
	}
	dfs(cur, prefix)
	return results
}
