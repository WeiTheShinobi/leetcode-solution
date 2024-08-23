use std::cmp::{max, min};

impl Solution {
    pub fn pair_sum(head: Option<Box<ListNode>>) -> i32 {
        let mut arr = vec![];
        let mut current = head;

        while let Some(node) = current {
            arr.push(node.val);
            current = node.next;
        }

        let mut result = 0;
        for i in 0..arr.len() / 2 {
            result = max(result, arr[i] + arr[arr.len()-i-1]);
        }
        result
    }
}