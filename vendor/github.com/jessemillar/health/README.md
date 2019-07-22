# health-check
A small Go package for quickly adding a `/health` endpoint to RESTful APIs

## Installation
```
go get github.com/jessemillar/health
```

## Usage
```
import "github.com/jessemillar/health"
```

```
func main() {
	port := ":8000"
	router := echo.New()

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.Start(port)
}
```
