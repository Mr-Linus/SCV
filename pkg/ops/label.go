package ops

import (
	"github.com/NJUPT-ISL/SCV/pkg/log"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func UpdateLabel(label map[string]string, name string) {
	client, err := kubernetes.NewForConfig(Config)
	if err != nil {
		panic(err)
	}
	node, err := client.CoreV1().Nodes().Get(name, v1.GetOptions{})
	if err != nil {
		log.ErrPrint(err)
	} else {
		if !CheckScvLabel(name) {
			nodeLabel := node.GetLabels()
			for i, v := range label {
				nodeLabel[i] = v
			}
			node.SetLabels(nodeLabel)
			if _, err := client.CoreV1().Nodes().Update(node); err != nil {
				log.ErrPrint(err)
			}
		}
	}
}

func CleanScvLabel(name string) {
	client, err := kubernetes.NewForConfig(Config)
	if err != nil {
		panic(err)
	}
	node, err := client.CoreV1().Nodes().Get(name, v1.GetOptions{})
	if err != nil {
		log.ErrPrint(err)
	} else {
		nodeLabel := node.GetLabels()
		for i := range nodeLabel {
			if strings.Contains(i, "scv/") {
				delete(nodeLabel, i)
			}
		}
		node.SetLabels(nodeLabel)
		if _, err := client.CoreV1().Nodes().Update(node); err != nil {
			log.ErrPrint(err)
		}
	}
}

func UpdateScvLabel(name string) {
	UpdateLabel(SetScvToMap(), name)
}

func CheckScvLabel(name string) bool {
	client, err := kubernetes.NewForConfig(Config)
	if err != nil {
		panic(err)
	}
	node, err := client.CoreV1().Nodes().Get(name, v1.GetOptions{})
	if err != nil {
		log.ErrPrint(err)
	} else {
		nodeLabel := node.GetLabels()
		for k, v := range SetScvToMap() {
			if value, ok := nodeLabel[k]; ok {
				if value != v {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
