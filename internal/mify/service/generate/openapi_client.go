package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/chebykinn/mify/internal/mify/config"
	"github.com/chebykinn/mify/internal/mify/util"
)

func (g *OpenAPIGenerator) makeClientEnrichedSchema(ctx *util.JobPoolContext, schemaPath string) (string, error) {
	doc, err := g.readSchema(ctx, schemaPath)
	if err != nil {
		return "", fmt.Errorf("failed to read schema: %s: %w", schemaPath, err)
	}

	pathsIface, ok := doc["paths"]
	if !ok {
		return "", fmt.Errorf("missing paths in schema: %s", schemaPath)
	}
	// TODO mapstructure
	paths := pathsIface.(map[interface{}]interface{})
	for path, v := range paths {
		ctx.Logger.Printf("processing path: %s\n", path)
		methods := v.(map[interface{}]interface{})
		if _, ok := methods["$ref"]; ok {
			return "", fmt.Errorf("paths with $ref are not supported yet")
		}
		for m, vv := range methods {
			ctx.Logger.Printf("processing method: %s\n", m)
			method := vv.(map[interface{}]interface{})
			method["tags"] = []string{"api"}
			methods[m] = method
		}
	}

	return g.saveEnrichedSchema(ctx, doc, schemaPath, CACHE_CLIENT_SUBDIR)
}

func (g *OpenAPIGenerator) doGenerateClient(ctx *util.JobPoolContext, clientName string, schemaPath string, targetPath string) error {
	langStr := string(GENERATOR_LANGUAGE_GO)
	path, err := config.DumpAssets(g.basePath, "openapi/client-template/"+langStr, "openapi/client-template")
	if err != nil {
		return fmt.Errorf("failed to dump assets: %w", err)
	}
	ctx.Logger.Printf("dumped path: %s\n", path)

	generatedPath := filepath.Join(g.basePath, targetPath, "generated", "api", "clients", clientName)

	packageName := makePackageName(clientName)

	err = runOpenapiGenerator(ctx, g.basePath, schemaPath, path, generatedPath, packageName, g.info)
	if err != nil {
		return fmt.Errorf("failed to run openapi-generator: %w", err)
	}

	// err = sanitizeServerHandlersImports(apiPath)
	// if err != nil {
		// return err
	// }

	// handlersPath := filepath.Join(g.basePath, targetPath, "handlers")
	// err = moveServerHandlers(apiPath, handlersPath)
	// if err != nil {
		// return err
	// }
	err = os.Remove(filepath.Join(generatedPath, "api"))
	if err != nil {
		return err
	}

	err = formatGenerated(generatedPath)
	if err != nil {
		return err
	}

	return nil
}

func makePackageName(clientName string) string {
	packageName := clientName
	if unicode.IsDigit(rune(clientName[0])) {
		packageName = "service_" + packageName
	}
	packageName = strings.ReplaceAll(packageName, "-", "_")

	packageName = packageName + "_client"
	return packageName
}