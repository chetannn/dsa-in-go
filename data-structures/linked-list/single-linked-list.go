package main

import "fmt"

type Node struct {
  data int
  next *Node
}

type LinkedList struct {
  head *Node
  size int
}

func (list *LinkedList) Size() int {
  return list.size
}

func (list *LinkedList) AddFirst(data int) {

    node := &Node { data: data, next: nil  }

  if list.head == nil {
    list.head = node
  } else {
    node.next = list.head
    list.head = node
  }

    list.size++
}

func (list *LinkedList) AddLast(data int) {

  if list.head == nil {
    list.AddFirst(data)
  } else {

    current := list.head
    
    for current.next != nil {
        current = current.next
    }

    node := &Node { data: data, next: nil  }
    current.next = node
  }

  list.size++

}

func (list *LinkedList) Reverse() {

  current := list.head
  var prev *Node;
  var temp *Node;

  for current != nil {
    
    temp = current.next
    current.next = prev
    prev = current

    current = temp
  }

  list.head = prev

}
func (list *LinkedList) Display() { current := list.head

  for current != nil {
     fmt.Printf("%v->", current.data)
    current = current.next
  }
}

func main() {
 
  list := LinkedList {}

  list.AddFirst(2)
  list.AddFirst(3)
  list.AddLast(4)

  list.Display()

  fmt.Println("\n")

  list.Reverse()

  list.Display()

}
