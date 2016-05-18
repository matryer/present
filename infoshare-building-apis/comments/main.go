package comments

import "net/http"

func init() {
	http.HandleFunc("/comments", handleComments)
}
