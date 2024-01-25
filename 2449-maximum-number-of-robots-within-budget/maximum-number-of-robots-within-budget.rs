use std::collections::vec_deque;

impl Solution {
    pub fn maximum_robots(charge_times: Vec<i32>, running_costs: Vec<i32>, budget: i64) -> i32 {
        let mut q = vec_deque::VecDeque::new();
        let mut result = 0;
        let mut s: i64 = 0;
        let mut left = 0;

        for (right, &ct) in charge_times.iter().enumerate() {
            while !q.is_empty() && ct > charge_times[*q.back().unwrap()] {
                q.pop_back();
            }

            q.push_back(right);
            s += running_costs[right] as i64;

            while !q.is_empty() && charge_times[q[0]] as i64 + (right - left + 1) as i64 * s as i64 > budget {
                s -= running_costs[left] as i64;

                while !q.is_empty() && q[0] <= left {
                    q.pop_front();
                }

                left += 1;
            }

            result = result.max(right - left + 1);
        }

        result as i32
    }
}