package main

import (
	"rate-limiter/cycled-array"
	"sync"
)

type seriesPool struct {
	buffers sync.Pool
}

type TimeSeries struct {
	data cycled_array.CycledArray[int64]
	mu   sync.Mutex
}

func initTimeSeriesPool(maxCap int) *seriesPool {
	return &seriesPool{
		buffers: sync.Pool{New: func() interface{} {
			return TimeSeries{data: cycled_array.CycledArray[int64]{arr: make([]int64, maxCap)}}
		}},
	}
}

func (sp *seriesPool) Get() *TimeSeries {
	return sp.buffers.Get().(*TimeSeries)
}

const maxBufferCap = 100 << 10 // 100Ki

func (sp *seriesPool) Put(t *TimeSeries) {
	if cap(t.data.arr) > maxBufferCap {
		return
	}

	t.data.nextPos = 0
	t.data.len = 0

	sp.buffers.Put(t)
}
