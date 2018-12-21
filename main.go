package main;

import 
(
"fmt"
"github.com/rajeshpachar/hellomod/hellotest"
"github.com/rajeshpachar/hellomod/child"
)

func main(){
fmt.Println("we are inside main func now");
hellotest.SayHello("my main calling me");

fmt.Println("now main is running child###");
child.HelloChild()
}
