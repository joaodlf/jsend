# jsend

jsend is a Go package that implements the JSend [specification](https://labs.omniti.com/labs/jsend).

## Installation

```
go get github.com/joaodlf/jsend
```
(no dependencies required)

## Usage

Simply pass a `http.ResponseWriter` to `jsend.Write()`. You can also pass a list of options to tailor your output:

```
jsend.Write(w,
    jsend.Data(map[string]interface{}{
        "user_id": 1,
        "email": "you@domain.com",
    }),
)
```

The above will output the following JSON encoded message:
```
{
    data: {
        email: "you@domain.com",
        user_id: 1
    },
    status: "success"
}
```

## Example

```
package main

import (
	"github.com/joaodlf/jsend"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.Data(map[string]interface{}{
				"user_id": 1,
				"email": "you@domain.com",
			}),
		)
	})

	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.StatusCode(400),
			jsend.Data(map[string]interface{}{
				"fail": true,
			}),
		)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		jsend.Write(w,
			jsend.StatusCode(500),
			jsend.Message("Error! (The 'code' and 'data' fields are optional here)"),
			jsend.Code(1),
			jsend.Data(map[string]interface{}{
				"error": true,
			}),
		)
	})

	http.ListenAndServe(":8001", nil)
}
```


