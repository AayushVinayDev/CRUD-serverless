package routes

type Route struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters()       {}
func (r *Router) setConfigRouters() {}
func (r *Router) RouterHealth()     {}
func (r *Router) RouterProduct()    {}
func (r *Router) EnableTimeout()    {}
func (r *Router) EnableCORS()       {}
func (r *Router) EnableRecovery()   {}
func (r *Router) EnableRequestID()  {}
func (r *Router) EnableRealIP()     {}
