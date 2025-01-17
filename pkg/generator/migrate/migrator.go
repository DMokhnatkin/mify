package migrate

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	gencontext "github.com/mify-io/mify/pkg/generator/gen-context"
	"go.uber.org/zap"
)

func MigrateSubstring(ctx *gencontext.GenContext, root string, prefix string, excludePrefix string, replace string) {
	_ = filepath.Walk(root, func(path string, info fs.FileInfo, _ error) error {
		dat, err := os.ReadFile(path)
		if err != nil {
			ctx.Logger.Warn("can't read file for migration "+path, zap.Error(err))
			return nil
		}

		if strings.Contains(string(dat), excludePrefix) {
			return nil
		}

		newDat := strings.ReplaceAll(string(dat), prefix, replace)
		if newDat == string(dat) {
			return nil
		}

		err = os.WriteFile(path, []byte(newDat), info.Mode())
		if err != nil {
			ctx.Logger.Warn("can't write file for migration "+path, zap.Error(err))
			return nil
		}

		return nil
	})
}
