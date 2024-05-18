package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)                          //RECURSION
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  // Recursive call to itself with n-1
  return n * factorial(n-1)
}

func main(){
  testcount(1)
  fmt.Println("Factorial is:")
  fmt.Println(factorial(5))
}
