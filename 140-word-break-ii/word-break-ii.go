func wordBreak(s string, wordDict []string) []string {
    words := make(map[string]bool)
    for _, word := range wordDict {
        words[word] = true
    }
    
    return dfs(s, words, make(map[int][]string), 0)
}

func dfs(s string, words map[string]bool, record map[int][]string, pos int) []string {
    if result, ok := record[pos]; ok {
        return result
    }
    var result []string
    
    for i := pos+1; i <= len(s); i++ {
        if words[s[pos:i]] {
            if i != len(s) {
                rest := dfs(s, words, record, i)
                for _, r := range rest {
                    result = append(result, s[pos:i] + " " + r)
                }                
            } else {
                result = append(result, s[pos:i])
                break
            }
        }
    }
    
    record[pos]  = result
    return result
}