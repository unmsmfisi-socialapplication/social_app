package application

import (
	"errors"
	"log"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
}

type PostRepository interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
}

type PostUseCase struct {
	repo PostRepository
}

func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}

func (l *PostUseCase) CreatePost(post domain.CreatePost) (*domain.Post, error) {
	dbPost, err := l.repo.CreatePost(post)

	if dbPost == nil {
		return dbPost, err
	}

	return dbPost, nil
}

func getLatitudeFromGeolocationAPI() (float64, error) {
    // Implement the function to get the user's latitude from the geolocation API
    var latitude float64
    var err error

    // Simulate the geolocation API call and get the latitude
    latitude, err = GeolocationAPICall()

    if err != nil {
        // Log the error to a file or send it to an error reporting service
        log.Println("Error getting latitude:", err)
        return 0.0, err
    }

    return latitude, nil
}

func getLongitudeFromGeolocationAPI() (float64, error) {
    // Implement the function to get the user's longitude from the geolocation API
    var longitude float64
    var err error

    // Simulate the geolocation API call and get the longitude
    longitude, err = GeolocationAPICall()

    if err != nil {
        // Log the error to a file or send it to an error reporting service
        log.Println("Error getting longitude:", err)
        return 0.0, err
    }

    return longitude, nil
}

// Function geolocation API call
func GeolocationAPICall() (float64, error) {
    // Simulate the API call and return latitude (for example, 37.7749) and no error
    return 37.7749, nil
}


func (l *PostUseCase) GetUserLocation() (*domain.Location, error) {
    latitude, err := getLatitudeFromGeolocationAPI()
    if err != nil { // Check for an error from getLatitudeFromGeolocationAPI()
        return nil, err
    }

    longitude, err := getLongitudeFromGeolocationAPI()
    if err != nil { // Check for an error from getLongitudeFromGeolocationAPI()
        return nil, err
    }

    location := &domain.Location{
        Latitude: latitude,
        Longitude: longitude,
    }

    return location, nil
}

