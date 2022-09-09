package main

import "fmt"

func main() {
  
  arr := []int{5,4,3,2,1}
  temp := 0

  for i := 1; i < len(arr); i++ {

    j := i-1
    temp = arr[i]

    for j >= 0 && arr[j] > temp {
       arr[j+1] = arr[j]
      j--
    }

    arr[j+1] = temp

  }


  fmt.Println(arr)

}
