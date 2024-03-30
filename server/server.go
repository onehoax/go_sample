package server

import (
	"flag"
	"html/template"
	"net/http"
)

var Addr *string = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ *template.Template = template.Must(template.New("qr").Parse(TemplateStr))

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}
