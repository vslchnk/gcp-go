package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	b "google.golang.org/api/cloudbilling/v1"
)

func main() {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"https://www.googleapis.com/auth/cloud-billing"},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)

	computeService, err := b.New(client)
	if err != nil {
		log.Fatal(err)
	}

	l, err := computeService.BillingAccounts.Projects.List("billingAccounts/0113EB-D2D3AA-FA24E8").Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l)
}
