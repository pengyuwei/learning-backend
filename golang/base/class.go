package main
import "fmt"

type User struct {
    id int
    name string
    age int
}

func (this *User) update(name string, age int){
    (*this).name = name
    (*this).age = age
}

func main() {
    var user User
    user.update("FF", 11)
    fmt.Printf(user.name)
}
