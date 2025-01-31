package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
    value int
    next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
    head *Node
}

// Add a new nodxs525-mukcha@xs-host603-020:~/Desktop/L1M4$ go mod init local/task1e to the end of the linked list
func (ll *LinkedList) Add(value int) {
    newNode := &Node{value: value}
    if ll.head == nil {
        ll.head = newNode
        return
    }package main 

    import "fmt"
    
    func sol(ljug int,sjug int) int {
        var lj int  =0 
        var sj int =0
        lj=ljug
        lj=lj-sjug
        sj=sjug
        sj=0
        sj=lj
        lj=ljug
        lj=lj-(sjug-sj)
        return lj
         
    }
    
    func main(){
        var ans=sol(7,3)
        fmt.Println(ans)
    }
    
    
    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

// Display the linked list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Print(current.value, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

// Delete a node by value
func (ll *LinkedList) Delete(value int) {
    if ll.head == nil {
        return
    }
    if ll.head.value == value {
        ll.head = ll.head.next
        return
    }
    current := ll.head
    for current.next != nil && current.next.value != value {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
    }
}

func main() {
    ll := &LinkedList{}

    ll.Add(1)
    ll.Add(2)
    ll.Add(3)

    fmt.Println("Linked List:")
    ll.Display()

    ll.Delete(2)
    fmt.Println("After deleting 2:")
    ll.Display()
}