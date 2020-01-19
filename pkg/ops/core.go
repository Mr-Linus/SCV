package ops

import (
	"errors"
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"reflect"
	"strconv"
	"strings"
)

var Config *rest.Config

func InitInClusterConfig(){
	err := errors.New("")
	// InClusterConfig
	log.Print("Init kubernetes Config. ")
	Config, err = rest.InClusterConfig()
	if err != nil {
		log.ErrPrint(err)
	}
}

func InitOutOfClusterConfig(){
	err := errors.New("")
	log.Print("Init kubernetes Config. ")
	Config, err = clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		log.ErrPrint(err)
	}
}

func SetScvToMap() map[string]string{
	m := make(map[string]string)
	elem := reflect.ValueOf(collection.Scv)
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if elem.Field(i).Type() == reflect.TypeOf(""){
			// TODO：add prefix e.g: isl.gpu/
			m["scv/"+relType.Field(i).Name] = strings.Replace(elem.Field(i).String()," ","-",-1)
		}else if elem.Field(i).Type() == reflect.TypeOf(true){
			m["scv/"+relType.Field(i).Name] = BoolToString(elem.Field(i).Bool())
		}else {
			m["scv/"+relType.Field(i).Name] = strconv.Itoa(int(elem.Field(i).Uint()))
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