package main
import "fmt"
func main(){
fmt.Println("palindrome program  checker")
var number,temp,reverse,rem int =45454,0,0,0
temp=number
for{
rem=number%10
reverse=reverse*10+rem
number/=10
if(number==0){
break
}
}
if(temp==reverse){
fmt.Printf("Number%d is a Palindrome",temp)
}else{
fmt.Printf("Number%d is not a Palindrome",temp)
}
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      