package service

import (
	"bufio"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HelperService struct {
}

func (s *HelperService) ChunkData(ctx *fiber.Ctx, data [][]byte) error {

	ctx.Set("Transfer-Encoding", "chunked")
	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {

		for i := 0; i < len(data); i++ {
			write, err := w.Write(data[i])
			if err != nil {
				return
			}
			println(write)
			err = w.Flush()
			if err != nil {
				print(err)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	})
	return nil
}
