package main
import "fmt"
import "newpack"
func main(){
var n int
fmt.Scan(&n)
result:=newpack.Square(n)
fmt.Printf("Square of %d is %d",n,result)
}