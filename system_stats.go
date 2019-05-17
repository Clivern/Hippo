// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"runtime"
	"time"
)

// SystemStats stuct
type SystemStats struct {
	EnableCPU bool
	EnableMem bool
	EnableGC  bool
	StartTime time.Time
	Stats     map[string]uint64
}

// NewSystemStats creates a new SystemStats
func NewSystemStats(enableCPU, enableMem, enableGC bool) *SystemStats {
	return &SystemStats{
		EnableCPU: enableCPU,
		EnableMem: enableMem,
		EnableGC:  enableGC,
	}
}

// Collect collects enabled stats
func (s *SystemStats) Collect() {
	s.Stats = make(map[string]uint64)
	mStats := runtime.MemStats{}
	if s.EnableMem {
		s.outputMemStats(&mStats)
	}
	if s.EnableGC {
		s.outputGCStats(&mStats)
	}
	if s.EnableCPU {
		s.outputCPUStats()
	}
	s.outputTimeStats()
}

// outputCPUStats sets CPU stats
func (s *SystemStats) outputCPUStats() {
	s.append("cpu.goroutines", uint64(runtime.NumGoroutine()))
	s.append("cpu.cgo_calls", uint64(runtime.NumCgoCall()))
}

// outputMemStats sets memory stats
func (s *SystemStats) outputMemStats(m *runtime.MemStats) {
	// General
	s.append("mem.alloc", m.Alloc)
	s.append("mem.total", m.TotalAlloc)
	s.append("mem.sys", m.Sys)
	s.append("mem.lookups", m.Lookups)
	s.append("mem.malloc", m.Mallocs)
	s.append("mem.frees", m.Frees)

	// Heap
	s.append("mem.heap.alloc", m.HeapAlloc)
	s.append("mem.heap.sys", m.HeapSys)
	s.append("mem.heap.idle", m.HeapIdle)
	s.append("mem.heap.inuse", m.HeapInuse)
	s.append("mem.heap.released", m.HeapReleased)
	s.append("mem.heap.objects", m.HeapObjects)

	// Stack
	s.append("mem.stack.inuse", m.StackInuse)
	s.append("mem.stack.sys", m.StackSys)
	s.append("mem.stack.mspan_inuse", m.MSpanInuse)
	s.append("mem.stack.mspan_sys", m.MSpanSys)
	s.append("mem.stack.mcache_inuse", m.MCacheInuse)
	s.append("mem.stack.mcache_sys", m.MCacheSys)
	s.append("mem.othersys", m.OtherSys)
}

// outputGCStats sets GC stats
func (s *SystemStats) outputGCStats(m *runtime.MemStats) {
	s.append("mem.gc.sys", m.GCSys)
	s.append("mem.gc.next", m.NextGC)
	s.append("mem.gc.last", m.LastGC)
	s.append("mem.gc.pause_total", m.PauseTotalNs)
	s.append("mem.gc.pause", m.PauseNs[(m.NumGC+255)%256])
	s.append("mem.gc.count", uint64(m.NumGC))
}

// outputTimeStats sets uptime
func (s *SystemStats) outputTimeStats() {
	s.append("uptime", uint64(time.Since(s.StartTime).Seconds()))
}

// GetStats get stats list
func (s *SystemStats) GetStats() map[string]uint64 {
	return s.Stats
}

// append add stats
func (s *SystemStats) append(key string, value uint64) {
	s.Stats[key] = value
}
