package main
import "os"
func main() {
	os.Chmod("/tmp/somefile", 0777)
	os.Chmod("/tmp/someotherfile", 0600)
	os.OpenFile("/tmp/thing", os.O_CREATE|os.O_WRONLY, 0666)
	os.OpenFile("/tmp/thing", os.O_CREATE|os.O_WRONLY, 0600)
}
