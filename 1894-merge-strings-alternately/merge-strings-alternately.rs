use std::cmp::{max, min};

impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {
        let mut result = String::new();
        let word1 = word1.as_bytes();
        let word2 = word2.as_bytes();
        
        for i in 0..max(word1.len(), word2.len()) {
            if i < word1.len() {
                result.push(word1[i] as char);
            }
            if i < word2.len() {
                result.push(word2[i] as char);
            }
        }
        result
    }
}