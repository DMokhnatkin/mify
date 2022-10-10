package layout

import "path"

type WorkspacePath struct {
	workspaceBasePathAbs string
	relativePath         string
}

type WorkspaceFolder struct {
	WorkspacePath
}

type WorkspaceFile struct {
	WorkspacePath
}

func (wp WorkspacePath) append(extraPaths ...string) WorkspaceFolder {
	newPath := wp.relativePath
	for _, p := range extraPaths {
		newPath = path.Join(newPath, p)
	}

	return WorkspaceFolder{
		WorkspacePath: WorkspacePath{
			workspaceBasePathAbs: wp.workspaceBasePathAbs,
			relativePath:         newPath,
		},
	}
}

func (wp WorkspacePath) file(name string) WorkspaceFile {
	return WorkspaceFile{
		WorkspacePath: WorkspacePath{
			workspaceBasePathAbs: wp.workspaceBasePathAbs,
			relativePath:         path.Join(wp.relativePath, name),
		},
	}
}

// Relative to workspace base directory path
func (wp WorkspacePath) Rel() string {
	return wp.relativePath
}

// Absolute path
func (wp WorkspacePath) Abs() string {
	return path.Join(wp.workspaceBasePathAbs, wp.relativePath)
}

type WorkspaceLayout struct {
	path WorkspacePath
}

func NewWorkspaceLayout(basePath string) WorkspaceLayout {
	return WorkspaceLayout{
		path: WorkspacePath{
			workspaceBasePathAbs: basePath,
			relativePath:         "",
		},
	}
}
