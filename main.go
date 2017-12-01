package main
import(
	"fmt"
	"net/http"
)

func index(w http.ResponseWrite, r *http.Request){
	fmt.Fprintf(w, "by dux")
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
