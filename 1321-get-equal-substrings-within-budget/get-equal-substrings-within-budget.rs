use std::cmp::max;

impl Solution {
   pub fn equal_substring(s: String, t: String, max_cost: i32) -> i32 {
        let mut result = 0;
        let mut curr = 0;
        let mut l: usize = 0;
        let mut max_cost = max_cost;
        for r in 0..s.len() {
            if s.as_bytes()[r] != t.as_bytes()[r] {
                let sb = s.as_bytes()[r] as i32;
                let tb = t.as_bytes()[r] as i32;
                let mut diff = if sb - tb < 0 { tb - sb } else { sb - tb };

                while diff > max_cost && l < r {
                    let mut diff2 = s.as_bytes()[l] as i32 - t.as_bytes()[l] as i32;
                    if diff2 < 0 {
                        diff2 = -diff2;
                    }
                    max_cost += diff2;
                    l += 1;
                    curr -= 1;
                }
                if diff > max_cost {
                    l += 1;
                    continue;
                }
                max_cost -= diff;
            }
            curr += 1;
            result = max(result, curr);
        }

        result
    }
}