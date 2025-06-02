package main

import (
	"log"
	"net/http"
	"sync/atomic"
	
)

//使用 mux.Handlefunc 来注册处理器，处理器可以只是一个与http.HandleFunc签名匹配的函数
//
type apiConfig struct {
	fileserverHits atomic.Int32
}





func main(){
	apiCfg := &apiConfig{
			fileserverHits: atomic.Int32{},
	}
	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()
	fileSeverHandler := http.StripPrefix("/app",http.FileServer(http.Dir(filepathRoot + "/app")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fileSeverHandler))
	mux.Handle("/app/assets/", http.StripPrefix("/app/assets", http.FileServer(http.Dir(filepathRoot + "/app/assets"))))
	mux.HandleFunc("GET /healthz", readinessHandler)
	mux.HandleFunc("GET /metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /reset", apiCfg.resetHandler)
	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,//MUX 会根据不同的根路径来处理
	}//这里是配置一个http服务器的结构体，设置监听地址和处理器
	

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}

