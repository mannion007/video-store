package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mannion007/video-store/videos"
)

func main() {

	s := videos.NewInMemoryRepository()

	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {

		// post only
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		defer r.Body.Close()

		var v *videos.Video

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unable to read request body"))
		}

		err = json.Unmarshal(b, &v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		err = s.Store(v)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to store video"))

			return
		}

		w.WriteHeader(http.StatusCreated)

	})

	http.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {

		// get only
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		defer r.Body.Close()

		id, ok := r.URL.Query()["id"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing id parameter"))

			return
		}

		vid := id[0]

		v, err := s.Retrieve(vid)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("unable to find video"))

			return
		}

		b, err := json.Marshal(v)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to marshal video"))

			return
		}

		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
