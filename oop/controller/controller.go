package controller

import "net/http"

type PresentsController struct{}

// The `ReadMany` method would be mapped to `GET /presents`.
// We aren’t mapping the methods to a pointer of the controller, because we don’t need to
// share any state between our RESTful methods.
func (p PresentsController) ReadMany(w http.ResponseWriter, r *http.Request) {
	// TODO: handle read many request
}

type RestManyReader interface {
	ReadMany(w http.ResponseWriter, r *http.Request)
}

type RestReader interface {
	Read(w http.ResponseWriter, r *http.Request)
}

type ReastCreator interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type RestUpdater interface {
	Update(w http.ResponseWriter, r *http.Request)
}

type RestReplacer interface {
	Replace(w http.ResponseWriter, r *http.Request)
}

type RestDeleter interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

// in a helper function, we can check to see if a controller implements that method
// and make the appropriate mapping:

type Router interface {
	Route(path string, method string, handler interface{})
}

func route(path string, controller interface{}, router Router) {

	if c, ok := controller.(RestReader); ok {
		router.Route(path+"/{id}", "GET", http.HandlerFunc(c.Read))
	}

	if c, ok := controller.(RestManyReader); ok {
		router.Route(path, "GET", http.HandlerFunc(c.ReadMany))
	}

	if c, ok := controller.(ReastCreator); ok {
		router.Route(path, "POST", http.HandlerFunc(c.Create))
	}

	if c, ok := controller.(RestDeleter); ok {
		router.Route(path+"/{id}", "DELETE", http.HandlerFunc(c.Delete))
	}
	if c, ok := controller.(RestReplacer); ok {
		router.Route(path+"/{id}", "PUT", http.HandlerFunc(c.Replace))
	}
	if c, ok := controller.(RestUpdater); ok {
		router.Route(path+"/{id}", "PATH", http.HandlerFunc(c.Update))
	}

	router.Route(path, "*", http.NotFoundHandler())
}

// Other ideas include:

// - Extracting the ID path parameter and changing the methods to take it as a parsed
//   argument: `Delete(w http.ResponseWriter, r *http.Request, id string)`
//
// - Count the number of mapping made, and at the end of Route make sure at least one
//   was done. If not, panic — because the developer probably made a mistake.
//
// - If you’re looking for a router that supports path parameters (like `{id}`) then
//   check out Gorilla’s mux package.
//
// - Add another method and interface pair that lets your controller map additional
//   routes that go beyond a traditional RESTful design.
 