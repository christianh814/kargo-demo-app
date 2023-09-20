# Simple Go
A simple HTTP application written in Go

To enable the `/greet` route modify:

* index.tmpl
* app.go
* handlers.go

You will also need to create a configmap/secret and mount it under
`/app/config` with the filename `config.json`. It should look like this

```json
{
	"message" : "some string"
}
```
