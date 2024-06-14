impl Solution {
    pub fn min_increment_for_unique(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();

        (1..nums.len()).fold(0, |acc, idx| {
            let val = nums[idx] - nums[idx - 1];
            if val <= 0 {
                nums[idx] = nums[idx-1] + 1;
                acc + 1 - val
            } else {
                acc
            }
        })
    }
}