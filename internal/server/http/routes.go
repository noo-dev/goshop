package http

import "goshop/internal/product/http"

func (s Server) MapRoutes() error {
	api := s.engine.Group("/api")
	http.InitProductRoutes(api, s.db, s.validator)
	return nil
}
