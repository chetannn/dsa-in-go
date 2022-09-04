package main

import "fmt"

func main() {

 fmt.Println("Merge Sort")

 first := []int{1,2,3}
 second := []int{4,5,6,7}

  merged := mergeArray(first, second)

  fmt.Println("After merge:", merged)

}


func mergeSort(arr []int)  int {
    
  mid := len(arr) / 2

  return mid
}

func mergeArray(first []int, second []int) []int {
    
   mix := make([]int, len(first) + len(second))
   i := 0
   j := 0
   k := 0

   for i < len(first) && j < len(second) {
     
     if(first[i] < second[j]) {
        mix[k] = first[i]
        i++
      } else {
        mix[k] = second[j]
        j++
      }

      k++
   }

   for i < len(first) {
     mix[k] = first[i]
     i++
     k++
   }

   for j < len(second) {
     mix[k] = second[j]
     j++
     k++
   }

   return mix
}
