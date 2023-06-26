package ziface

type IRouter interface {
	PreRouter(request IRequest)
	Handler(request IRequest)
	AfterRouter(request IRequest)
}
