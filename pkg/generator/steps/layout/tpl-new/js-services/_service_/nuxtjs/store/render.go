package store

import (
	"embed"
	"path"

	gencontext "github.com/mify-io/mify/pkg/generator/gen-context"
	"github.com/mify-io/mify/pkg/util/render"
)

//go:embed *.tpl
var templates embed.FS

func Render(ctx *gencontext.GenContext) error {
	curPath := path.Join(ctx.GetWorkspace().BasePath, ctx.GetWorkspace().GetJsServiceRelPath(ctx.GetServiceName()), "store")
	return render.RenderMany(
		templates,
		render.NewFile(ctx, path.Join(curPath, "index.js")).SetFlags(render.NewFlags().SkipExisting()),
		render.NewFile(ctx, path.Join(curPath, "mifycontext.js")).SetFlags(render.NewFlags().SkipExisting()),
	)
}
