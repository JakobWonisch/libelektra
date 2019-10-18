package main

import (
	"net/http"
	"regexp"

	elektra "github.com/ElektraInitiative/go-elektra/kdb"
)

// getFindHandler searches for Keys via a Regex expression.
//
// Arguments:
// 		regex	the regex expression to find keys, URL path param.
//
// Response Code:
//		200 OK if the request was successfull.
// 		400 Bad Request if the REGEX is invalid.
//
// Returns: String array with found keys.
//
// Example: `curl localhost:33333/kdbFind/versi*`
func getFindHandler(w http.ResponseWriter, r *http.Request) {
	query := parseKeyNameFromURL(r)
	regex, err := regexp.Compile(query)

	if err != nil {
		badRequest(w)
		return
	}

	root, err := elektra.NewKey("/")

	if err != nil {
		writeError(w, err) // this should not happen
		return
	}

	handle := getHandle(r)
	ks, err := getKeySet(handle, root)

	if err != nil {
		writeError(w, err)
		return
	}

	results := []string{}

	for _, key := range ks.KeyNames() {
		if regex.MatchString(key) {
			results = append(results, key)
		}
	}

	writeResponse(w, results)
}
