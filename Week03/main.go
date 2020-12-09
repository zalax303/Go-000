package Week03

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	// http
	g.Go(func() error {
		s := &http.Server{Addr: ":8080", Handler: &handler{}}
		go func() {
			<-ctx.Done()
			_ = s.Shutdown(ctx)
		}()

		return s.ListenAndServe()
	})

	// signal
	g.Go(func() error {
		return signalMonitor(ctx)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("quit group err is %+v", err)
	}

}

type handler struct {
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world!"))
}

func signalMonitor(ctx context.Context) error {
	ch := make(chan os.Signal)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	select {
	case c := <-ch:
		return fmt.Errorf("quit is :%+v", c)
	case <-ctx.Done():
		return fmt.Errorf("quit group")

	}
}
