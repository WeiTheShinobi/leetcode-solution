impl Solution {
    pub fn subset_xor_sum(nums: Vec<i32>) -> i32 {
        nums.iter().fold(0, |acc, v| acc | v) << (nums.len() - 1)
    }
}