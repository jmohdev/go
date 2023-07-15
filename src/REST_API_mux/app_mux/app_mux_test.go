package app_mux

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// Mock-up class: http server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// get response from Mock-up http server
	// ts.URL: /app.go + "/" from app.go defined "func NewHandler" Index path
	resp, err := http.Get(ts.URL)
	// Check assert(respone) No Error
	assert.NoError(err)
	// Check assert(resp.StatusCode)
	// resp.StatusCode(src) ==(Equal) http.StatusOK(expected result)
	// http.StatusOK = 200
	assert.Equal(http.StatusOK, resp.StatusCode)

	// Read response body
	data, _ := ioutil.ReadAll(resp.Body)
	// Check assert(string(resp.Body))
	// string(data)(src) ==(Equal) "Hello world"(expected result)
	assert.Equal("Hello world", string(data))
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	// Mock-up class: http server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// get response from Mock-up http server
	// app.go Not defined "/users"path: call parent path("/")
	resp, err := http.Get(ts.URL + "/users")
	// Check assert(respone) No Error
	assert.NoError(err)
	// Check assert(resp.StatusCode)
	// resp.StatusCode(src) ==(Equal) http.StatusOK(expected interface)
	// http.StatusOK = 200
	assert.Equal(http.StatusOK, resp.StatusCode)

	// Read response body
	data, _ := ioutil.ReadAll(resp.Body)
	// Check assert(string(resp.Body))
	// string(data)(src) >=(Contains) "Get UserInfo"(containd interface)
	assert.Equal(string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 89")

	resp, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 56")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// POST Method
	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"name":"jmoh", "email":"jmoh.developer@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	// strconv.Itoa: Integer → String
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.Name, user2.Name)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Not Delete Method in http Default
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 1")
	// check resp.body with log
	log.Print(string(data))

	// POST Method
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"name":"jmoh", "email":"jmoh.developer@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID: 1")
	// check resp.body with log
	log.Print(string(data))
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "name":"updated", "email":"updated"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 1")

	// POST Method
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"name":"jmoh", "email":"jmoh.developer@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "name":"json"}`, user.ID)

	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("json", updateUser.Name)
	assert.Equal(user.Email, updateUser.Email)
}

func TestUsers_WithUserData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"name":"jmoh", "email":"jmoh.developer@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"name":"json", "email":"json@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	users := []*User{}
	json.NewDecoder(resp.Body).Decode(&users)
	// data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.Equal(2, len(users))
	// assert.NotZero(len(data))
}
