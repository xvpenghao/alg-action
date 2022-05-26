#!/usr/bin/env python
# -*- coding: utf-8 -*-


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None:
            return None
        while root:
            if root.val < p.val and root.val < q.val:
                root = root.right
            elif root.val > p.val and root.val > q.val:
                root = root.left
            else:
                break
        return root

    def lowestCommonAncestor2(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None:
            return None
        if root.val < p.val and root.val < q.val:
            return self.lowestCommonAncestor2(root.right, p, q)
        elif root.val > p.val and root.val > q.val:
            return self.lowestCommonAncestor2(root.left, p, q)
        return root


if __name__ == '__main__':
    root = TreeNode(6)

    root.left = TreeNode(2)
    root.left.left = TreeNode(0)
    root.left.right = TreeNode(4)
    root.left.right.left = TreeNode(3)
    root.left.right.right = TreeNode(5)

    root.right = TreeNode(8)
    root.right.left = TreeNode(7)
    root.right.left = TreeNode(9)

    s = Solution()
    res = s.lowestCommonAncestor(root, TreeNode(3), TreeNode(5))
    if res:
        print(res.val)
    else:
        print("错误")
