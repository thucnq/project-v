package api

import (
	"unsafe"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"project-v/internal/app-api/handler/health"
	"project-v/internal/middleware"
	"project-v/pkg/l"
)

// Server ...
type Server struct {
	r *fiber.App
}

// HTTPErrorResponse ...
type HTTPErrorResponse struct {
	Status     string                 `json:"status"`
	Code       uint32                 `json:"code"`
	Message    string                 `json:"message"`
	DevMessage interface{}            `json:"dev_message" swaggerignore:"true"`
	Errors     map[string]interface{} `json:"errors"`
	RID        string                 `json:"rid"`
}

func getString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//	@title			Workflow API
//	@version		2.0
//	@description	Workflow API .

//	@contact.name	API Support
//	@contact.url	http://localhost
//	@contact.email

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host						localhost:5000
// @BasePath					localhost:5000
// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						authToken
// @securitydefinitions.apikey	BearerAuth
// @in							header
// @name						authToken
// @securityDefinitions.apikey	WorkspaceIDKey
// @in							header
// @name						authToken
// @securityDefinitions.apikey	LangKey
// @in							header
// @name						authToken
func New(env string) *Server {
	r := fiber.New(
		fiber.Config{
			ErrorHandler:    customErrorHandler(env),
			WriteBufferSize: 15 * 4096,
			ReadBufferSize:  15 * 4096,
		},
	)

	return &Server{
		r,
	}
}

// Listen ...
func (o Server) Listen(add string) error {
	return o.r.Listen(add)
}

// Middleware ...
func (o Server) Middleware(ll l.Logger) {
	o.r.Use(cors.New())
	o.r.Use(pprof.New())
	o.r.Use(requestid.New())
	o.r.Use(compress.New())
	o.r.Use(middleware.NewLogging(ll))
}

// InitHealth ...
func (o Server) InitHealth(healthHandler health.Controller) {
	o.r.Get("/health", healthHandler.Health)
	o.r.Get("/liveness", healthHandler.Liveness)
}

// InitMetrics ...
func (o Server) InitMetrics() {
	prometheus := fiberprometheus.New("workflow")
	prometheus.RegisterAt(o.r, "/metrics")
	o.r.Use(prometheus.Middleware)
}

// InitLogHandler ...
func (o Server) InitLogHandler() {
	o.r.Get("/log/level", adaptor.HTTPHandlerFunc(l.ServeHTTP))
	o.r.Put("/log/level", adaptor.HTTPHandlerFunc(l.ServeHTTP))
}
