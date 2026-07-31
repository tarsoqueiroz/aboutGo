package main

import (
	"fmt"

	"github.com/tarsoqueiroz/gerenciador-pods/pods"
	"github.com/tarsoqueiroz/gerenciador-pods/utils"
)

func main() {
	utils.Info("Iniciando gerenciador de pods")

	podsList := pods.ListPods()
	utils.Info(fmt.Sprintf("Encontrados %d pods", len(podsList)))

	for _, pod := range podsList {
		utils.Info(fmt.Sprintf("Pod: %s/%s - %s",
			pod.Namespace, pod.Name, pod.Status))
	}

	// ❌ Isto não compila (validatePod não é exportada)
	// pods.validatePod(pods.Pod{Name: "teste"})
}
