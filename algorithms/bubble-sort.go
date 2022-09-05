package main

import "fmt"

func main() {

  arr := []int{5,3,2,1,3,5,6}
  temp := 0

  for i := 0; i < len(arr); i++ {
    for j :=i+1; j < len(arr); j++ {

      if arr[j] < arr[i] {
          temp = arr[i]
          arr[i] = arr[j]
          arr[j] = temp
      }
    }  

  }

  fmt.Println("Sorted arr:", arr)

}
