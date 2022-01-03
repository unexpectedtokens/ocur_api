package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/unexpectedtokens/ocur_api/router/routes"
)


func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Write([]byte("application is up and running"))
}



func SetUpRoutes(){
	router := httprouter.New()
	router.GET("/events", routes.GetEvents)
	router.POST("/events", routes.CreateEvent)
	router.GET("/events/:id", routes.GetEvent)
	router.GET("/test/", ping)

	_cors := cors.Options{
        AllowedMethods: []string{
			"POST", 
			"OPTIONS", 
			"GET", 
			"PUT", 
			"UPDATE", 
			"PATCH", 
			"HEAD", 
			"DELETE",
		},
        AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
    }
	handler := cors.New(_cors).Handler(router)
	fmt.Println("Setting op listening on port 8080")
	panic(http.ListenAndServe("localhost:8080", handler))
}