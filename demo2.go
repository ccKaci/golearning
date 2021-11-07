package 

// start HTTP server
func StartHttpServer(srv *http.Server) error {
        http.HandelFunc("/demo2", DemoServer)
        fmt.Println("http server start")
        err := srv.ListenAndServe()
        return err
}

// add a HTTP hanlder
func DemoServer(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "it is demo2\n")
}

//
func main() {
        ctx := context.Background()
        // withCancle -> cancle()     to cancle context
        ctx, cancel := context.WithCancel(ctx)
        // use errgroup to do the goroutine cancel
        group, errCtx := errgroup.WithContext(ctx)
        // http server
        srv := &http.Server{Addr: ":9090"}

        group.Go(func() error {
				return StartHttpServer(srv)
        })

        group.Go(func() error {
                <- errCtx.Done()            // might be closed
                fmt.Println("http server stop")
                return srv.Shutdown(errCtx)
        })

        chanel := make(chan os.Signal, 1)
        signal.Notify(chanel)

        group.Go(func() error {
                for {
                        select {
                                case <- errCtx.Done():
                                        return errCtx.Err()
                                case <- chanel:
                                        cancel()
                        }
                }
                return nil
        })
        if err := group.Wait(); err != nil {
                fmt.Println("group error: ", err)
        }
        fmt.Println("all group done!")

}
