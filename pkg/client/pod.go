package client

import (
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/config"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/types"
	"github.com/jinghzhu/kutils/pod"
)

func init() {
	initDefaultPodClient()
}

func initDefaultPodClient() {
	c, err := pod.New(types.GetDefaultCtx(), "", config.GetConfig().GetKubeconfigPath())
	if err != nil {
		panic(err)
	}
	defaultPodClient = c
}

func GetDefaultPodClient() *pod.Client {
	return defaultPodClient
}
