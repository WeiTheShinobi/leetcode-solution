// Definition for a binary tree node.
  // #[derive(Debug, PartialEq, Eq)]
  // pub struct TreeNode {
  //   pub val: i32,
  //   pub left: Option<Rc<RefCell<TreeNode>>>,
  //   pub right: Option<Rc<RefCell<TreeNode>>>,
  // }
  // 
  // impl TreeNode {
  //   #[inline]
  //   pub fn new(val: i32) -> Self {
  //     TreeNode {
  //       val,
  //       left: None,
  //       right: None
  //     }
  //   }
  // }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = vec![];

        fn r(node: &Option<Rc<RefCell<TreeNode>>>, arr: &mut Vec<i32>) {
            if let Some(n) = node {
                r(&n.borrow().left, arr);
                r(&n.borrow().right, arr);
                arr.push(n.borrow().val);
            }
        }

        r(&root, &mut result);
        result
    }
}