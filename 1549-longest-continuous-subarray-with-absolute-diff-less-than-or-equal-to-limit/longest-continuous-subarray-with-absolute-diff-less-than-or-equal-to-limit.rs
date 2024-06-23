use std::cmp::{max, min};
use std::collections::BinaryHeap;

#[derive(Eq, PartialEq, Ord, PartialOrd, Debug)]
struct Item(i32, usize);

impl Solution {
    pub fn longest_subarray(nums: Vec<i32>, limit: i32) -> i32 {
        let mut maxh = BinaryHeap::new();
        let mut minh = BinaryHeap::new();
        let mut result = 0;
        let mut l = 0;
        for (i, &val) in nums.iter().enumerate() {
            maxh.push(Item(val, i));
            minh.push(Item(-val, i));

            while maxh.peek().unwrap().0 + minh.peek().unwrap().0 > limit {
                l = min(maxh.peek().unwrap().1, minh.peek().unwrap().1) + 1;
                while maxh.peek().unwrap().1 < l { maxh.pop(); }
                while minh.peek().unwrap().1 < l { minh.pop(); }
            }

            result = max(result, i - l + 1);
        }

        result as i32
    }
}
