package dependency

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "The Name Of The Lord Is A Strong Tower.The Righteous Run Into It!!")
}
