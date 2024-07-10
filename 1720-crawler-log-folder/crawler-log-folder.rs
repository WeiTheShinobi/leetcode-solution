impl Solution {
    pub fn min_operations(logs: Vec<String>) -> i32 {
        logs.iter().fold(0, |acc, x| {
            match x.as_str() {
                "./" => acc,
                "../" => if acc > 0 { acc - 1 } else { acc }
                _ => acc + 1,
            }
        })
    }
}