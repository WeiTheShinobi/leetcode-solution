use std::collections::{HashMap};

impl Solution {
    pub fn frequency_sort(s: String) -> String {
        let mut map = HashMap::new();

        s.chars().for_each(|x| *map.entry(x).or_insert(0) += 1);

        let mut v = map.into_iter().collect::<Vec<(char, i32)>>();
        v.sort_unstable_by_key(|pair| -pair.1);
        v.into_iter()
            .flat_map(|(c, n)| vec![c; n as usize])
            .collect()
    }
}