impl Solution {
    pub fn special_array(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        let (mut l, mut r): (i32, i32) = (-1, nums.len() as i32);
        while l + 1 != r {
            let mid = (r + l) / 2;
            let count = nums.
                iter().
                filter(|&x| x >= &mid).
                count() as i32;
            if count > mid {
                l = mid
            } else {
                r = mid
            }
        }

        if r == nums.
            iter().
            filter(|&x| x >= &r).
            count() as i32 {
            r
        } else { -1 }
    }
}