package application

import (
	"errors"
	"log"
	"github.com/unmsmfisi-socialapplication/social_app/internal/post/domain"

)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrIncompleteData = errors.New("incomplete data")
	ErrPostNotFound = errors.New("post not found")
)

type PostUseCaseInterface interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
	GetUserLocation() (*domain.Location, error)
	GetMultimedia(postId int64) ([]byte, error)
}

type PostRepository interface {
	CreatePost(post domain.CreatePost) (*domain.Post, error)
	GetMultimedia(postId int64) ([]byte, error)
	UpdatePost(postId int64, post domain.CreatePost) (*domain.Post, error)
}

type PostUseCase struct {
	repo PostRepository
}

func NewPostUseCase(r PostRepository) *PostUseCase {
	return &PostUseCase{repo: r}
}
//create
func (l *PostUseCase) CreatePost(post domain.CreatePost) (*domain.Post, error) {
	dbPost, err := l.repo.CreatePost(post)

	if dbPost == nil {
		return dbPost, err
	}

	return dbPost, nil
}
//update
func (l *PostUseCase) UpdatePost(postId int64, post domain.CreatePost) (*domain.Post, error) {
    dbPost, err := l.repo.UpdatePost(postId, post)

    if dbPost == nil {
        return dbPost, err
    }

    return dbPost, nil
}

//others
func (l *PostUseCase) GetMultimedia(postId int64) ([]byte, error) {

    // Get the multimedia data from the repository
    multimedia, err := l.repo.GetMultimedia(postId)

    // Handle errors
    if err != nil {
		log.Println("error: ", err)
    }

    return multimedia, nil
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

