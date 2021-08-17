package router

import "github.com/gorilla/mux"

// GerarRouter retorna um router
func GerarRouter() *mux.Router {
	return mux.NewRouter()
}
