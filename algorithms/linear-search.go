package main

import "fmt"

func main() {

  fmt.Println("Linear search")

  arr := []int {7,4,9,23,10}

  found := linearSearch(arr, 23)


  if found {
    fmt.Println("Element found")
  } else {
    fmt.Println("Element not found")
  }


}

func linearSearch(arr []int, target int) bool {
  
  found := false

  for i := 0; i < len(arr); i++ {

    if(arr[i] == target) {
        found = true
        return found
    }
  }

  return found
}
