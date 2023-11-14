package service

import (
	"github.com/aiviaio/go-binance/v2"
	"github.com/gin-gonic/gin"
	"sync"
	"testByAivia/utils"
)

func (s *Service) GetPrice(ctx *gin.Context) {
	client := binance.NewClient("", "")

	symbols, err := utils.GetSymbols(ctx, client)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error fetching symbols"})
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(symbols))

	ch := make(chan map[string]string, len(symbols))

	for _, symbol := range symbols {
		go utils.GetPrice(ctx, client, symbol, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		ctx.JSON(200, data)
	}
}
