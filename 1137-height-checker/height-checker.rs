impl Solution {
    pub fn height_checker(heights: Vec<i32>) -> i32 {
        let mut s = heights.clone();
        s.sort();
        heights.iter().enumerate().filter(|x| *x.1 != s[x.0]).count() as i32
    }
}