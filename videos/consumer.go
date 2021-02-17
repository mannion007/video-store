package consumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Video struct {
	ID          string `json:"id" pact:"example=123"`
	Name        string `json:"name" pact:"example=the matrix"`
	Description string `json:"description" pact:"example=the most overrated film, ever"`
}

type Consumer interface {
	GetVideo(id string) *Video
}

type HttpConsumer struct {
	Port int
}

func (h HttpConsumer) GetVideo(id string) (*Video, error) {

	url := fmt.Sprintf("http://localhost:%d/videos", h.Port)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("the video of id %s was not found", id)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var v *Video

	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, fmt.Errorf("getvideo: failed to unmarshal response %v", err)
	}

	return v, nil
}
