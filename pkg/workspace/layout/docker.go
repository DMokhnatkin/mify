package layout

func (gl GoLayout) Dockerfile(serviceName string) WorkspaceFile {
	return gl.Cmd(serviceName).file("Dockerfile")
}

func (jsl JsServiceLayout) Dockerfile() WorkspaceFile {
	return jsl.path.file("Dockerfile")
}
