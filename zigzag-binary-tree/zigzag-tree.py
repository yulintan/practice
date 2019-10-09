class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def zigzagLevelOrder(self, root):
        result = []
        if not root:
            return result

        self.helper(root, result, 0, True)

        return result

    def helper(self, node, result, level, inorder):
        if len(result) == level:
            result.append([])

        if inorder:
            result[level].append(node.val)
        else:
            result[level].insert(0, node.val)

        if node.left is not None:
            self.helper(node.left, result, level+1, not inorder)

        if node.right is not None:
            self.helper(node.right, result, level+1, not inorder)
