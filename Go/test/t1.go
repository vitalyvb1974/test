package main
import "fmt"

func main() {
  fmt.Println("Hello!!")
  var i,j int
  i = 5
  j = 10
  fmt.Printf("%d,%d\n", i,j)
  i,j=j,i
  fmt.Printf("%d,%d\n", i,j)

  var arr [5]int
  arr[0]=1
  arr[1]=2
  arr[2]=3
  arr[3]=4
  arr[4]=5

  for _, val := range arr{
    fmt.Printf("%d\n", val)
  }
}