package comments

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Comment struct {
	Key     *datastore.Key `json:"id" datastore:"-"`
	Body    string         `json:"body"`
	Author  string         `json:"author"`
	Created time.Time      `json:"created"`
}

func handleComments(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	switch r.Method {
	case "GET":
		handleGetComments(ctx, w, r)
	case "POST":
		handleAddComment(ctx, w, r)
	}
}

func handleAddComment(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.Created = time.Now()
	commentKey := datastore.NewIncompleteKey(ctx, "Comment", nil)
	commentKey, err = datastore.Put(ctx, commentKey, &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.Key = commentKey
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handleGetComments(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	comments := make([]Comment, 0)
	keys, err := datastore.NewQuery("Comment").Order("-Created").GetAll(ctx, &comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, key := range keys {
		comments[i].Key = key
	}
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
