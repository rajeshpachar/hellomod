package child;

import (
"fmt"
"github.com/rajeshpachar/hellomod/hellotest"
)

func HelloChild(){
fmt.Println("child is looking good ")
hellotest.SayHello("child is calling parent")
}
