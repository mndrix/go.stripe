package stripe

import (
	"net/http"

	"golang.org/x/net/context"
)

func getHttpClient(ctx context.Context) *http.Client {
	return http.DefaultClient
}
