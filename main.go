package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var tty string

var css = `
<head>
<style>
a {
	border: 1px solid black;
	margin: 5px;
	padding: 5px;
	display: box;
	text-decoration: none;
}

a, a:visited, a:active {
	color:black;
}

a:hove {
	color: grey;
}
</style>
</head>
`

func template() string {
	start := "<html>" + css + "<body>"
	end := "</body></html>"
	middle := ""
	for i := 0; i <= 180; i = i + 15 {
		middle = middle + fmt.Sprintf("<a href='/change?value=%d'>%d</a>", i, i)
	}

	return start + middle + end
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, template())
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
	var ttyArg = flag.String("tty", "/dev/ttyACM0", "tty that should be used to communicate with Arduino")
	flag.Parse()
	tty = *ttyArg

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/change", changeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
