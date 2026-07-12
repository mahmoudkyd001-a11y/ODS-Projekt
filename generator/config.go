package generator

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// generateConfigFiles legt .env, config.go, configSvc.go und version an.
func generateConfigFiles(serverConf ServerConfig) {
	// 1) .env
	fileName := ".env"
	filePath := filepath.Join(Config.Path, fileName)
	templateFile := "templates/common/core/app.env.tmpl" // falls du eine andere Vorlage willst, passe hier an
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		createFileFromTemplate(filePath, templateFile, serverConf)
	}

	// 2) config.go
	fileName = "config.go"
	filePath = filepath.Join(Config.Path, CorePkg, fileName)
	templateFile = "templates/common/core/config.go.tmpl"
	createFileFromTemplate(filePath, templateFile, serverConf)

	// 3) configSvc.go
	fileName = "configSvc.go"
	filePath = filepath.Join(Config.Path, CorePkg, fileName)
	templateFile = "templates/common/core/configSvc.go.tmpl"
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		createFileFromTemplate(filePath, templateFile, serverConf)
	}

	// 4) version (und Symlink)
	fileName = "version"
	filePath = filepath.Join(Config.Path, CorePkg, fileName)
	templateFile = "templates/common/core/version"
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		createFileFromTemplate(filePath, templateFile, serverConf)
		linkPath := filepath.Join(Config.Path, fileName)
		// symlink target must be relative to the link's own directory, not to the CWD
		relTarget, err := filepath.Rel(filepath.Dir(linkPath), filePath)
		if err != nil {
			relTarget = filePath
		}
		if err := os.Symlink(relTarget, linkPath); err != nil {
			log.Warn().Err(err).Str("source", relTarget).Str("target", linkPath).Msg("Could not create symbolic Link, please create it manually")
		}
	}
}
