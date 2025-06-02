package main

import (
	"net/http"
	"log"
)
//使用 mux.Handlefunc 来注册处理器，处理器可以只是一个与http.HandleFunc签名匹配的函数
//
func main(){
	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()
	mux.Handle("/app", http.StripPrefix("/app",http.FileServer(http.Dir(filepathRoot + "/app"))))
	mux.Handle("/app/assets/", http.StripPrefix("/app/assets", http.FileServer(http.Dir(filepathRoot + "/app/assets"))))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
									w.Header().Set("Content-Type", "text/plain; charset=utf-8")
									w.WriteHeader(http.StatusOK)
									w.Write([]byte("OK"))

	})	
	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,//MUX 会根据不同的根路径来处理
	}//这里是配置一个http服务器的结构体，设置监听地址和处理器
	
	err := server.ListenAndServe()
	//启动服务器，监听指定端口，处理请求
	if err != nil{
		panic(err)
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}