// We can still use private data in our debug command!

package main

import (
	"fmt"

	"github.com/vasi/internalpackages"
	"github.com/vasi/internalpackages/internal"
)

func main() {
	f := internalpackages.ParseFile("foo")
	ff := f.(*internal.MyFileImpl)
	fmt.Printf("dump: %s\n", ff.PrivateData)
}
