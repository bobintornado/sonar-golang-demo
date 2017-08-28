package main
import (
	_ "crypto/md5"
	"fmt"
	"os"
)
func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
