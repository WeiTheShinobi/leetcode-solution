use std::cmp::{max, min};


impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut dp: Vec<Vec<i32>> = vec![vec![0; 2]; prices.len()+1];
        dp[0][1] = i32::MIN;
        for i in 0..prices.len() {
            dp[i+1][0] = max(dp[i][0], dp[i][1] + prices[i]);
            dp[i+1][1] = max(dp[i][1], dp[i][0] - prices[i]);
        }

        dp[prices.len()][0]
    }
}