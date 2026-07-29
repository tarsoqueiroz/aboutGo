package pods

// ListPods - exportada
func ListPods() []Pod {
	// Simulando dados
	return []Pod{
		{Name: "nginx", Namespace: "default", Status: "Running"},
		{Name: "redis", Namespace: "cache", Status: "Pending"},
		{Name: "api", Namespace: "production", Status: "Running"},
	}
}

// validatePod - NÃO exportada
func validatePod(p Pod) bool {
	return p.Name != ""
}
