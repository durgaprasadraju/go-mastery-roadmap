package solutions

type Trie struct{ children map[rune]*Trie; end bool }

func NewTrie() *Trie { return &Trie{children: make(map[rune]*Trie)} }

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = &Trie{children: make(map[rune]*Trie)}
		}
		cur = cur.children[ch]
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, ch := range word {
		if cur.children[ch] == nil { return false }
		cur = cur.children[ch]
	}
	return cur.end
}