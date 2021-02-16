package videos

import "fmt"

// VideoID is an identifier for a Video
type VideoID string

// Video is a representation of a video
type Video struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// VideoRepository defines all of the bahaviours required for a respository of Videos
type VideoRepository interface {
	Store(*Video) error
	Retrieve(string) (*Video, error)
}

// InMemoryVideoRepository is VideoRepository with no concrete persistance. When the app stops running, all stored Videos are gone forever
type InMemoryVideoRepository struct {
	storage map[string]*Video
}

// NewInMemoryRepository is a factory for InMemoryVideoRepository
func NewInMemoryRepository() InMemoryVideoRepository {

	s := make(map[string]*Video)

	return InMemoryVideoRepository{
		storage: s,
	}
}

// Store adds the provided video to storage
func (ivr InMemoryVideoRepository) Store(v *Video) error {
	ivr.storage[v.ID] = v

	return nil
}

// Retrieve finds the video with the supplied id in storage and returns it
func (ivr InMemoryVideoRepository) Retrieve(id string) (*Video, error) {
	v, ok := ivr.storage[id]

	if !ok {
		return nil, fmt.Errorf("no video with id %s in storage", id)
	}

	return v, nil
}
