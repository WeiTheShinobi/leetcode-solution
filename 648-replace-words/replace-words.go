func replaceWords(dictionary []string, sentence string) string {
	type trie struct {
		isEnd bool
		t     []*trie
	}

	t := &trie{t: make([]*trie, 26)}

	for _, s := range dictionary {
		curr := t
		for i, c := range s {
			if curr.t[c-'a'] != nil {
				curr = curr.t[c-'a']
			} else {
				curr.t[c-'a'] = &trie{
					t: make([]*trie, 26),
				}
				curr = curr.t[c-'a']
			}
			if i == len(s)-1 {
				curr.isEnd = true
			}
		}
	}

	ss := strings.Split(sentence, " ")

	var builder strings.Builder
	for i, s := range ss {
    
		curr := t
		var currBuilder strings.Builder
		for _, c := range s {
			if curr.t[c-'a'] == nil {
				break
			}
			currBuilder.WriteByte(byte(c))
			curr = curr.t[c-'a']
			if curr.isEnd {
				break
			}
		}
		if curr.isEnd {
			builder.WriteString(currBuilder.String())
		} else {
      builder.WriteString(s)
    }
		if i != len(ss)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}
