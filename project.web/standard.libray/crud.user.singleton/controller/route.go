// Go Api server
// @jeffotoni

package controller

import (
	"net/http"

	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/api.user/controller/handler"
	cf "github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/pkg/config"
	mw "github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/pkg/middleware"
	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/pkg/pkg/cors"
)

func StartServer(cfg cf.Config) *GoServerHttp {

	///////////////////////////////////////
	/////
	// Create our middleware.
	// mdlw := prommiddleware.NewDefault()

	// DefaultServeMux
	mux := http.NewServeMux()

	corsx := cors.Domain()

	// POST handler /ping
	handlerApiPing := http.HandlerFunc(handler.Ping)
	mux.Handle(Endpoint().Ping, mw.Use(handlerApiPing,
		mw.CustomHeaders(),
		mw.Gzip(),
		mw.MaxClients(cf.MaxClients),
		mw.Logger("auth/ping"),
	))

	///////////////////////////////////////////////////////////
	// Endpoints
	//////////////////////////////////////
	// user
	// user/{uuid}
	/////////////////////////////////////
	mux.Handle("/", mw.Use(http.HandlerFunc(handler.HomeHandler),
		//mw.MetricsPrometheusDinamic(),
		mw.Cors(),
		mw.CustomHeaders(),
		mw.AutJwt(),
		//mw.AuthJwtNot([]string{"/products"}),
		mw.Logger("/")))

	// withMetrics := mw.MetricsPrometheus(handler)
	// middpromet := mdlw.Handler("", mux)

	handlerCors := corsx.Handler(mux)

	// println(cfg.Host)
	// Create the HTML Server
	ApiServer := GoServerHttp{
		server: &http.Server{
			Addr:    cfg.Host,
			Handler: handlerCors,
			//Handler: middpromet,
			//ReadTimeout:  time.Millisecond * 600,
			//WriteTimeout: time.Millisecond * 400,
			//ReadTimeout:  cfg.ReadTimeout,
			//WriteTimeout: cfg.WriteTimeout,
			// IdleTimeout:    1000 * time.Millisecond,
			//MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	ApiServer.wg.Add(1)

	// Start the listener
	//go func() {
	Show(cfg)
	ApiServer.server.ListenAndServe()
	ApiServer.wg.Done()
	//}()

	// Serve our metrics.
	// Prometheus
	// Metrics
	// go func() {
	// 	if err := http.ListenAndServe(":"+cfp.PORT_METRICS, promhttp.Handler()); err != nil {
	// 		log.Printf("Eror while serving metrics: %s", err)
	// 	}
	// }()

	return &ApiServer
}