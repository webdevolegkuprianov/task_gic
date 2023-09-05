package stdservice

import (
	"context"
	"fmt"
)

type stdService struct {
	ctx context.Context
	ch  chan int
}

func RegisterStdService(ctx context.Context, ch chan int) {

	serv := &stdService{
		ctx: ctx,
		ch:  ch,
	}

	go serv.runStdout()

}

func (s *stdService) runStdout() {

	for {
		select {
		case v, ok := <-s.ch:
			if !ok {
				return
			}
			fmt.Printf("кол-во запросов: %d\n", v)
		case <-s.ctx.Done():
			return
		}
	}

}
