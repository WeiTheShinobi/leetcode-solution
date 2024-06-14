use std::collections::HashSet;

impl Solution {
    pub fn min_increment_for_unique(mut nums: Vec<i32>) -> i32 {
        nums.sort();
        let mut m = HashSet::new();
        nums.iter().for_each(|&x| {
            m.insert(x);
        });

        let mut result = 0;
        let mut i = -1;
        for j in 1..nums.len() {
            if nums[j] == nums[j-1] {
                if i < nums[j] {
                    i = nums[j] + 1;
                }
                while m.contains(&i) {
                    i += 1;
                }
                result += i - nums[j];
                m.insert(i);
            }
        }

        result
    }
}