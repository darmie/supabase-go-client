package supabase

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/supabase-community/gotrue-go/endpoints"
	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase/postgrest-go"
)

type Client struct {
	client  *http.Client
	Auth    *endpoints.Client
	Storage *storage_go.Client
	DB      *postgrest.Client
}

func NewClient(projectRef string, apiKey string) *Client {
	dbURL, err := url.Parse(fmt.Sprintf("%s.superbase.co%s", projectRef, "/rest/v1"))
	if err != nil {
		panic(err)
	}
	storageURL, err := url.Parse(fmt.Sprintf("%s.superbase.co%s", projectRef, "/storage/v1"))
	if err != nil {
		panic(err)
	}

	return &Client{
		DB:      postgrest.NewClient(dbURL.String(), "", map[string]string{"apiKey": apiKey, "Authorization": fmt.Sprintf("Bearer %s", apiKey)}),
		Storage: storage_go.NewClient(storageURL.String(), apiKey, nil),
		Auth:    endpoints.New(projectRef, apiKey),
	}
}
