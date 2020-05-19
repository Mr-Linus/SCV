package collection

import (
	"errors"
	"github.com/NJUPT-ISL/SCV/pkg/log"
)

func CalculateBestGPUWithHighMode() GPU {
	var (
		Scores []Score
		Max    uint64 = 0
		Device GPU
	)
	for _, g := range GPUs {
		if g.Health == "Unhealthy" {
			Scores = append(Scores, Score{
				Device: g,
				Score:  0,
			})
			continue
		} else {
			Scores = append(Scores, Score{
				Device: g,
				Score:  g.CalculateScoreWithHighMode(),
			})
		}
	}
	for _, s := range Scores {
		if s.Score > Max {
			Max = s.Score
			Device = s.Device
		}
	}
	return Device
}

func CalculateBestGPUWithFullMode() GPU {
	var (
		Scores []Score
		Max    uint64 = 0
		Device GPU
	)
	for _, g := range GPUs {
		if g.Health == "Unhealthy" {
			Scores = append(Scores, Score{
				Device: g,
				Score:  0,
			})
			continue
		} else {
			Scores = append(Scores, Score{
				Device: g,
				Score:  g.CalculateScoreWithFullMode(),
			})
		}
	}
	for _, s := range Scores {
		if s.Score > Max {
			Max = s.Score
			Device = s.Device
		}
	}
	return Device
}

func CalculateBestGPUWithLowPowerMode() GPU {
	var (
		Scores []Score
		Max    uint64 = 0
		Device GPU
	)
	for _, g := range GPUs {
		if g.Health == "Unhealthy" {
			Scores = append(Scores, Score{
				Device: g,
				Score:  0,
			})
			continue
		} else {
			Scores = append(Scores, Score{
				Device: g,
				Score:  g.CalculateScoreWithLowPowerMode(),
			})
		}
	}
	for _, s := range Scores {
		if s.Score > Max {
			Max = s.Score
			Device = s.Device
		}
	}
	return Device
}

func (g *GPU) CalculateScoreWithWeight(FreeW uint64, MemW uint64, MemCW uint64, CoreW uint64, BandC uint64, PoW uint64) uint64 {
	maxFree, err := CalculateMemoryMaxGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	maxMemClock, err := CalculateMemoryClockMaxGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	mem, err := CalculateMemoryMaxGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	maxBand, err := CalculateMaxBandwidthGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	maxCore, err := CalculateCoreMaxGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	maxP, err := CalculateMaxPowerGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	return g.FreeMemory/maxFree*FreeW + uint64(g.MemoryClock/maxMemClock)*MemCW + g.Memory/mem*MemW + uint64(g.Bandwidth/maxBand)*BandC + uint64(g.Cores/maxCore)*CoreW + uint64(1-g.Power/maxP)*PoW
}

func (g *GPU) CalculateScoreWithHighMode() uint64 {
	return g.CalculateScoreWithWeight(8, 3, 4, 10, 5, 0)
}

func (g *GPU) CalculateScoreWithFullMode() uint64 {
	return g.CalculateScoreWithWeight(10, 8, 5, 4, 4, 0)
}

func (g *GPU) CalculateScoreWithLowPowerMode() uint64 {
	return g.CalculateScoreWithWeight(5, 2, 1, 1, 3, 10)
}

func CalculateSCVLevel() string {
	MaxMem, err := CalculateMemoryMaxGPU()
	if err != nil {
		log.ErrPrint(err)
	}
	if MaxMem > 10000 {
		return "High"
	}
	if MaxMem >= 4000 {
		return "Medium"
	}
	return "Low"
}
func CalculateSCVMemorySum() uint64 {
	sum, err := CalculateGPUMemorySum()
	if err != nil {
		log.ErrPrint(err)
	}
	return sum
}

func CalculateGPUMemorySum() (uint64, error) {
	var Sum uint64 = 0
	for _, g := range GPUs {
		Sum += g.FreeMemory
	}
	if Sum == 0 {
		return 0, errors.New("The Sum of the GPU memory is 0. ")
	}
	return Sum, nil
}

func CalculateCoreMaxGPU() (uint, error) {
	var maxCore uint = 0
	for _, g := range GPUs {
		if g.Cores > maxCore {
			maxCore = g.Cores
		}
	}
	if maxCore == 0 {
		return 0, errors.New("The Max Core of the GPUs is 0. ")
	}
	return maxCore, nil
}

func CalculateMemoryFreeMaxGPU() (uint64, error) {
	var maxMem uint64 = 0
	for _, g := range GPUs {
		if g.FreeMemory > maxMem {
			maxMem = g.FreeMemory
		}
	}
	if maxMem == 0 {
		return 0, errors.New("The Max Free Memory of the GPUs is 0. ")
	}
	return maxMem, nil
}

func CalculateMemoryClockMaxGPU() (uint, error) {
	var maxMemC uint = 0
	for _, g := range GPUs {
		if g.MemoryClock > maxMemC {
			maxMemC = g.MemoryClock
		}
	}
	if maxMemC == 0 {
		return 0, errors.New("The Max Memory Clock of the GPUs is 0. ")
	}
	return maxMemC, nil
}

func CalculateMemoryMaxGPU() (uint64, error) {
	var maxMem uint64 = 0
	for _, g := range GPUs {
		if g.Memory > maxMem {
			maxMem = g.Memory
		}
	}
	if maxMem == 0 {
		return 0, errors.New("The Max Memory of the GPUs is 0. ")
	}
	return maxMem, nil
}

func CalculateMaxBandwidthGPU() (uint, error) {
	var maxBand uint = 0
	for _, g := range GPUs {
		if g.Bandwidth > maxBand {
			maxBand = g.Bandwidth
		}
	}
	if maxBand == 0 {
		return 0, errors.New("The Max Memory of the GPUs is 0. ")
	}
	return maxBand, nil
}

func CalculateMaxPowerGPU() (uint, error) {
	var maxPower uint = 0
	for _, g := range GPUs {
		if g.Power > maxPower {
			maxPower = g.Power
		}
	}
	if maxPower == 0 {
		return 0, errors.New("The Min Power of the GPUs is 0. ")
	}
	return maxPower, nil
}
