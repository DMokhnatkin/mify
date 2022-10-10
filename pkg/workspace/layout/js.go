package layout

type JsLayout struct {
	workspace WorkspaceLayout
	path      WorkspaceFolder
}

func (wl WorkspaceLayout) JsServices() JsLayout {
	return JsLayout{
		workspace: wl,
		path:      wl.path.append("js-services"),
	}
}

func (gl JsLayout) Root() WorkspaceFolder {
	return gl.path
}

func (gl JsLayout) JsPackageJson() WorkspaceFile {
	return gl.path.file("package.json")
}

// Js service

type JsServiceLayout struct {
	js   JsLayout
	path WorkspaceFolder
}

func (jl JsLayout) JsService(serviceName string) JsServiceLayout {
	return JsServiceLayout{
		js:   jl,
		path: jl.path.append(serviceName),
	}
}

func (jsl JsServiceLayout) PackageJson() WorkspaceFile {
	return jsl.path.file("package.json")
}

func (jsl JsServiceLayout) YarnLock() WorkspaceFile {
	return jsl.path.file("yarn.lock")
}

func (jsl JsServiceLayout) NuxtConfig() WorkspaceFile {
	return jsl.path.file("nuxt.config.js")
}

func (jsl JsServiceLayout) Pages() WorkspaceFolder {
	return jsl.path.append("pages")
}

func (jsl JsServiceLayout) PagesIndex() WorkspaceFile {
	return jsl.path.append("pages").file("index.vue")
}

func (jsl JsServiceLayout) Components() WorkspaceFolder {
	return jsl.path.append("components")
}

func (jsl JsServiceLayout) ComponentsSampleVue() WorkspaceFile {
	return jsl.path.append("components").file("sample.vue")
}

func (jsl JsServiceLayout) Generated() WorkspaceFolder {
	return jsl.path.append("generated")
}
