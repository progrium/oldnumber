package main

import (
	"log"
	"net/http"
	"os"
)

var voiceTwiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say voice="woman">Hello. Jeff Lindsay no longer receives calls at this number. His new number is 408 218 0596.</Say>
    <Gather timeout="10" finishOnKey="1" numDigits="1">
        <Say voice="woman">Press 1 to have this number sent to your phone via text. Again, Jeff Lindsay's new number is 408 218 0596.</Say>
    </Gather>
</Response>
`

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/xml")
		w.Write([]byte(voiceTwiml))
	})

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
