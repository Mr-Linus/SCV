package collection

import (
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
)

func AddGPU() error {
	var health string = "Healthy"
	err := nvml.Init()
	if  err != nil{
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil{
			log.ErrPrint(err)
		}
	}()
	count, err := nvml.GetDeviceCount()
	if err != nil {
		log.ErrPrint(err)
	}
	for i := uint(0); i < count; i++{

		device, err := nvml.NewDevice(i)
		if err != nil{
			log.ErrPrint(err)
		}
		status,err := device.Status()
		if err != nil{
			log.ErrPrint(err)
			health = "Unhealthy"
		}
		GPUs = append(GPUs,GPU{
			ID: i,
			health:    	health,
			Model:      *device.Model,
			Power:      *device.Power,
			Memory:     *device.Memory,
			MemoryClock: *device.Clocks.Memory,
			FreeMemory: *status.Memory.Global.Free,
			Cores:      *device.Clocks.Cores,
			Bandwidth:  *device.PCI.Bandwidth,
			Device: *device,
		})
	}
	return err
}



func UpdateGPU() error {
	var NewGPUs []GPU
	err := nvml.Init()
	if err != nil{
		log.ErrPrint(err)
	}
	defer func() {
		if err := nvml.Shutdown(); err != nil{
			log.ErrPrint(err)
		}
	}()
	for _, g := range GPUs{
		var health string = "Healthy"
		status,err := g.Device.Status()
		if err != nil{
			log.ErrPrint(err)
			health = "Unhealthy"
		}
		g.health = health
		g.Memory = *status.Memory.Global.Free
		GPUs = append(NewGPUs,g)
	}
	GPUs = NewGPUs
	return err
}

