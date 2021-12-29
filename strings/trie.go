package strings

const MaxPrefix = 1000

type Trie struct {
	nodes  [][]int
	nNodes int
	list   []string
	nList  int
}

func (trie *Trie) InitTrie() {
	for i := 0; i < MaxPrefix; i++ {
		trie.nodes[i] = make([]int, MaxPrefix)
		for j := 0; j < MaxPrefix; j++ {
			trie.nodes[i][j] = -1
		}
	}
}

func (trie *Trie) Insert(ss string) {
	node := 0
	for i := 0; i < len(ss); i++ {
		let := int(ss[i] - 'a')
		if trie.nodes[node][let] == -1 {
			trie.nodes[node][let] = trie.nNodes
			node = trie.nNodes
			trie.nNodes++
		} else {
			node = trie.nodes[node][let]
		}
	}
}

func (trie *Trie) ListPrefix() {
	trie.list = make([]string, trie.nNodes)
	trie.nList = 0

}

func (trie *Trie) listPrefixRec(pre string, node int) {
	if node >= trie.nNodes {
		return
	}
	trie.list[trie.nList] = pre
	trie.nList++
	for i := 0; i < 26; i++ {
		if trie.nodes[node][i] != -1 {
			trie.listPrefixRec(pre+string(rune(i+'a')), trie.nodes[node][i])
		}
	}
}
