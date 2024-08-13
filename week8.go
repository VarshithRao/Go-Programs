package main
import "fmt"
import "net/http"
func main(){
http.HandleFunc("/",func(w http.ResponseWriter,r*http.Request){
if r.Method != http.MethodPost{
http.ServeFile(w,r,"form.html")
return
}
err:=r.ParseForm()
if err!=nil{
http.Error(w,"Error parsing form data",http.StatusBadRequest)
return
}
name:=r.Form.Get("name")
email:=r.Form.Get("email")
msg:=r.Form.Get("Message")
fmt.Printf("Name1:%s\n",name)
fmt.Printf("Email1:%s\n",email)
fmt.Printf("msg1:%s\n",msg)
fmt.Fprintf(w,"Received form submission:\nName:%s\nEmail:%s\nMessage:%s\n",name,email,msg)
})
fmt.Println("Serve is running on localhost:2425")
http.ListenAndServe(":2425",nil)
}