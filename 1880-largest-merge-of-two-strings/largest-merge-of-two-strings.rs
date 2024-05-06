impl Solution {
    pub fn largest_merge(word1: String, word2: String) -> String {
        let mut result: Vec<char> = vec!();

        let (mut i, mut j) = (0, 0);

        while i < word1.len() || j < word2.len() {
            if i < word1.len() && word1[i..] > word2[j..] {
                result.push(word1.chars().nth(i).unwrap());
                i += 1;
            } else {
                result.push(word2.chars().nth(j).unwrap());
                j += 1;
            }
        }

        result.into_iter().collect()
    }
}