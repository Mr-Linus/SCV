package collection

import (
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

func AddGPU() error {
	var health string = "Healthy"
	err := nvml.Init()
	if err != nil {
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil {
			log.ErrPrint(err)
		}
	}()
	count, err := nvml.GetDeviceCount()
	if err != nil {
		log.ErrPrint(err)
	}
	for i := uint(0); i < count; i++ {

		device, err := nvml.NewDevice(i)
		if err != nil {
			log.ErrPrint(err)
		}
		status, err := device.Status()
		if err != nil {
			log.ErrPrint(err)
			health = "Unhealthy"
		}
		GPUs = append(GPUs, GPU{
			ID:          i,
			Health:      health,
			Model:       *device.Model,
			Power:       *device.Power,
			Memory:      *device.Memory,
			MemoryClock: *device.Clocks.Memory,
			FreeMemory:  *status.Memory.Global.Free,
			Cores:       *device.Clocks.Cores,
			Bandwidth:   *device.PCI.Bandwidth,
			Device:      *device,
		})
	}
	return err
}

func UpdateGPU() error {
	var NewGPUs []GPU
	var health string = "Healthy"
	err := nvml.Init()
	if err != nil {
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil {
			log.ErrPrint(err)
		}
	}()
	count, err := nvml.GetDeviceCount()
	if err != nil {
		log.ErrPrint(err)
	}
	for i := uint(0); i < count; i++ {
		device, err := nvml.NewDevice(i)
		if err != nil {
			log.ErrPrint(err)
		}
		status, err := device.Status()
		if err != nil {
			log.ErrPrint(err)
			health = "Unhealthy"
		}
		NewGPUs = append(NewGPUs, GPU{
			ID:          i,
			Health:      health,
			Model:       *device.Model,
			Power:       *device.Power,
			Memory:      *device.Memory,
			MemoryClock: *device.Clocks.Memory,
			FreeMemory:  *status.Memory.Global.Free,
			Cores:       *device.Clocks.Cores,
			Bandwidth:   *device.PCI.Bandwidth,
			Device:      *device,
		})
	}
	GPUs = NewGPUs
	return err
}
