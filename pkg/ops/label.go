package ops

import (
	"github.com/NJUPT-ISL/SCV/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)


func UpdateLabel(label map[string]string, name string){
	clientset, err := kubernetes.NewForConfig(Config)
	if err != nil {
		panic(err)
	}
	node, err := clientset.CoreV1().Nodes().Get(name,metav1.GetOptions{})
	if err != nil{
		log.ErrPrint(err)
	}
	nodeLabel := node.GetLabels()
	for i,v := range label{
		nodeLabel[i]=v
	}
	node.SetLabels(nodeLabel)
	if _,err :=  clientset.CoreV1().Nodes().Update(node);err != nil{
		log.ErrPrint(err)
	}
}

func UpdateScvLabel(name string){
	UpdateLabel(SetScvToMap(),name)
}