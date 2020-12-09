package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	s1 := http.Server{
		Addr: ":9000",
	}

	s2 := http.Server{
		Addr: ":9001",
	}

	g.Go(func() error {
		if err := s1.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := s2.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	c := make(chan os.Signal, 1)

	shut := make(chan struct{})

	// Passing no signals to Notify means that
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case <-c:
				cancel()
				return
			default:

			}
		}
	}()

	go func() {
		for {
			select {
			// 接收cancel 信号
			case <-ctx.Done():
				_ = s1.Shutdown(context.Background())
				_ = s2.Shutdown(context.Background())
				shut <- struct{}{}
				return
			default:

			}
		}
	}()

	if err := g.Wait(); err != nil {
		log.Println("errgroup.Err", err)
	}
	// 接收 shutdown 信号
	<-shut

}
