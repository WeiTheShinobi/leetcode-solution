impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        if c.count_ones() <= 1 {
            return true;
        }

        let limit = (c as f64).sqrt() as i32;

        for i in 0..limit {
            let a = i * i;
            let b = c - a;

            let mut lo = i;
            let mut hi = (b as f64).sqrt() as i32;


            while lo <= hi {
                let mid = lo + (hi - lo) / 2;
                let mids = mid * mid;

                if b == mids {
                    return true;
                }

                if b < mids {
                    hi = mid - 1;
                } else {
                    lo = mid + 1;
                }
            }
        }

        false
    }
}