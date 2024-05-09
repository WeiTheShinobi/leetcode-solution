use std::{cmp::*, vec};

impl Solution {
    pub fn maximum_happiness_sum(mut happiness: Vec<i32>, k: i32) -> i64 {
        let mut result: i64 = 0;

        happiness.sort_by(|a, b| b.cmp(a));
        for i in 0..k {
            result += max((happiness[i as usize] - i) as i64, 0);
        }

        result
    }
}