use std::cmp::{max, min};

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut result = 0;
        let mut min_price = prices[0];
        for &price in prices.iter().skip(1) {
            min_price = min(price, min_price);
            result = max(result, price-min_price);
        }   
        
        result
    }
}