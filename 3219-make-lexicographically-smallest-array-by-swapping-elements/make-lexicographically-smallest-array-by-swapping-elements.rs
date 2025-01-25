impl Solution {
    pub fn lexicographically_smallest_array(nums: Vec<i32>, limit: i32) -> Vec<i32> {
        let mut result = vec![0; nums.len()];
        let mut sort = vec![];
        for i in 0..nums.len() {
            sort.push(i as i32);
        }
        sort.sort_by_key(|a| nums[*a as usize]);

        let mut i = 0;
        while i < sort.len() {
            let start = i;
            i += 1;
            while i < sort.len() && nums[sort[i] as usize] - nums[sort[i - 1] as usize] <= limit {
                i += 1;
            }
            let mut s: Vec<i32> = sort[start..i].to_vec();
            s.sort();

            for j in start..i {
                result[s[j-start] as usize] = nums[sort[j] as usize];
            }
        }
        result
    }
}
