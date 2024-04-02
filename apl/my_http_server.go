package apl

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"techbookfest16-sample/domain/objects"

	"net/http"
	"os"
	"os/signal"
	"time"
)

// myHttpServer はGo標準のhttpサーバー機能を使うための構造体。
// EchoやGinなどのサードパーティー製WAF(Web Application Framework)を使う選択肢もアリ。
type myHttpServer struct {
	mux       *http.ServeMux
	httpPort  int
	defaultDb *sql.DB // 実際にはDBのスキーマ分け等、工夫が必要

	// 実際にはoriginチェックやmethodチェック、認証・認可について、各種ミドルウェアが必要
}

func NewMyHttpServer(defaultDb *sql.DB, httpPort int) *myHttpServer {
	return &myHttpServer{
		mux:       http.NewServeMux(),
		httpPort:  httpPort,
		defaultDb: defaultDb,
	}
}

func (mhs *myHttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// ミドルウェアを差し込む場合はここで細工すること
	// 細工次第では、GET専用やPOST専用などにできたりする
	mhs.mux.HandleFunc(pattern, handler)
}

func (mhs *myHttpServer) EntryUsecase(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// ミドルウェアを差し込む場合はここで細工すること
	// 細工次第では、GET専用やPOST専用などにできたりする
	mhs.mux.HandleFunc(pattern, handler)
}

func (mhs myHttpServer) GetNo(r *http.Request) (objects.No, error) {
	id, err := strconv.ParseInt(r.URL.Path, 10, 64)
	if err != nil {
		return 0, err
	}
	return objects.No(id), nil
}

func (mhs myHttpServer) GetCode(r *http.Request) string {
	return r.URL.Path
}

func (mhs *myHttpServer) EntryUsecaseWithNo(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// ミドルウェアを差し込む場合はここで細工すること
	// 細工次第では、GET専用やPOST専用などにできたりする
	mhs.mux.Handle(pattern, http.StripPrefix(pattern, http.HandlerFunc(handler)))
}

func (mhs *myHttpServer) ListenAndServe() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", mhs.httpPort),
		Handler:           mhs.mux, // 共通のミドルウェアを差し込む場合はここで細工すること
		ReadHeaderTimeout: 20 * time.Second,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err) // 実際にはDBにもログを書き込むと良い
	}
}
