// a treap implementation 

package treap

/*
func insert(root *treapNode, key int) (*treapNode, bool) {
    var res bool = true
    if root == nil {
        return &treapNode{key, rand.Int(), nil, nil}, res
    } else if (root.key > key) {
        root.left, res = insert(root.left, key)
        if root.left.priority < root.priority {
            root = rightRot(root)
        }
    } else if (root.key < key) {
        root.right, res = insert(root.right, key)
        if root.right.priority < root.priority {
            root = leftRot(root)
        }
    } else {
        return root, false
    }
    return root, res
}

// remove the node with key and return the root, if the second return is false,
// denotes a false remove operation which means key was not found in the treap
func remove(root *treapNode, key int) (*treapNode, bool) {
    var res bool
    if root == nil {
        return root, false
    } else if (root.key < key) {    // right
        root.right, res = remove(root.right, key)
    } else if (root.key > key) {
        root.left, res = remove(root.left, key)
    } else {
        if root.left == nil && root.right == nil {  // found the leaf node, just remove it
            root = nil
        } else if root.left == nil {
            root = root.right
        } else if root.right == nil {
            root = root.left
        } else if root.left.priority < root.right.priority {
            rightRotroot
            remove(root.right, key)
        } else {
            leftRotroot
            remove(root.left, key)
        }
        res = true
    }
    return root, res
}
*/
