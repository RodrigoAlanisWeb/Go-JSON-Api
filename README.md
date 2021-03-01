# Golang JSON API
## A **Golang** json api with **mux** and **http**
### My api not works with a db like mongodb or mysql , my api works with a JSON variable , the json variable is modified It is modified depending on the action of the client

---

<!-- Ul -->
## Packages

* gorilla/mux
* net/http
* fmt   
* strconv
* encoding/json
* io/ioutil

---

## Routes

```go
r.HandleFunc("/", index)
r.HandleFunc("/tasks", getTasks)
r.HandleFunc("/create", createTask).Methods("POST")
r.HandleFunc("/task/{id}", getTask)
r.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
r.HandleFunc("/update/{id}", updateTask).Methods("PUT")
```
---

### By: Rodrigo Alanis Web | Web Developer
### WebSite: [rodrigoalanis.com](https://portafolio-ra.herokuapp.com/)
### GitHub: [RodrigoAlanisWeb](https://github.com/RodrigoAlanisWeb)