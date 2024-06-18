use std::cmp::Ordering;

impl Solution {
    pub fn max_profit_assignment(difficulty: Vec<i32>, profit: Vec<i32>, mut worker: Vec<i32>) -> i32 {
        let n = difficulty.len();
        let mut memo = vec![(0,0);n];
        for i in 0..n {
            memo[i] = (difficulty[i], profit[i]);
        }
        memo.sort();
        worker.sort();
        let mut ci = 0;
        let mut max = 0;
        let mut result = 0;
        for v in worker {
            while ci < n && memo[ci].0 <= v {
                max = std::cmp::max(max, memo[ci].1);
                ci += 1;
            }
            ci = ci.saturating_sub(1);
            result += max;
        }

        result
    }
}