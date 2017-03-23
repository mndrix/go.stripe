// +build appengine

package stripe

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

func getHttpClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}
