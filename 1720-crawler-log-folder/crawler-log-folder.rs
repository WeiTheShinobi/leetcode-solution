impl Solution {
    pub fn min_operations(logs: Vec<String>) -> i32 {
        logs.iter().fold(0, |acc, x| {
            match x.as_str() {
                "./" => acc,
                "../" => acc - (acc > 0) as i32,
                _ => acc + 1,
            }
        })
    }
}