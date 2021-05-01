package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。


启动多个server，能够监听不同端口，确保当某一个出错的时候，其它的也能退出。在各个server退出之前，要执行一些清理；所有server退出之后，主程序退出之前，也可能需要执行一些东西
进一步要考虑，所谓退出前的执行动作，要不要考虑可扩展，就是允许你的用户注册自己的退出之前回调。更加进一步的考虑是，如果允许扩展，那么怎么定义执行顺序。比如说，数据库连接必须要等所有的请求处理完成才能释放，这种依赖问题
如果万一退出执行时间很长，用户不会希望自己一直卡在那里，所以要有考虑强制退出的机制
https://mp.weixin.qq.com/s/B9F1Ta0QWJZNpkzq8URGRQ
https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869
https://lailin.xyz/post/go-training-week3-errgroup.html **
*/
var (
	httpServer01 *http.Server
	httpServer02 *http.Server
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx) //有error会先cancle

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)

	g.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
			cancel()
			w.Write([]byte("shutdown01"))
		})
		httpServer01 = &http.Server{
			Handler:      mux,
			Addr:         fmt.Sprintf(":9001"),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		if err := httpServer01.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/shutdown/trigger", func(w http.ResponseWriter, r *http.Request) {
			cancel()
			w.Write([]byte("shutdown02"))
		})
		httpServer02 = &http.Server{
			Handler:      mux,
			Addr:         fmt.Sprintf(":9002"),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		if err := httpServer02.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		select {
		case <-quit:
			fmt.Println("system interupt, exit...")
		case <-ctx.Done():
			fmt.Println("errgroup exit...")
		}

		ctxShut, cancelShut := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancelShut()
		if err := httpServer01.Shutdown(ctxShut); err != nil {
			fmt.Println("serve01 quit, err:", err)
		}
		if err := httpServer02.Shutdown(ctxShut); err != nil {
			fmt.Println("serve02 quit, err:", err)
		}

		return errors.New("server quit")
	})

	if err := g.Wait(); err != nil {
		fmt.Println("serve shutdown, err:", err)
	}
}
