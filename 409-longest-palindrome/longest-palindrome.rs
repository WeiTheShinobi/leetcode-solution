impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        s.bytes().fold(vec![0; 58], |mut acc, x| {
            acc[(x - b'A') as usize] += 1;
            acc
        }).iter().fold(0, |acc, &x|
            acc + if acc & 1 == 0 {
                x
            } else {
                x - (x&1)
            }
        )
    }
}