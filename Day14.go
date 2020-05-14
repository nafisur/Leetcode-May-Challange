package leetcode

/*

Implement Trie (Prefix Tree)
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/
type Trie struct {
	child [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	current := this
	for _, ch := range word {
		index := ch - 'a'
		if current.child[index] == nil {
			current.child[index] = &Trie{}
		}
		current = current.child[index]
	}
	current.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	current := this
	for _, ch := range word {
		index := ch - 'a'
		if current.child[index] == nil {
			return false
		}
		current = current.child[index]
	}
	return current.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, ch := range prefix {
		index := ch - 'a'
		if current.child[index] == nil {
			return false
		}
		current = current.child[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
