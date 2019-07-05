package mux

import (
    "crypto/subtle"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
)

func endAPICall(w http.ResponseWriter, httpStatus int, anyStruct interface{}) {

    result, err := json.MarshalIndent(anyStruct, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(httpStatus)
    w.Write(result)
}

func BasicAuth(w http.ResponseWriter, r *http.Request, username, password, realm string) bool {

    user, pass, ok := r.BasicAuth()

    if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
        w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
        w.WriteHeader(401)
        w.Write([]byte("Unauthorised.\n"))
        return false
    }

    return true
}

func routers() *mux.Router {
    username := "apiuser"
    password := "apipass"

    v2Path := "/v2"
    healthPath := "/health"

    topRouter := mux.NewRouter().StrictSlash(true)
    healthRouter := mux.NewRouter().PathPrefix(healthPath).Subrouter().StrictSlash(true)
    v2Router := mux.NewRouter().PathPrefix(v2Path).Subrouter().StrictSlash(true)

    healthRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        endAPICall(w, 200, "PONG")
    })

    v2Router.HandleFunc("/smallcat", func(w http.ResponseWriter, r *http.Request) {
        endAPICall(w, 200, "Small Meow")
    })

    bigMeowFn := func(w http.ResponseWriter, r *http.Request) {
        endAPICall(w, 200, "Big MEOW")
    }

    v2Router.HandleFunc("/bigcat", bigMeowFn)

    topRouter.PathPrefix(healthPath).Handler(negroni.New(
        /* Health-check routes are unprotected */
        negroni.Wrap(healthRouter),
    ))

    topRouter.PathPrefix(v2Path).Handler(negroni.New(
        negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
            if BasicAuth(w, r, username, password, "Provide user name and password") {
                /* Call the next handler iff Basic-Auth succeeded */
                next(w, r)
            }
        }),
        negroni.Wrap(v2Router),
    ))

    return topRouter
}

func TestMux1() {
    if r := routers(); r != nil {
        log.Fatal("Server exited:", http.ListenAndServe(":3000", r))
    }
}