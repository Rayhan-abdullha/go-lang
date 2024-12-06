// package main
// import (
//     "encoding/json"
//     "net/http"
// )
// type User struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
// }

// func main() {
//     http.HandleFunc("/user", handleJSON)
//     http.ListenAndServe(":8080", nil)
// }

// func handleJSON(w http.ResponseWriter, r *http.Request) {
//     var user User
//     json.NewDecoder(r.Body).Decode(&user)
//     w.Write([]byte("Received: " + user.Name))
// }
