package consumer

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

const existantID string = "c9fdf481-81c2-4e49-99c0-dcb3b0200bcc"
const nonExistantID string = "47766cc5-a5db-4fcf-a475-62e363e06e3c"

const videoName string = "terminator"
const videoDescription string = "an old sci fi horror film"

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)

var pact dsl.Pact

func TestMain(m *testing.M) {
	pact = setupPact()

	defer pact.Teardown()

	code := m.Run()

	pact.WritePact()

	os.Exit(code)
}

func setupPact() dsl.Pact {
	return dsl.Pact{
		Consumer: "videoService",
		Provider: "ratingService",
		Host:     "localhost",
		LogLevel: "DEBUG",
		PactDir:  pactDir,
	}

}

func TestConsumer_VideoExists(t *testing.T) {

	var videoExists = func() (err error) {

		subject := HttpConsumer{
			Port: pact.Server.Port,
		}

		v, err := subject.GetVideo(existantID)

		if v.ID != existantID {
			return fmt.Errorf("id mismatch, wanted %s, got %s", existantID, v.ID)
		}

		if v.Name != videoName {
			return fmt.Errorf("name mismatch, wanted %s, got %s", videoName, v.Name)
		}

		if v.Description != videoDescription {
			return fmt.Errorf("description mismatch, wanted %s, got %s", videoDescription, v.Description)
		}

		return err
	}

	pact.
		AddInteraction().
		Given(`video "terminator" exists`).
		UponReceiving(`a request to retrieve "terminator"`).
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String("/videos"),
			Query:  dsl.MapMatcher{"id": dsl.String(existantID)},
		}).
		WillRespondWith(dsl.Response{
			Status:  http.StatusOK,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.String(fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s"}`, existantID, videoName, videoDescription)),
		})

	if err := pact.Verify(videoExists); err != nil {
		log.Fatalf("error on verify: %v", err)
	}
}

func TestConsumer_VideoNotExists(t *testing.T) {

	var videoNotExists = func() (err error) {
		subject := HttpConsumer{
			Port: pact.Server.Port,
		}

		_, err = subject.GetVideo(nonExistantID)

		if err == nil {
			log.Fatalf("was expecting an error")
		}

		expectedErr := fmt.Sprintf("the video of id %s was not found", nonExistantID)

		if err.Error() != expectedErr {
			log.Fatalf("not the expected error")
		}

		return nil
	}

	pact.
		AddInteraction().
		Given(`video "scream" does not exist`).
		UponReceiving(`a request to retrieve "scream"`).
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String("/videos"),
			Query:  dsl.MapMatcher{"id": dsl.String(nonExistantID)},
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusNotFound,
		})

	if err := pact.Verify(videoNotExists); err != nil {
		log.Fatalf("error on verify: %v", err)
	}
}
