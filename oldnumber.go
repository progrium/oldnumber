package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var spokenNumber = "5 1 2, 7 8 5, 1 1 9 4"
var textNumber = "512 785 1194"

var voiceTwiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say voice="woman">Hello. Jeff Lindsay no longer receives calls at this number. His new number is, {SPOKEN}.</Say>
	<Sms>My new number: {TEXT}</Sms>
	<Say voice="woman">If you can receive them, I just texted the number to your phone. Again, his new number is, {SPOKEN}. Goodbye.</Say>
</Response>
`

var smsTwiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Sms>I don't get texts at this number anymore. My new number: {TEXT}</Sms>
</Response>
`

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	voiceTwiml = strings.Replace(voiceTwiml, "{SPOKEN}", spokenNumber, -1)
	voiceTwiml = strings.Replace(voiceTwiml, "{TEXT}", textNumber, -1)
	smsTwiml = strings.Replace(smsTwiml, "{TEXT}", textNumber, -1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/xml")
		if r.URL.Path == "/voice" {
			w.Write([]byte(voiceTwiml))
		} else {
			w.Write([]byte(smsTwiml))
		}
	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
