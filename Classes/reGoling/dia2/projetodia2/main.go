package main

import (
	"fmt"

	// Importando nossos pacotes locais
	k8sclient "github.com/seu-usuario/projetodia2/k8s"
	"github.com/seu-usuario/projetodia2/utils"

	// Alias "_" para ignorar (útil para init())
	_ "github.com/lib/pq" // Inicializa o driver PostgreSQL
)

func main() {
	// Usando função exportada do pacote k8s
	podName := k8sclient.GetPodName("default", "nginx-pod")
	fmt.Println("Nome do Pod:", podName)

	// Usando função exportada do pacote utils
	upper := utils.UpperExportada("hello kubernetes")
	fmt.Println("Upper:", upper)

	// ❌ Isso NÃO funciona - parseYaml não é exportada
	// k8s.parseYaml("teste")  // Erro de compilação!

	// Usando struct exportada
	pod := k8s.PodInfo{
		Name:      "nginx-pod",
		Namespace: "default",
		Ready:     true,
	}
	fmt.Printf("Pod: %+v\n", pod)

	// ✅ CORRETO: usar o alias para simplificar
	// k8s é o nome do pacote, não precisa de alias a menos que haja conflito
}
