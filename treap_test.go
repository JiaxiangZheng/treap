package treap

import (
    "testing"
    "fmt"
    "time"
    "math/rand"
)

func init() {
    rand.Seed(0)
}

func TestTreap(t *testing.T) {
    tree := NewTreap()
    const SIZE = 1 << 20
    tic := time.Now()
    tree.Insert(Item{0, 't'})
    for i := 1; i<SIZE; i++ {
        tree.Insert(Item{i, 'a'})
    }
    fmt.Println(time.Now().Sub(tic))
    fmt.Printf("height of tree is %d for %d elements\n", tree.Height(), tree.Len())
    tv := rand.Int()%SIZE
    if !tree.IsExist(tv) {
        t.Errorf("%d should exist in the treap\n", tv)
    }
    node := tree.Index(0)
    if node == nil || node.item.value != 't' {
        t.Errorf("the item with key=%d should exist and the value should be %c\n", 0, 't')
    }
}

func TestEmptyTreap(t *testing.T) {
    tree := NewTreap()
    if tree.Len() != 0 {
        t.Errorf("NewTreap() should return an empty treap\n")
    }
}

func TestRemove(t *testing.T) {
    tree := NewTreap()
    tree.Insert(Item{0, 'a'})
    tree.Insert(Item{1, 'b'})
    tree.Insert(Item{2, 'c'})
    tree.Insert(Item{3, 'd'})

    if tree.Remove(-3) {
        t.Errorf("-3 is not exist in the treap, expect return false, but return true\n")
    }
    if !tree.Remove(2) {
        t.Errorf("2 is exist in the treap, expect return true, but get false\n")
    }
    tree.Preorder(func(item Item) {
        fmt.Printf("%v ", item)
    })
    fmt.Printf("\n")
}
