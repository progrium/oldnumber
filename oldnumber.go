package main

import (
	"log"
	"net/http"
	"os"
)

var voiceTwiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say voice="woman">Hello. Jeff Lindsay no longer receives calls at this number. His new number is 4 0 8, 2 1 8, 0 5 9 6.</Say>
	<Sms>My new number: 408 218 0596</Sms>
	<Say voice="woman">If you can receive them, I just texted the number to you. Again, his new number is 4 0 8, 2 1 8, 0 5 9 6. Goodbye.</Say>
</Response>
`

var smsTwiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Sms>I don't get texts at this number anymore. My new number: 408 218 0596</Sms>
</Response>
`

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

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
