// it isn't binary tree

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
    return search(cloned, target->val);  
}

TreeNode* search(TreeNode* tree, int val) {
    if (tree->val == val) return tree;

    if (tree->left != NULL) {
        auto found = search(tree->left, val);
        if (found != NULL) {
            return found;
        }
    }

    if (tree->right != NULL) {
        auto found = search(tree->right, val);
        if (found != NULL) {
            return found;
        }
    }

    return NULL;
}