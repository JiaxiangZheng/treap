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
    const SIZE = 1 << 5
    tree.Insert(Item{0, 't'})
    for i := 1; i<SIZE; i++ {
        tree.Insert(Item{i, 'a'})
    }
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
    if tree.Len() != 3 {
        t.Errorf("after remove an element from treap, the size should be 3\n")
    }
    tree.Preorder(func(item Item) {
        fmt.Printf("%v ", item)
    })
    fmt.Printf("\n")
}

func TestBenchmark(t *testing.T) {
    tree := NewTreap()
    const SIZE = 1 << 22
    tic := time.Now()
    for i := 0; i<SIZE; i++ {
        tree.Insert(Item{i, 'a'})
    }
    fmt.Printf("insert %d elements (height %d) use %v\n", SIZE, tree.Height(), time.Now().Sub(tic))

    const REMOVE_SIZE = 1 << 22
    tic = time.Now()
    for i := 0; i<REMOVE_SIZE; i++ {
        if !tree.Remove(i) {
            fmt.Printf("remove %d failed\n", i)
        }
    }
    fmt.Printf("remove %d elements use %v\n", REMOVE_SIZE, time.Now().Sub(tic))
    if tree.Len() != SIZE-REMOVE_SIZE {
        t.Errorf("failed in remove\n")
    }

    tic = time.Now()
    m := make(map[int]string)
    for i := 0; i<SIZE; i++ {
        m[i] = "a"
    }
    fmt.Printf("[map] insert %d elements use %v\n", REMOVE_SIZE, time.Now().Sub(tic))

    tic = time.Now()
    s := make([]int, SIZE)
    for i := 0; i<SIZE; i++ {
        s[i] = i
    }
    fmt.Printf("[slice] insert %d elements use %v\n", REMOVE_SIZE, time.Now().Sub(tic))
}
