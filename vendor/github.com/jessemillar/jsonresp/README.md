# json-response
JSON wants returned values to be in key-value format. This package removes a couple lines of code I got tired of copypasta-ing between projects.

## Installation
```
go get github.com/jessemillar/jsonresp
```

## Example Usage
```
import "github.com/jessemillar/jsonresp"
```
```
func SampleFunction(writer http.ResponseWriter, request *http.Request) {
	err := functionThatReturnsAnError()
	if err != nil {
		jsonresp.New(writer, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}
}
```
