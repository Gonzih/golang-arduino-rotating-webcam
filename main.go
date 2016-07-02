package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var tty string

var template string = `
<html>
<body>
<a href='/change?value=0'>0</a>
<a href='/change?value=45'>45</a>
<a href='/change?value=90'>90</a>
<a href='/change?value=135'>135</a>
<a href='/change?value=180'>180</a>
</form>
</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, template)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query()["value"]
	if value != nil {
		v := value[0]
		log.Printf("Value is %v\n", v)

		out := []byte(v)
		err := ioutil.WriteFile(tty, out, 1660)

		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", 302)
	} else {
		io.WriteString(w, "please provide value\n")
	}
}

func main() {
	var ttyArg = flag.String("tty", "/dev/ttyACM1", "tty that should be used to communicate with Arduino")
	flag.Parse()
	tty = *ttyArg

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/change", changeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
