package controllers

import "github.com/mesbera/go-blog/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUserEntryCount)).Methods("GET")

	//Posts routes

	//Comments routes
	s.Router.HandleFunc("/comments", middlewares.SetMiddlewareJSON(s.CreateComment)).Methods("POST")
}
