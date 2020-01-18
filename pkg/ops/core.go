package ops

import (
	"errors"
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"k8s.io/client-go/rest"
	"reflect"
	"strconv"
)

var Config *rest.Config

func InitInClusterConfig(){
	err := errors.New("")
	// InClusterConfig
	log.Print("Init kubernetes Config. ")
	Config, err = rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
}

func SetScvToMap() map[string]string{
	m := make(map[string]string)
	elem := reflect.ValueOf(collection.Scv)
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if elem.Field(i).Type() == reflect.TypeOf(""){
			m[relType.Field(i).Name] = elem.Field(i).String()
		}else if elem.Field(i).Type() == reflect.TypeOf(true){
			m[relType.Field(i).Name] = BoolToString(elem.Field(i).Bool())
		}else {
			m[relType.Field(i).Name] = strconv.Itoa(int(elem.Field(i).Uint()))
		}
	}
	return m
}

func BoolToString(value bool) string{
	if value{
		return "True"
	}else {
		return "False"
	}
}