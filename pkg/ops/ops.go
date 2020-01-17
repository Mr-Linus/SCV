package ops

import (
	"encoding/json"
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/log"
)

func InitSCV(mode string){
	log.Print("Init SCV, mode:"+mode)
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
	log.Print("Init SCV, mode:"+mode)
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
	s,err := json.Marshal(&collection.Scv)
	if err != nil{
		log.ErrPrint(err)
	}
	log.Print("SCV Info: "+string(s))
}