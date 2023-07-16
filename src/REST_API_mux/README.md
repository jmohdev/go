# REST_API

REST: REpresentational State Tranfer  
"Method(GET, POST, PUT, DELETE)+ URI" CRUD manipulating  
※ CRUD: Create Read Update Delete

```bash
├── app
│   ├── app_mux.go
│   └── app_mux_test.go
├── main.go
└── README.md
```
*** 
## /main.go

### func main  
Go의 표준 패키지인 net/http 패키지에서 제공하는 Server, Client 기능 사용
```go
package main

import (
	"REST_API_mux/app_mux"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", app_mux.NewHandler())
}

```
http://localhost:3000/

***
## /app_mux/app_mux.go

```go
package app_mux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Updated User struct for ""(null) manipulating
type UpdateUser struct {
	ID           int       `json:"id"`
	UpdatedName  bool      `json:"updated_name"`
	Name         string    `json:"name"`
	UpdatedEmail bool      `json:"updated_email"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns the route variables for the current request
	vars := mux.Vars(r)
	// strconv.Atoi: String → Integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", id)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", id)
		return
	}
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID: ", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", updateUser.ID)
		return
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func NewHandler() http.Handler {
	// Initialize userMap, lastID
	userMap = make(map[int]*User)
	lastID = 0
	// mux.NewRouter: gorilla/mux
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	// .Methods("GET"): gorilla/mux
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	// .Methods("POST"): gorilla/mux
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	// {id:[0-9]+}: gorilla/mux
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	return mux
}
```

### TEST RESULT : assert.Contains, assert.Equal
- PASS : assert.Contains(string(data), "Get UserInfo")  
	① call "/users", usersHandler  
	② resp.Body = "Get UserInfo by /users/{id}"  
	③ string(data) = "Get UserInfo by /users/{id}"  
	④ Contains "Get UserInfo"
	```go
		// Read response body
		data, _ := ioutil.ReadAll(resp.Body)
		// Check assert(string(resp.Body))
		// string(data)(src) >=(Contains) "Get UserInfo"(containd interface)
		assert.Contains(string(data), "Get UserInfo")
	```
- FAILED : assert.Equal("Hello world", string(data))  
	① call "/users", usersHandler  
	② resp.Body = "Get UserInfo by /users/{id}"  
	③ string(data) = "Get UserInfo by /users/{id}"  
	Error : Not equal
	```go
		// Read response body
		data, _ := ioutil.ReadAll(resp.Body)
		// Check assert(string(resp.Body))
		// string(data)(src) ==(Equal) "Hello world"(expected result)
		assert.Equal("Hello world", string(data))
	```
### TEST RESULT : Not defined "/users/89" path in app.go
```go
func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User ID: 89")
}
```
- FAILED : Not defined "/users/89" path in app.go  
	① TestGetUserInfo call path("/users/89")  
	② Not defined "/users/89" path in app.go  
	③ call parent path("/")  
	④ call indexHandler  
	⑤ reponse = "Hello world"  
	Error : "Hello world" does not contain "User ID: 89"
	```go
	func indexHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	}

	func usersHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get UserInfo by /users/{id}")
	}

	func NewHandler() http.Handler {
		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		mux.HandleFunc("/users", usersHandler)

	return mux
	}
	```
- PASS : Defined "/users/89" path in app.go  
	① TestGetUserInfo call path("/users/89")  
	② Defined "/users/89" path in app.go  
	③ call path("/users/89")  
	④ call getUserInfo89Handler  
	⑤ reponse = "User ID: 89"
	```go
	func indexHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	}

	func usersHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get UserInfo by /users/{id}")
	}

	func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "User ID: 89")
	}

	func NewHandler() http.Handler {
		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		mux.HandleFunc("/users", usersHandler)
		mux.HandleFunc("/users/89", getUserInfo89Handler)

		return mux
	}
	```

## github.com/gorilla/mux

```
go get -u github.com/gorilla/mux
```
~~mux := http.NewServeMux()~~  
mux := mux.NewRouter()

~~mux.HandleFunc("/users/89", getUserInfo89Handler)~~  
mux.HandleFunc("/users/{id:[0-9]+}", getUserInfo89Handler)

~~fmt.Fprint(w, "User ID: 89")~~  
vars := mux.Vars(r)  
fmt.Fprint(w, "User ID: ", vars["id"])

__BEFORE__
```go

func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User ID: 89")
}

func NewHandler() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/users", usersHandler)
    mux.HandleFunc("/users/89", getUserInfo89Handler)

    return mux
}
```
__AFTER__
```go
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
    // mux.Vars returns the route variables for the current request
    vars := mux.Vars(r)
	fmt.Fprint(w, "User ID: ", vars["id"])
}

func NewHandler() http.Handler {
    mux := mux.NewRouter()
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/users", usersHandler)
    mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

    return mux
}
```
### TEST RESULT : StatusCreated
- FAILED  
	① TestCreateUser call path("/users") with POST method  
	② call usersHandler with GET method  
    ③ reponse = "Get UserInfo by /users/{id}"  
	Error: Not equal // expected: 201(Created), actual: 200(OK)
```go
func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// POST Method
	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"name":"jmoh", "email":"jmoh.developer@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
}
```

- PASS  
	① TestCreateUser call path("/users") with POST method  
	② call createUserHandler with POST method  
    ③ reponse = json

```go
// User struct
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// new instance
	user := new(User)
	// check err
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	user.ID = 2
	user.CreatedAt = time.Now()
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func NewHandler() http.Handler {
	// mux.NewRouter: gorilla/mux
    mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
    // .Methods("GET"): gorilla/mux
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	// .Methods("POST"): gorilla/mux
    mux.HandleFunc("/users", createUserHandler).Methods("POST")
	// {id:[0-9]+}: gorilla/mux
    mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}
```
### TEST : Method Transfer
https://github.com/frigus02/RESTer

### TEST : DELETE Method
Mock-up > DELETE Method > (test) POST Method > (test) DELETE Method > (test) Response
```go
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Not Delete Method in http Default
	// Need "http.NewRequest"
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 1")
	// check resp.body with log
	log.Print(string(data))

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
```
```go
mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", id)
		return
	}
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID: ", id)
}
```

### TEST : PUT Method(UPDATE)
Mock-up > PUT Method > (test) POST Method > (test) PUT Method > (test) Compare Data
```go
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
```
```go
mux.HandleFunc("/users", updateUserHandler).Methods("PUT")

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", updateUser.ID)
		return
	}

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}
```

### TEST : Data List
Mock-up > POST Method > (test) POST Method > (test) GET Method > (test) Data List
```go
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
```
```go
func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns the route variables for the current request
	vars := mux.Vars(r)
	// strconv.Atoi: String → Integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID: ", id)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}
```

### ""(null) Data manipulating
```go
// User struct
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Updated User struct for ""(null) Data manipulating
type UpdateUser struct {
	ID           int       `json:"id"`
	UpdatedName  bool      `json:"updated_name"`
	Name         string    `json:"name"`
	UpdatedEmail bool      `json:"updated_email"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
}
```