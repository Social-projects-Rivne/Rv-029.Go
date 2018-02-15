package middlewares

import (
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type authorithedHandler struct {
	Handler func(w http.ResponseWriter, r *http.Request)
}

func (a *authorithedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler(w, r)
}

// Check if user authenticated
func TestAuthenticatedMiddleware(t *testing.T) {
	handler := &authorithedHandler{
		Handler: func(w http.ResponseWriter, r *http.Request) {
			userContext := r.Context().Value("user")
			fmt.Fprint(w, userContext)
		},
	}

	user := &models.User{UUID: gocql.TimeUUID()}

	jwtToken, err := jwt.GenerateToken(user)
	if err != nil {
		t.Fatal("Token wasn't generated")
	}

	request := httptest.NewRequest("GET", "http://localhost/", nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

	w := httptest.NewRecorder()
	AuthenticatedMiddleware(handler)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if user.UUID.String() != string(body) {
		t.Error("Context should have a correct user id")
	}
}
