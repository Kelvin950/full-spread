package firebaseclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"bytes"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/kelvin950/spread/internals/core/domain"
	"google.golang.org/api/option"
)

type IFirebaseClient interface {
	CreateUser(email, password string) (string, error)
	LoginUser(email, password string) (string, error)
}

type FirebaseClient struct {
	firebase *firebase.App
	ApiKey   string
}

type varRes struct {
	UID string `json:"localId"`
}

func NewFirebaseClient(fapikey string) (*FirebaseClient, error) {

	absPath := "../confusionserver-d01b1-firebase-adminsdk-udbmv-4c13b154f0.json"

	fb, err := firebase.NewApp(context.TODO(), &firebase.Config{}, option.WithCredentialsFile(absPath))
	return &FirebaseClient{
		firebase: fb,
		ApiKey:   fapikey,
	}, err
}

func (f FirebaseClient) CreateUser(email, password string) (string, error) {

	authclient, err := f.firebase.Auth(context.TODO())

	if err != nil {
		return "", err
	}

	usertocreate := (&auth.UserToCreate{}).Email(email).Password(password).EmailVerified(false)

	newUser, err := authclient.CreateUser(context.TODO(), usertocreate)
	if err != nil {

		return "", domain.ApiError{
			Code:   http.StatusBadRequest,
			ErrVal: err,
		}
	}
	return newUser.UID, err
}

func (f FirebaseClient) LoginUser(email, password string) (string, error) {
	data := map[string]string{"email": email, "password": password}
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	// Make POST request with JSON data
	resp, err := http.Post(fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", f.ApiKey), "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to login, status code: " + resp.Status)
	}
	defer resp.Body.Close()

	var ll varRes

	err = json.NewDecoder(resp.Body).Decode(&ll)

	if err != nil {
		return "", err
	}

	return ll.UID, nil

}
