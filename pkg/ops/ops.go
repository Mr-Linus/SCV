package ops

import (
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/log"

)

func InitSCV(mode string){
	switch mode {
		case  "Full":
			collection.InitSCVWithFullMode()
		case  "High":
			collection.InitSCVWithHighMode()
		case  "LowPower":
			collection.InitSCVWithLowPowerMode()
	}
}

func UpdateSCV(mode string){
	switch mode {
		case  "Full":
			collection.UpdateSCVWithFullMode()
		case  "High":
			collection.UpdateSCVWithHighMode()
		case  "LowPower":
			collection.UpdateSCVWithLowPowerMode()
	}
}

func PrintSCV(){
	log.LogPrint(collection.Scv.Model)
}