impl Solution {
    pub fn number_of_subarrays(nums: Vec<i32>, k: i32) -> i32 {
        let mut indices = vec![];

        for i in 0..nums.len() {
            if nums[i] % 2 == 1 { indices.push(i); }
        }

        let mut left = 0;
        let mut right = (k - 1) as usize;
        let mut result = 0;

        while right < indices.len() {
            let cur_start = indices[left] as i32;
            let mut prev_index = -1;
            if left != 0 { prev_index = indices[left-1] as i32; }

            let cur_end = indices[right] as i32;
            let mut next_index = nums.len() as i32;
            if right != indices.len()-1 { next_index = indices[right+1] as i32; }
            
            result += (cur_start - prev_index) * (next_index - cur_end);

            left += 1;
            right += 1;
        }
        result
    }
}