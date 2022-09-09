package main


import "fmt"

func main() {


  arr := []int {5,4,3,2,1}

  current := 0


  for i := 0; i < len(arr) - 1; i++ {

    current = i

    for j := i+1; j < len(arr); j++ {
        
      if arr[j] < arr[current] {
        current = j
      }

    }

    if i != current {
      temp := arr[i]
      arr[i] = arr[current]
      arr[current] = temp
    }


  }


  fmt.Println(arr)


}

