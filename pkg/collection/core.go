package collection

import (
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

type SCV struct {
	Gpu bool	// If the node has GPU, the value is true.
	Health string	// If all GPU is unhealthy,the value is unhealthy.
	Model string	// If set High mode,it is the best GPU of the GPUs.
	Level string
	Power uint		// If set High mode,it is the max power of the GPUs.
	Memory uint64	// The Max of the free memory the single GPU.
	MemorySum uint64 // The Sum of the free memory of the GPUs.
	MemoryClock uint
	FreeMemory uint64
	Cores uint
	Bandwidth uint
}

type GPU struct {
	ID uint
	Health string
	Model string
	Power uint
	Memory uint64
	MemoryClock uint
	FreeMemory uint64
	Cores uint
	Bandwidth uint
	Device nvml.Device
}

type Score struct {
	Device GPU
	Score uint64
}

var (
	Scv SCV
	GPUs []GPU
	Mode string
	Node string
)

func CheckHealth() string {
	for _, g := range GPUs{
		if g.Health == "Healthy"{
			return "Healthy"
		}
	}
	return "Unhealthy"
}

func CheckGPU() bool {
	err := nvml.Init()
	if err != nil{
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil{
			log.ErrPrint(err)
		}
	}()
	count, err := nvml.GetDeviceCount()
	if err != nil{
		log.ErrPrint(err)
	}
	if count > 0 {
		return true
	}
	return false
}

func InitSCVWithHighMode(){
	InitModeSCV(CalculateBestGPUWithHighMode)
}

func InitSCVWithLowPowerMode(){
	InitModeSCV(CalculateBestGPUWithLowPowerMode)
}

func InitSCVWithFullMode(){
	InitModeSCV(CalculateBestGPUWithFullMode)
}

func UpdateSCVWithHighMode(){
	UpdateModeSCV(CalculateBestGPUWithHighMode)
}

func UpdateSCVWithFullMode(){
	UpdateModeSCV(CalculateBestGPUWithFullMode)
}

func UpdateSCVWithLowPowerMode(){
	UpdateModeSCV(CalculateBestGPUWithLowPowerMode)
}

func InitModeSCV(Mode func() GPU){
	if err := AddGPU(); err != nil{
		log.ErrPrint(err)
	}
	Device := Mode()
	Scv = SCV{
		Gpu:         CheckGPU(),
		Health:      CheckHealth(),
		Model:       Device.Model,
		Level:       CalculateSCVLevel(),
		Power:       Device.Power,
		Memory:      Device.Memory,
		MemorySum:   CalculateSCVMemorySum(),
		MemoryClock: Device.MemoryClock,
		FreeMemory:  Device.FreeMemory,
		Cores:       Device.Cores,
		Bandwidth:   Device.Bandwidth,
	}
}

func UpdateModeSCV(Mode func() GPU){
	if err := UpdateGPU(); err != nil {
		log.ErrPrint(err)
	}
	Device := Mode()
	Scv = SCV{
		Gpu:         CheckGPU(),
		Health:      CheckHealth(),
		Model:       Device.Model,
		Level:       CalculateSCVLevel(),
		Power:       Device.Power,
		Memory:      Device.Memory,
		MemorySum:   CalculateSCVMemorySum(),
		MemoryClock: Device.MemoryClock,
		FreeMemory:  Device.FreeMemory,
		Cores:       Device.Cores,
		Bandwidth:   Device.Bandwidth,
	}
}