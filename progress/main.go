package main

import (
	"math/rand"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func main() {
	bar := progressbar.Default(100)

	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(50 * time.Millisecond)
	}

	// 初始化全局进度管理器
	p := mpb.New(mpb.WithWidth(60)) // 进度条总宽度

	// 第一个任务
	total1 := 100
	bar1 := p.New(int64(total1),
		mpb.BarStyle().Lbound("[").Filler("=").Tip(">").Padding("-").Rbound("]"),
		mpb.PrependDecorators(
			decor.Name("任务A:", decor.WC{W: 8}),
			decor.CountersNoUnit("%d / %d"),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)

	// 第二个任务
	total2 := 200
	bar2 := p.New(int64(total2),
		mpb.BarStyle().Lbound("[").Filler("█").Tip(">").Padding("░").Rbound("]"),
		mpb.PrependDecorators(
			decor.Name("任务B:", decor.WC{W: 8}),
			decor.CountersNoUnit("%d / %d"),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WCSyncSpace)),
	)

	// 模拟第一个任务
	go func() {
		for i := 0; i < total1; i++ {
			bar1.Increment()
			time.Sleep(time.Duration(rand.Intn(50)+20) * time.Millisecond)
		}
	}()

	// 模拟第二个任务
	go func() {
		for i := 0; i < total2; i++ {
			bar2.Increment()
			time.Sleep(time.Duration(rand.Intn(30)+10) * time.Millisecond)
		}
	}()

	// 等待所有任务结束
	p.Wait()
}
