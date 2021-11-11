package workspace

import (
	"fmt"
	"path/filepath"

	"github.com/chebykinn/mify/internal/mify/config"
	"github.com/chebykinn/mify/internal/mify/core"
)

func CreateWorkspace(ctx *core.Context, dir string, name string) error {
	fmt.Printf("creating workspace %s\n", name)

	context := Context{
		Name:     name,
		BasePath: filepath.Join(dir, name),
		GoRoot:   filepath.Join(dir, "go_services"),
	}

	if err := RenderTemplateTree(ctx, context); err != nil {
		return err
	}

	if err := createYaml(filepath.Join(dir, name)); err != nil {
		return err
	}
	return nil
}

// private

func createYaml(dir string) error {
	fmt.Printf("creating yaml in %s\n", dir)

	conf := config.WorkspaceConfig{
		WorkspaceName: dir,
		GitHost: "repo.com",
		GitNamespace: "namespace",
		GitRepository: "somerepo",
	}
	return config.SaveWorkspaceConfig(dir, conf)
}

