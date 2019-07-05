package mux

// import(
// 	"github.com/gorilla/mux"
// 	"log"
// 	"net/http"
// )
// func TestM(){
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", handler)
// 	r.Use(simpleMw)
	
// 	api := r.PathPrefix("/api").Subrouter()
// 	api.HandleFunc("/things", apiThings)
// 	api.HandleFunc("/things/compress", apiCompressAllTheThings)

// 	http.ListenAndServe(":3000", r)
// }

// func handler(){
// 	log.Fatalln("aaa")
// }
// func apiThings(){
// 	log.Fatalln("apiThings")
// }
// func apiCompressAllTheThings(){
// 	log.Fatalln("apiCompressAllTheThings")
// }
// func simpleMw(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Do stuff here
//         log.Println(r.RequestURI)
//         // Call the next handler, which can be another middleware in the chain, or the final handler.
//         next.ServeHTTP(w, r)
//     })
// }