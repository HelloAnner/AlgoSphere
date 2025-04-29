package tree

// https://leetcode.cn/problems/implement-trie-prefix-tree/?envType=study-plan-v2&envId=top-interview-150
// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

// 前缀树: 每个节点表示一个字符，从根节点到叶子节点的路径表示一个单词

// 时间复杂度:
// insert: O(n)
// search: O(n)
// startsWith: O(n)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}


func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			// key 是字符，value 是子节点
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, c := range word {
		// 如果当前字符不存在，则创建一个新节点
		if cur.children[c] == nil {
			cur.children[c] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, c := range word {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return cur.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, c := range prefix {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return true
}