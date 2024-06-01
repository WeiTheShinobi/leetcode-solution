impl Solution {
    pub fn score_of_string(s: String) -> i32 {
        s.chars().
            zip(s.chars().skip(1)).
            map(|(a, b)| (a as u8).abs_diff(b as u8) as i32).
            sum()
    }
}