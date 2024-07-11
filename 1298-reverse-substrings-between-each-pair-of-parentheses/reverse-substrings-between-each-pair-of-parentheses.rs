impl Solution {
    pub fn reverse_parentheses(s: String) -> String {
        let mut s = s.into_bytes();

        let mut stack = vec![];
        for end in 0..s.len() {
            match s[end] {
                b'(' => stack.push(end + 1),
                b')' => {
                    let start = stack.pop().unwrap();
                    let slice = &mut s[start..end];
                    slice.reverse();
                }
                _ => {}
            }
        }

        let mut bracket_index = 0;
        for index in 0..s.len() {
            match s[index] {
                b'(' | b')' => continue,
                _ => s[bracket_index] = s[index],
            }
            bracket_index += 1;
        }

        s.resize(bracket_index, 0);

        unsafe { String::from_utf8_unchecked(s) }
    }
}