package routes

import "net/http"

func HomeHandler(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("Hello World"))
}
