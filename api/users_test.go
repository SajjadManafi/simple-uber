package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SajjadManafi/simple-uber/contract"
	"github.com/SajjadManafi/simple-uber/internal/token"
	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"gender":    user.Gender,
				"email":     user.Email,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, "rider", time.Minute)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUserCreate(t, recorder.Body, user)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username":  "InvalidUser#1",
				"password":  password,
				"full_name": user.FullName,
				"gender":    user.Gender,
				"email":     user.Email,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "InvalidUser#1", "rider", time.Minute)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, TestServer.tokenMaker)

			TestServer.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)

		})
	}
}

func TestGetUser(t *testing.T) {
	var ID int32
	user, _ := randomUser(t)

	testCases := []struct {
		name          string
		setupStore    func(store contract.Store)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupStore: func(store contract.Store) {
				arg := models.CreateUserParams{
					Username:       user.Username,
					FullName:       user.FullName,
					HashedPassword: user.HashedPassword,
					Gender:         user.Gender,
					Email:          user.Email,
				}
				user, err := store.CreateUser(context.Background(), arg)
				require.NoError(t, err)
				ID = user.ID
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, "rider", time.Minute)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUserGet(t, recorder.Body, user)
			},
		},
		{
			name: "NotFound",
			setupStore: func(store contract.Store) {
				err := store.DeleteUser(context.Background(), ID)
				require.NoError(t, err)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, "rider", time.Minute)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			setupStore: func(store contract.Store) {
				ID = 0
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, "rider", time.Minute)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := TestServer.store
			server := NewTestServer(t, store)

			recorder := httptest.NewRecorder()

			tc.setupStore(server.store)

			url := fmt.Sprintf("/api/users/%d", ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			log.Println(recorder.Body)
			tc.checkResponse(recorder)
		})
	}

}

func randomUser(t *testing.T) (user models.User, password string) {
	password = util.RandomString(6)
	hashesdPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = models.User{
		Username:       util.RandomUsername(),
		FullName:       util.RandomUsername(),
		Gender:         util.RandomGender(),
		Email:          util.RandomEmail(),
		HashedPassword: hashesdPassword,
	}

	return

}

func requireBodyMatchUserCreate(t *testing.T, body *bytes.Buffer, user models.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser models.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)

	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}

func requireBodyMatchUserGet(t *testing.T, body *bytes.Buffer, user models.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser models.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)

	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.NotEmpty(t, gotUser.HashedPassword)
}
