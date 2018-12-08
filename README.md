# jsend

jsend is a Go package that implements the JSend [specification](https://labs.omniti.com/labs/jsend).

## Installation

`go get github.com/joaodlf/jsend`

(no dependencies required)

## Example

```
package main

import (
	"github.com/joaodlf/jsend"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w)
	})

	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.Message("Success! I default to a status code of 200!"),
			jsend.Data(map[string]interface{}{
				"fail": false,
			}),
		)
	})

	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.StatusCode(400),
			jsend.Message("Fail!"),
			jsend.Data(map[string]interface{}{
				"fail": true,
			}),
		)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.StatusCode(500),
			jsend.Message("Error! (The 'code' field is optional here)"),
			jsend.Code(1),
		)
	})

	http.ListenAndServe(":8001", nil)
}
```


