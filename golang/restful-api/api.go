package main

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  // "strconv"
  // golang+mux的性能大约比Python+Flask快20-30%
  "github.com/gorilla/mux"
)

// struct 成员首字母必须大写，否则就是unexport变量
type User struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Age int `json:"age"`
}
var users []User

func get(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")

  user_id := pathParams["user_id"];
  fmt.Printf("GET user_id=%v\n", user_id);
  ret := 0;
  if (ret == -1) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(`{"message": "need a crcode"}`))
  }

  w.WriteHeader(http.StatusOK)
  // w.Write([]byte("{code:0, users:[1,2,3]}"))
  json.NewEncoder(w).Encode(&users)
}

func post(w http.ResponseWriter, r *http.Request) {
  var user User
  user.Id = len(users) + 1
  fmt.Printf("POST user_id=%v\n", user.Id)
  
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {

    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  users = append(users, user)

  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(&user)
  // w.Write([]byte(`{"Created": "1"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  user_id := pathParams["user_id"]
  fmt.Printf("PUT user_id=%v\n", user_id)

  var user = users[0]
  user.Age = 9

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusAccepted)
  json.NewEncoder(w).Encode(&user)
}

func delete(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  user_id := pathParams["user_id"]
  fmt.Printf("DELETE user_id=%v\n", user_id)
  // remove the last item
  users = users[:len(users)-1]
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(&users)
}


func reg_router() {
  r := mux.NewRouter()

  api := r.PathPrefix("/api/v1").Subrouter()
  api.HandleFunc("/users", get).Methods(http.MethodGet)
  api.HandleFunc("/users", post).Methods(http.MethodPost)
  api.HandleFunc("/users/{user_id}", put).Methods(http.MethodPut)
  api.HandleFunc("/users/{user_id}", delete).Methods(http.MethodDelete)

  log.Fatal(http.ListenAndServe(":9001", r))
}

func main() {
  fmt.Printf("Listening 9001...\n");
  reg_router()
}