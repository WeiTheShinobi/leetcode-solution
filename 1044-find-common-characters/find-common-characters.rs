impl Solution {
    pub fn common_chars(words: Vec<String>) -> Vec<String> {
        words.iter().
            map(|x| x.bytes().
                fold(vec![0; 26], |mut acc, x| {
                    acc[(x - b'a') as usize] += 1;
                    acc
                })).
            fold(vec![i32::MAX; 26], |acc, x| {
                acc.iter().zip(x.iter()).map(|val| {
                    std::cmp::min(*val.0, *val.1)
                }).collect()
            }).
            iter().enumerate().
            fold(vec![], |mut acc, en| {
                let s = std::char::from_u32((b'a' + en.0 as u8) as u32).unwrap().to_string();
                (0..*en.1).into_iter().for_each(|_| {
                    acc.push(s.clone());
                });
                acc
            })
    }
}