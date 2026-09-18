package manager

import (
	"net/http"
)

type Middleware func(http.Handler)http.Handler

type Manager struct{
	globalMiddlewares []Middleware
}

func NewManager() *Manager{
	return &Manager{
		globalMiddlewares: make([]Middleware,0),
	}
}
func (mngr *Manager)Use(middlewares ...Middleware){
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	// return mngr //builder pattern
}

func (mngr *Manager)With(handaler http.Handler, middlewares ...Middleware)http.Handler{
		h:= handaler
		//middlewares = [logger,Hudai,]
		//huddai(logger(http.handlefunc(getproduct)))
		for _,middleware:= range middlewares{
			h = middleware(h)
		}
		//

		return h
	
}

func (mngr *Manager)WrappedRouter(handaler http.Handler)http.Handler{

	h:=handaler
		for _,middleware:= range mngr.globalMiddlewares{
			h = middleware(h)
		}
		return h
}