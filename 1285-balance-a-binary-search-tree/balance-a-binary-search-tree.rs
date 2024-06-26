
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn balance_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut arr = Vec::new();
        Self::inorder(&root, &mut arr);
        Self::split(&arr, 0, arr.len())
    }

    fn inorder(node: &Option<Rc<RefCell<TreeNode>>>, arr: &mut Vec<i32>) {
        if let Some(node) = node {
            let node = node.borrow();
            Self::inorder(&node.left, arr);
            arr.push(node.val);
            Self::inorder(&node.right, arr);
        }
    }

    fn split(arr: &Vec<i32>, i: usize, j: usize) -> Option<Rc<RefCell<TreeNode>>> {
        if i == j {
            return None;
        }
        let m = (i + j) / 2;
        Some(Rc::new(RefCell::new(
            TreeNode {
                val: arr[m],
                left: Self::split(arr, i, m),
                right: Self::split(arr, m + 1, j),
            }
        )))
    }
}