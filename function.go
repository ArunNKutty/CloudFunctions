package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

var substrings = []string{"Google","Oracle","Microsoft","Amazon"}
// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func CopyRight(w http.ResponseWriter, r *http.Request) {
	queryParamDisplayHandler(w, r)
}


func checkSubstrings(inputString string) string {
	for _, sub := range substrings {
		if strings.Contains(inputString, sub) {
			inputString = strings.ReplaceAll(inputString,sub,"©"+sub)
		}
	}
	return inputString
}

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	inputSentence := req.FormValue("sentence")
	if inputSentence != "" {
		io.WriteString(res, "New Copyrighted Sentence: "+checkSubstrings(inputSentence))
	} else {
		io.WriteString(res, "New Copyrighted Sentence: "+"The Input sentence is empty")
	}
}
