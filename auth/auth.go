package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Team-Podo/podoting-server/utils"
	"google.golang.org/api/option"
	"log"
	"os"
)

var Firebase fb

type fb struct {
	app *firebase.App
}

func init() {
	path := os.Getenv("FIREBASE_APPLICATION_CREDENTIALS")
	fmt.Println("firebase application path:", path)
	if path == "" {
		path = "firebase-sdk.json"
	}

	opt := option.WithCredentialsFile(utils.RootPath() + path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	} else {
		log.Println("Firebase App Initialized")
	}

	Firebase.app = app
}

func (f *fb) SearchUser(uid string) (*auth.UserRecord, error) {
	client, err := f.app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v\n", err)
	}

	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v\n", err)
	}

	return user, nil
}

func (f *fb) CreateCustomToken(uid string) (string, error) {
	client, err := f.app.Auth(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		return "", fmt.Errorf("error minting custom token: %v\n", err)
	}

	return token, nil
}

func (f *fb) VerifyToken(token string, ctx context.Context) (*auth.Token, error) {
	client, err := f.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v\n", err)
	}

	decodedToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v\n", err)
	}

	return decodedToken, nil
}
