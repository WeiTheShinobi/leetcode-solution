impl Solution {
    pub fn num_steps(s: String) -> i32 {
        s.bytes()
            .rev()
            .map(|b| (b - b'0') as i32)
            .scan((0, 0), |(s, c), b| {
                *s += match (b + *c) & 1 {
                    0 => 1,
                    _ => {
                        *c = 1;
                        2
                    }
                };
                Some((*s, *c))
            })
            .take(s.len() - 1)
            .last()
            .map(|(s, c)| s + c)
            .unwrap_or_default()
    }
}