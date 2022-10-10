package layout

func (wl WorkspaceLayout) MifyCache() WorkspaceFolder {
	return wl.path.append(".mify")
}

func (wl WorkspaceLayout) Logs() WorkspaceFolder {
	return wl.path.append(".mify", "logs")
}

func (wl WorkspaceLayout) ServiceCache(serviceName string) WorkspaceFolder {
	return wl.path.append(".mify", "services", serviceName)
}
