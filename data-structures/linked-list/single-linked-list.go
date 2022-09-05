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
    list.size++
  }


}

func (list *LinkedList) DeleteFirst() {

  if list.size == 1 {
    list.head = nil
  }

    list.head = list.head.next
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

func (list *LinkedList) GetAt(index int) *Node {

  if index > 0 && index > list.size {
    panic("Index out of bound")
  }
  
  current := list.head

  for i := 0; i < index; i++ {
    current = current.next
  }

  return current

}

func (list *LinkedList) InsertAt(index int, data int) {

  if index > list.size {
    panic("Index out of bound")
  }

  if index == 0 {
    list.AddFirst(data)
  }

  if index == list.size - 1 {
      list.AddLast(data)
    } else {
    prevNode := list.GetAt(index - 1)
    newNode := &Node { data: data, next: nil }
  
    newNode.next = prevNode.next.next
    prevNode.next = newNode
    list.size++

    }

}

func (list *LinkedList) DeleteAt(index int) {


  if index > list.size {
    panic("Index out of bound")
  }

  if index == 0 {
    list.DeleteFirst()
  }


  if index == list.size - 1 {

    current := list.head

    for current.next != nil {
      current = current.next
    }

    current = nil

  } else {
    prevNode := list.GetAt(index - 1)
    prevNode.next = prevNode.next.next
  }

  list.size--
    
}

func (list *LinkedList) Display() { 

  current := list.head

  for current != nil {
     fmt.Printf("%v->", current.data)
    current = current.next
  }

  fmt.Println()
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

  node := list.GetAt(2)

  fmt.Println("Node at index 2:", node)

  list.DeleteAt(2)

  list.Display()

  list.InsertAt(2, 5)

  list.Display()

}
