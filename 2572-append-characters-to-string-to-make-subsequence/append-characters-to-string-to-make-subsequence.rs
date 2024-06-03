impl Solution {
    pub fn append_characters(s: String, t: String) -> i32 {
        (t.len() - s.
            bytes().
            fold(0 , |acc, x| {
                if acc < t.len() && x == t.as_bytes()[acc] {
                    acc + 1
                } else {
                    acc
                }
            })) as i32
    }
}