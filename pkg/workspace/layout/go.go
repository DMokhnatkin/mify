package layout

type GoLayout struct {
	workspace WorkspaceLayout
	path      WorkspaceFolder
}

func (wl WorkspaceLayout) GoServices() GoLayout {
	return GoLayout{
		workspace: wl,
		path:      wl.path.append("go-services"),
	}
}

func (gl GoLayout) Root() WorkspaceFolder {
	return gl.path
}

func (gl GoLayout) GoMod() WorkspaceFile {
	return gl.path.file("go.mod")
}

func (gl GoLayout) GoSum() WorkspaceFile {
	return gl.path.file("go.sum")
}

func (gl GoLayout) App(serviceName string) WorkspaceFolder {
	return gl.path.append("internal", serviceName, "app")
}

func (gl GoLayout) Generated(serviceName string) WorkspaceFolder {
	return gl.path.append("internal", serviceName, "generated")
}

func (gl GoLayout) GeneratedApp(serviceName string) WorkspaceFolder {
	return gl.path.append("internal", serviceName, "generated", "app")
}

func (gl GoLayout) GeneratedCore(serviceName string) WorkspaceFolder {
	return gl.path.append("internal", serviceName, "generated", "core")
}

func (gl GoLayout) PostgresConfig() WorkspaceFolder {
	return gl.path.append("internal", "pkg", "generated", "postgres")
}

func (gl GoLayout) Cmd(serviceName string) WorkspaceFolder {
	return gl.path.append("cmd", serviceName)
}
