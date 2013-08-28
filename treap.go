// A simple implementation of treap. Key in treap is assumed to be unique.
package treap

import (
    "time"
    "fmt"
    "math/rand"
)

type Treap struct {
    size int
    root *Node
}
type Item struct {
    key int
    value interface{}
}

type Node struct {
    item Item
    priority int
    left, right *Node
}

func NewNode(item Item, priority int) *Node {
    return &Node{item, priority, nil, nil}
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Returns a treap with no item
func NewTreap() *Treap {
    return &Treap{0, nil}
}
func (t *Treap) Len() int {
    return t.size
}

// Returns the height of the treap
func (t *Treap) Height() int {
    return height(t.root)
}

func height(root *Node) int {
    if root == nil {
        return 0
    }
    leftHeight, rightHeight := height(root.left), height(root.right)
    if leftHeight > rightHeight {
        return leftHeight + 1
    } else {
        return rightHeight + 1
    }
}

// Returns true if an item with key is exist in the treap 
func (t *Treap) IsExist(key int) bool {
    return t.Index(key) != nil
}

// Returns the item with key. If the item is not exist, return nil 
func (t *Treap) Index(key int) *Node {
    return index(t.root, key)
}

func index(root *Node, key int) *Node {
    if root == nil {
        return nil
    }
    if root.item.key < key {
        return index(root.right, key)
    } else if root.item.key > key {
        return index(root.left, key)
    }
    return root
}

// Insert insert a new item to the treap. Currently, it support no duplicate
// keys, returns true if the item is inserted to the treap, false if there is
// already an item with the same key
func (t *Treap) Insert (item Item) bool {
    res, flag := insert(t.root, NewNode(item, rand.Int()))
    if flag {
        t.root = res
        t.size++
        return true
    }
    return false
}

func rightRotate(node *Node) *Node {
    if node == nil || node.left == nil {
        panic("error occured : node and node's left child must not be nil")
    }
    res := node.left
    node.left = res.right
    res.right = node
    return res
}
func leftRotate(node *Node) *Node {
    if node == nil || node.right == nil {
        panic("error occured : node and node's right child must not be nil")
    }
    res := node.right
    node.right = res.left
    res.left = node
    return res
}
func insert(root *Node, node *Node) (*Node, bool) {
    if root == nil {
        return node, true
    }
    var flag bool = true
    var res *Node = nil
    if node.item.key < root.item.key {
        res, flag = insert(root.left, node)
        if !flag { return root, false }
        root.left = res
		if root.left.priority < root.priority {
            root = rightRotate(root)
		}
	} else if root.item.key < node.item.key {
        res, flag = insert(root.right, node)
        if !flag { return root, false }
        root.right = res
		if root.right.priority < root.priority {
			root = leftRotate(root)
		}
	}
	return root, flag
}

// Delete delete the item with the parameter key
func (t *Treap) Remove (key int) bool {
    res, flag := remove(t.root, key)
    if flag {
        t.root = res
        t.size--
        return true
    }
    return false
}

func remove(root *Node, key int) (*Node, bool) {
    var flag bool = true
    var res *Node = nil
    if root == nil {
        return root, false
    } else if root.item.key > key { //left child
        res, flag = remove(root.left, key)
        if flag == false {
            return root, false
        }
        root.left = res
    } else if root.item.key < key {
        res, flag = remove(root.right, key)
        if flag == false {
            return root, false
        }
        root.right = res
    } else {    // found the item, remove it
        if root.left == nil && root.right == nil {  // leaf node
            return nil, true
        } else if root.left == nil {
            return root.right, true
        } else if root.right == nil {
            return root.left, true
        } else {
            pre, cur := root, root.left
            for cur.right != nil {
                pre = cur
                cur = cur.right
            }
            pre.right = nil
            cur.left = root.left
            cur.right = root.right
            cur.priority = root.priority
            return cur, true
        }
    }
    return root, flag
}

func (t *Treap) Preorder(fun func(Item)) {
    preorder(t.root, fun)
}

func preorder(root *Node, fun func(Item)) {
    if root == nil {
        return
    }
    fun(root.item)
    preorder(root.left, fun)
    preorder(root.right, fun)
}

func (t *Treap) Inorder(fun func(Item)) {
    inorder(t.root, fun)
}

func inorder(root *Node, fun func(Item)) {
    if root == nil {
        return
    }
    inorder(root.left, fun)
    fun(root.item)
    inorder(root.right, fun)
}
