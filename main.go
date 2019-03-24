package loong

import (
	"context"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/runner-mei/loong/log"
	"github.com/runner-mei/loong/util"
)

const MyContextKey = "my-context-key"

var ErrNotFound = echo.ErrNotFound
var ErrBadArgument = util.ErrBadArgument
var WithHTTPCode = util.WithHTTPCode
var Wrap = util.Wrap

type Error = util.Error
type HTTPError = util.HTTPError

type Context struct {
	echo.Context
	StdContext context.Context

	CtxLogger log.Logger
	LogArray  []string
}

type Result struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data,omitempty"`
	Error    *Error      `json:"error,omitempty"`
	Messages []string    `json:"messages,omitempty"`
}

func (c *Context) ReturnResult(code int, i interface{}, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(code, i)
	}
	return c.JSON(code, &Result{Success: true, Data: i})
}

func (c *Context) ReturnCreatedResult(i interface{}, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(http.StatusCreated, i)
	}
	return c.JSON(http.StatusCreated, &Result{Success: true, Data: i})
}

func (c *Context) ReturnUpdatedResult(i interface{}, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, &Result{Success: true, Data: i})
}

func (c *Context) ReturnDeletedResult(i interface{}, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, &Result{Success: true, Data: i})
}

func (c *Context) ReturnQueryResult(i interface{}, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, &Result{Success: true, Data: i})
}

func (c *Context) ReturnCountResult(i int64, notWrapped ...bool) error {
	if len(notWrapped) > 0 && notWrapped[0] {
		return c.JSON(http.StatusOK, i)
	}
	return c.JSON(http.StatusOK, &Result{Success: true, Data: i})
}

func (c *Context) ReturnError(err error, code ...int) error {
	var httpCode int
	if len(code) > 0 {
		httpCode = code[0]
	}
	if httpCode == 0 {
		if a, ok := err.(interface {
			HTTPCode() int
		}); ok {
			httpCode = a.HTTPCode()
		} else {
			httpCode = http.StatusInternalServerError
		}
	}
	return c.JSON(httpCode, &Result{Success: false, Messages: c.LogArray, Error: util.ToError(err, httpCode)})
}

var _ echo.Context = &Context{}

// HandlerFunc defines a function to serve HTTP requests.
type HandlerFunc func(*Context) error

// MiddlewareFunc defines a function to process middleware.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// Route contains a handler and information for matching against requests.
type Route = echo.Route

// Validator is the interface that wraps the Validate function.
type Validator = echo.Validator

// Renderer is the interface that wraps the Render function.
type Renderer interface {
	Render(io.Writer, string, interface{}, Context) error
}

// Map defines a generic map of type `map[string]interface{}`.
type Map map[string]interface{}

// MIME types
const (
	MIMEApplicationJSON                  = echo.MIMEApplicationJSON
	MIMEApplicationJSONCharsetUTF8       = echo.MIMEApplicationJSONCharsetUTF8
	MIMEApplicationJavaScript            = echo.MIMEApplicationJavaScript
	MIMEApplicationJavaScriptCharsetUTF8 = echo.MIMEApplicationJavaScriptCharsetUTF8
	MIMEApplicationXML                   = echo.MIMEApplicationXML
	MIMEApplicationXMLCharsetUTF8        = echo.MIMEApplicationXMLCharsetUTF8
	MIMETextXML                          = echo.MIMETextXML
	MIMETextXMLCharsetUTF8               = echo.MIMETextXMLCharsetUTF8
	MIMEApplicationForm                  = echo.MIMEApplicationForm
	MIMEApplicationProtobuf              = echo.MIMEApplicationProtobuf
	MIMEApplicationMsgpack               = echo.MIMEApplicationMsgpack
	MIMETextHTML                         = echo.MIMETextHTML
	MIMETextHTMLCharsetUTF8              = echo.MIMETextHTMLCharsetUTF8
	MIMETextPlain                        = echo.MIMETextPlain
	MIMETextPlainCharsetUTF8             = echo.MIMETextPlainCharsetUTF8
	MIMEMultipartForm                    = echo.MIMEMultipartForm
	MIMEOctetStream                      = echo.MIMEOctetStream
)

type Party interface {
	Use(middleware ...MiddlewareFunc)

	// CONNECT implements `Echo#CONNECT()` for sub-routes within the Group.
	CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// DELETE implements `Echo#DELETE()` for sub-routes within the Group.
	DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// GET implements `Echo#GET()` for sub-routes within the Group.
	GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// HEAD implements `Echo#HEAD()` for sub-routes within the Group.
	HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// OPTIONS implements `Echo#OPTIONS()` for sub-routes within the Group.
	OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// PATCH implements `Echo#PATCH()` for sub-routes within the Group.
	PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// POST implements `Echo#POST()` for sub-routes within the Group.
	POST(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// PUT implements `Echo#PUT()` for sub-routes within the Group.
	PUT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// TRACE implements `Echo#TRACE()` for sub-routes within the Group.
	TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route

	// Any implements `Echo#Any()` for sub-routes within the Group.
	Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []*Route

	// Match implements `Echo#Match()` for sub-routes within the Group.
	Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []*Route

	// Group creates a new sub-group with prefix and optional sub-group-level middleware.
	Group(prefix string, middleware ...MiddlewareFunc) Party

	// Static implements `Echo#Static()` for sub-routes within the Group.
	Static(prefix, root string)

	// File implements `Echo#File()` for sub-routes within the Group.
	File(path, file string)

	// Add implements `Echo#Add()` for sub-routes within the Group.
	Add(method, path string, handler HandlerFunc, middleware ...MiddlewareFunc) *Route
}

type Engine struct {
	*echo.Echo

	Logger log.Logger
}

func (e *Engine) convertHandler(h HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if actx, ok := ctx.(*Context); ok {
			return h(actx)
		}

		return h(ctx.Get(MyContextKey).(*Context))
	}
}

func (e *Engine) convertFromHandler(h echo.HandlerFunc) HandlerFunc {
	return func(ctx *Context) error {
		return h(ctx)
	}
}

func (e *Engine) convertFromPreHandler(h echo.HandlerFunc) HandlerFunc {
	return func(ctx *Context) error {
		return h(ctx.Context)
	}
}

func (e *Engine) convertMiddleware(middleware MiddlewareFunc) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return e.convertHandler(middleware(e.convertFromHandler(h)))
	}
}

func (e *Engine) convertPreMiddleware(middleware MiddlewareFunc) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return e.convertHandler(middleware(e.convertFromPreHandler(h)))
	}
}

func (e *Engine) convertMiddlewares(middlewares []MiddlewareFunc) []echo.MiddlewareFunc {
	funcs := make([]echo.MiddlewareFunc, len(middlewares))
	for idx := range middlewares {
		funcs[idx] = e.convertMiddleware(middlewares[idx])
	}
	return funcs
}

func (e *Engine) convertPreMiddlewares(middlewares []MiddlewareFunc) []echo.MiddlewareFunc {
	funcs := make([]echo.MiddlewareFunc, len(middlewares))
	for idx := range middlewares {
		funcs[idx] = e.convertPreMiddleware(middlewares[idx])
	}
	return funcs
}

// Pre adds middleware to the chain which is run before router.
func (e *Engine) Pre(middlewares ...MiddlewareFunc) {
	e.Echo.Pre(e.convertPreMiddlewares(middlewares)...)
}

// Use adds middleware to the chain which is run after router.
func (e *Engine) Use(middlewares ...MiddlewareFunc) {
	e.Echo.Use(e.convertMiddlewares(middlewares)...)
}

// CONNECT registers a new CONNECT route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.CONNECT(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (e *Engine) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.DELETE(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (e *Engine) GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.GET(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.HEAD(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.OPTIONS(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// PATCH registers a new PATCH route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.PATCH(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) POST(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.POST(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.PUT(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// TRACE registers a new TRACE route for a path with matching handler in the
// router with optional route-level middleware.
func (e *Engine) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.TRACE(path, e.convertHandler(h), e.convertMiddlewares(m)...)
}

// Any registers a new route for all HTTP methods and path with matching handler
// in the router with optional route-level middleware.
func (e *Engine) Any(path string, handler HandlerFunc, m ...MiddlewareFunc) []*Route {
	return e.Echo.Any(path, e.convertHandler(handler), e.convertMiddlewares(m)...)
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (e *Engine) Add(method, path string, handler HandlerFunc, m ...MiddlewareFunc) *Route {
	return e.Echo.Add(method, path, e.convertHandler(handler), e.convertMiddlewares(m)...)
}

// Match implements `Echo#Match()` for sub-routes within the Group.
func (e *Engine) Match(methods []string, path string, handler HandlerFunc, m ...MiddlewareFunc) []*Route {
	return e.Echo.Match(methods, path, e.convertHandler(handler), e.convertMiddlewares(m)...)
}

// File registers a new route with path to serve a static file with optional route-level middleware.
func (e *Engine) File(path, file string) {
	e.Echo.File(path, file)
}

// Static implements `Echo#Static()` for sub-routes within the Group.
func (e *Engine) Static(prefix, root string) {
	e.Echo.Static(prefix, root)
}

func (e *Engine) Group(prefix string, m ...MiddlewareFunc) Party {
	g := e.Echo.Group(prefix, e.convertMiddlewares(m)...)
	return &Group{e, g}
}

type Group struct {
	engine *Engine
	group  *echo.Group
}

// Use adds middleware to the chain which is run after router.
func (e *Group) Use(middlewares ...MiddlewareFunc) {
	e.group.Use(e.engine.convertMiddlewares(middlewares)...)
}

// CONNECT registers a new CONNECT route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.CONNECT(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (g *Group) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.DELETE(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (g *Group) GET(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.GET(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.HEAD(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.OPTIONS(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// PATCH registers a new PATCH route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.PATCH(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) POST(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.POST(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.PUT(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// TRACE registers a new TRACE route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.TRACE(path, g.engine.convertHandler(h), g.engine.convertMiddlewares(m)...)
}

// Any registers a new route for all HTTP methods and path with matching handler
// in the router with optional route-level middleware.
func (g *Group) Any(path string, handler HandlerFunc, m ...MiddlewareFunc) []*Route {
	return g.group.Any(path, g.engine.convertHandler(handler), g.engine.convertMiddlewares(m)...)
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (g *Group) Add(method, path string, handler HandlerFunc, m ...MiddlewareFunc) *Route {
	return g.group.Add(method, path, g.engine.convertHandler(handler), g.engine.convertMiddlewares(m)...)
}

// Match implements `Echo#Match()` for sub-routes within the Group.
func (g *Group) Match(methods []string, path string, handler HandlerFunc, m ...MiddlewareFunc) []*Route {
	return g.group.Match(methods, path, g.engine.convertHandler(handler), g.engine.convertMiddlewares(m)...)
}

// File registers a new route with path to serve a static file with optional route-level middleware.
func (g *Group) File(path, file string) {
	g.group.File(path, file)
}

// Static implements `Echo#Static()` for sub-routes within the Group.
func (g *Group) Static(prefix, root string) {
	g.group.Static(prefix, root)
}

func (g *Group) Group(prefix string, m ...MiddlewareFunc) Party {
	sg := g.group.Group(prefix, g.engine.convertMiddlewares(m)...)
	return &Group{g.engine, sg}
}

func New() *Engine {
	e := &Engine{
		Echo: echo.New(),
	}
	e.Echo.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req := ctx.Request()
			actx := &Context{
				Context:    ctx,
				StdContext: req.Context(),
			}
			if e.Logger != nil {
				actx.CtxLogger = e.Logger.With(log.String("http.method", req.Method), log.Stringer("http.url", req.URL))
				actx.StdContext = log.ContextWithLogger(actx.StdContext, actx.CtxLogger)
			}
			ctx.Set(MyContextKey, actx)
			return next(ctx)
		}
	})
	e.Pre(Tracing("w"))
	e.Echo.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: func(c echo.Context) string {
			m := c.FormValue("_method")
			if m != "" {
				return m
			}
			return c.QueryParam("_method")
		}}))

	// Middleware
	e.Echo.Use(middleware.Logger())
	e.Echo.Use(middleware.Recover())
	e.Echo.HTTPErrorHandler = echo.HTTPErrorHandler(func(err error, c echo.Context) {
		if e.Logger != nil {
			e.Logger.Warn("处理请求发生错误", log.Error(err), log.String("url", c.Request().RequestURI))
		}
		e.Echo.DefaultHTTPErrorHandler(err, c)
	})
	return e
}
