package experiments
import(
	"fmt"
	"testing"
)
type person struct{
	name string
	age int
}

func TestStruct(t *testing.T){
	var p person
	fmt.Println(p)
	p = person{
		name : "wang",// p.name = "wang"
		age : 18,    //p.age = 18
	}
	fmt.Println(p)
}