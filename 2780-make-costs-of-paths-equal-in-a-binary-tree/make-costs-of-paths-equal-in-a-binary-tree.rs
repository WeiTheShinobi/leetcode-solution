use std::cmp::{max, min};

impl Solution {
    pub fn min_increments(n: i32, mut cost: Vec<i32>) -> i32 {
        let mut result = 0;
        for i in (1..n/2+1).rev() {
            let i = i as usize;
            let (mut left, mut right) = (cost[i * 2 - 1], cost[i * 2]);
            (left, right) = (min(left, right), max(left, right));
            result += right - left;
            cost[i - 1] += right
        }

        result
    }
}