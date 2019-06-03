// This initial file was created via Github Web Editor to go get it later on - Does it work?
// go get -u github.com/tinysammy/play/go/json
// yes, the previous commit was able to be "pulled via go get
// now let's continue via local copy - interestingly the upper files like README.md under
// play/READM.md were pulled too

package main

import "fmt"
// let's use an alias to import our package
import in  "github.com/tinysammy/play/go/json/input"
import out "github.com/tinysammy/play/go/json/output"

func main() {
  fmt.Println("package Hello main")
  in.Hello()
  out.Hello()
}
