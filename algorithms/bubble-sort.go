package main

import "fmt"

func main() {

  arr := []int{5,3,2,1,3,5,6}
  temp := 0

  bubbleSort(arr, temp)

  fmt.Println("Sorted arr:", arr)

}


func bubbleSort(arr []int, temp int) {

   isSorted := false

    for i := 0; i < len(arr); i++ {

      isSorted = true

    for j :=1; j < len(arr) - i; j++ {

      if arr[j] < arr[j - 1] {
          temp = arr[j]
          arr[j] = arr[j-1]
          arr[j-1] = temp

          isSorted = false
      }

      if isSorted {
        return
      }
    }

  }


}
