package http

import (
	productHttp "goshop/internal/product/http"
	userHttp "goshop/internal/user/http"
)

func (s Server) MapRoutes() error {
	api := s.engine.Group("/api")
	productHttp.InitProductRoutes(api, s.db, s.validator)
	userHttp.InitUserRoutes(api, s.db, s.validator)
	return nil
}
