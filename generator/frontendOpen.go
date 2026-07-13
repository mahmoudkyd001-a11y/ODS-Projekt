package generator

import (
	fs "dredger/fileUtils"
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rs/zerolog/log"
)

func generateEmptyFrontend(_ *openapi3.T, conf GeneratorConfig) {
	frontendPath := filepath.Join(conf.OutputPath, "web")
	fs.GenerateFolder(frontendPath)
	createFileFromTemplate(filepath.Join(frontendPath, "README.md"), "templates/common/web/README.md.tmpl", conf)
}

func generateFrontend(spec *openapi3.T, conf GeneratorConfig) {
	generateOpenAPIDoc(conf)

	// NEU am 01.07
	if authExt, ok := spec.Info.Extensions["x-ui-auth"]; ok {
		if authMap, ok := authExt.(map[string]interface{}); ok {
			if totpVal, ok := authMap["totp"].(bool); ok {
				conf.AuthConfig.AddTOTP = totpVal
			}
		}
	}

	// create folders
	restPath := filepath.Join(conf.OutputPath, "rest")
	frontendPath := filepath.Join(conf.OutputPath, "web")
	javascriptPath := filepath.Join(frontendPath, "js")
	stylesheetPath := filepath.Join(frontendPath, "css")
	imagesPath := filepath.Join(frontendPath, "images")
	fontsPath := filepath.Join(stylesheetPath, "fonts")
	pagesPath := filepath.Join(frontendPath, "pages")
	localesPath := filepath.Join(pagesPath, "locales")
	publicPath := filepath.Join(frontendPath, "public")
	docPath := filepath.Join(frontendPath, "doc")

	fs.GenerateFolder(frontendPath)
	fs.GenerateFolder(javascriptPath)
	fs.GenerateFolder(stylesheetPath)
	fs.GenerateFolder(imagesPath)
	fs.GenerateFolder(fontsPath)
	fs.GenerateFolder(pagesPath)
	fs.GenerateFolder(localesPath)
	fs.GenerateFolder(publicPath)
	fs.GenerateFolder(docPath)

	// files in root directory
	createFileFromTemplate(filepath.Join(frontendPath, "README.md"), "templates/common/web/README.md.tmpl", conf)

	// files in javascript directory
	fs.CopyWebFile("common/web/js", javascriptPath, "bootstrap.bundle.min.js", true)
	fs.CopyWebFile("common/web/js", javascriptPath, "htmx.min.js", true)
	fs.CopyWebFile("common/web/js", javascriptPath, "hyperscript.js", true)
	fs.CopyWebFile("common/web/js", javascriptPath, "sse.js", true)
	fs.CopyWebFile("common/web/js", javascriptPath, "rapidoc-min.js", true)
	fs.CopyWebFile("common/web/js", javascriptPath, "elements.min.js", true)

	// files in stylesheet directory
	fs.CopyWebFile("common/web/css", stylesheetPath, "bootstrap-icons.min.css", true)
	fs.CopyWebFile("common/web/css/fonts", fontsPath, "bootstrap-icons.woff", true)
	fs.CopyWebFile("common/web/css/fonts", fontsPath, "bootstrap-icons.woff2", true)
	fs.CopyWebFile("common/web/css", stylesheetPath, "bootstrap.min.css", true)
	fs.CopyWebFile("common/web/css", stylesheetPath, "pico.min.css", true)
	fs.CopyWebFile("common/web/css", stylesheetPath, "pico.colors.min.css", true)
	fs.CopyWebFile("common/web/css", stylesheetPath, "elements.min.css", true)

	// files in images directory
	fs.CopyWebFile("common/web/images", imagesPath, "favicon.ico", false)

	// files in web directory
	fs.CopyWebFile("common/web", frontendPath, "web.go", true)

	// files in pages directory
	fs.CopyWebFile("common/web/pages", restPath, "render.go", true)
	if _, err := os.Stat(filepath.Join(pagesPath, "languages.templ")); errors.Is(err, os.ErrNotExist) {
		createFileFromTemplate(filepath.Join(pagesPath, "languages.templ"), "templates/common/web/pages/languages.templ.tmpl", conf)
	}
	if spec.Paths.Find("/index.html") != nil && spec.Paths.Find("/index.html").Operations()[http.MethodGet] != nil && slices.Contains(spec.Paths.Find("/index.html").Operations()[http.MethodGet].Tags, "builtin") {
		if _, err := os.Stat(filepath.Join(pagesPath, "index.templ")); errors.Is(err, os.ErrNotExist) {
			createFileFromTemplate(filepath.Join(pagesPath, "index.templ"), "templates/common/web/pages/index.templ.tmpl", conf)
			createFileFromTemplate(filepath.Join(pagesPath, "content.templ"), "templates/common/web/pages/content.templ.tmpl", conf)
		}
		op := openapi3.NewOperation()
		op.AddResponse(http.StatusOK, createOAPIResponse("The service delivers index page"))
		updateOAPIOperation(op, "GetIndex", "successfully deliver index page", "200")
		spec.AddOperation("/index.html", http.MethodGet, op)
	}
	if spec.Paths.Find("/") != nil && spec.Paths.Find("/").Operations()[http.MethodGet] != nil && slices.Contains(spec.Paths.Find("/").Operations()[http.MethodGet].Tags, "builtin") {
		op := openapi3.NewOperation()
		op.AddResponse(http.StatusOK, createOAPIResponse("The service delivers index page"))
		updateOAPIOperation(op, "GetRoot", "successfully deliver index page", "200")
		spec.AddOperation("/", http.MethodGet, op)
	}
	if spec.Paths.Find("/content.html") != nil && spec.Paths.Find("/content.html").Operations()[http.MethodGet] != nil && slices.Contains(spec.Paths.Find("/content.html").Operations()[http.MethodGet].Tags, "builtin") {
		if _, err := os.Stat(filepath.Join(pagesPath, "content.templ")); errors.Is(err, os.ErrNotExist) {
			createFileFromTemplate(filepath.Join(pagesPath, "content.templ"), "templates/common/web/pages/content.templ.tmpl", conf)
		}
		op := openapi3.NewOperation()
		op.AddResponse(http.StatusOK, createOAPIResponse("The service delivers content page"))
		updateOAPIOperation(op, "GetContent", "successfully deliver content page", "200")
		spec.AddOperation("/content.html", http.MethodGet, op)
	}

	// files in public directory
	fs.CopyWebFile(path.Join("common/web", "public"), publicPath, "README.md", false)

	// files in doc directory
	fs.CopyWebFile(path.Join("common/web", "doc"), docPath, "README.md", false)

	// support for events
	if spec.Paths.Find("/events") != nil && spec.Paths.Find("/events").Operations()[http.MethodGet] != nil && slices.Contains(spec.Paths.Find("/events").Operations()[http.MethodGet].Tags, "builtin") {
		//log.Debug().Msg("Generating default /events endpoint.")
		createFileFromTemplate(filepath.Join(restPath, "progress.go"), "templates/common/web/pages/progress.go.tmpl", conf)
		createFileFromTemplate(filepath.Join(restPath, "notice.go"), "templates/common/web/pages/notice.go.tmpl", conf)

		op := openapi3.NewOperation()
		op.AddResponse(http.StatusOK, createOAPIResponse("The service support sse"))
		updateOAPIOperation(op, "HandleEvents", "support for sse", "200")
		spec.AddOperation("/events", http.MethodGet, op)
		spec.AddOperation("/events", http.MethodPost, op)
	}

	// SSE server for real-time UI updates (progress bars, notifications)
	createFileFromTemplate(filepath.Join(conf.OutputPath, "core", "sse.go"), "templates/common/core/sse.go.tmpl", conf)
	createFileFromTemplate(filepath.Join(restPath, "handleEvents.go"), "templates/openapi/rest/handleEvents.go.tmpl", conf)

	log.Info().Msg("Created Frontend successfully.")
	// NEU - Formulare generieren wenn Schemas mit x-label: "form" vorhanden
	schemas := createSchemas(spec)
	if schemas.IsNotEmpty {
		type FormConfig struct {
			GeneratorConfig
			Schemas Schemas
		}
		formConf := FormConfig{
			GeneratorConfig: conf,
			Schemas:         schemas,
		}
		createFileFromTemplate(
			filepath.Join(pagesPath, "layout.templ"),
			"templates/common/web/pages/layout.templ.tmpl",
			formConf,
		)
		createFileFromTemplate(
			filepath.Join(pagesPath, "form.templ"),
			"templates/common/web/pages/form.templ.tmpl",
			formConf,
		)
		log.Info().Msg("Generated form templates from schema.")
	}

	// am 01.07
	if conf.AuthConfig.AddTOTP {
		createFileFromTemplate(filepath.Join(restPath, "authHandler.go"), "templates/openapi/rest/authHandler.go.tmpl", conf)
		createFileFromTemplate(filepath.Join(restPath, "handlerExt.go"), "templates/openapi/rest/handlerExt.go.tmpl", conf)
		createFileFromTemplate(filepath.Join(pagesPath, "register.templ"), "templates/common/web/pages/register.templ.tmpl", conf)
		createFileFromTemplate(filepath.Join(pagesPath, "login.templ"), "templates/common/web/pages/login.templ.tmpl", conf)
		createFileFromTemplate(filepath.Join(pagesPath, "totp_setup.templ"), "templates/common/web/pages/totp_setup.templ.tmpl", conf)
		createFileFromTemplate(filepath.Join(pagesPath, "totp_verify.templ"), "templates/common/web/pages/totp_verify.templ.tmpl", conf)
		log.Info().Msg("Generated TOTP auth templates.")
	}

}

// function to get the port specified in the OpenAPI Spec
func getServerPort(spec *openapi3.T) (port int16) {
	if spec.Servers != nil {
		serverSpec := spec.Servers[0]
		if portSpec := serverSpec.Variables["port"]; portSpec != nil {
			portStr := portSpec.Default
			if portSpec.Enum != nil {
				portStr = portSpec.Enum[0]
			}

			port, err := strconv.ParseInt(portStr, 10, 16)
			if err != nil {
				log.Warn().Msg("Failed to convert port, using 8080 instead.")
				return 8080
			} else {
				return int16(port)
			}
		} else {
			log.Warn().Msg("Failed to convert port, using 8080 instead.")
			return 8080
		}
	} else {
		log.Warn().Msg("Failed to convert port, using 8080 instead.")
		return 8080
	}
}

func createSchemas(spec *openapi3.T) (schemas Schemas) {
	schemas.List = make([]SchemaConf, 0)
	schemas.IsNotEmpty = false

	if spec != nil && spec.Components != nil && spec.Components.Schemas != nil {
		schemaStrings := toString(reflect.ValueOf(spec.Components.Schemas).MapKeys())

		for i := range schemaStrings {
			tmpSchemaName := schemaStrings[i]

			// check if schema has x-label == "form" -> if yes add schema to list
			schemaInformation, _ := spec.Components.Schemas[tmpSchemaName].Value.MarshalJSON()
			if strings.Contains(string(schemaInformation[:]), "\"x-label\":\"form\"") {
				var schema SchemaConf

				// add names
				schema.Name = strings.ReplaceAll(strings.ToLower(tmpSchemaName), " ", "")
				schema.H1Name = strings.Title(tmpSchemaName)
				schema.ComponentName = strings.ReplaceAll(schema.H1Name, " ", "")

				// add properties
				schema.Properties = make([]PropertyConf, 0)
				tmpSchemaPropertyNames := reflect.ValueOf(spec.Components.Schemas[tmpSchemaName].Value.Properties).MapKeys()
				for j := range tmpSchemaPropertyNames {
					tmpSchemaPropertyName := tmpSchemaPropertyNames[j].Interface().(string)
					var tmpPropertyConf PropertyConf
					tmpPropertyConf.Name = tmpSchemaPropertyName
					tmpPropertyConf.LabelName = strings.Title(tmpSchemaPropertyName)
					types := spec.Components.Schemas[tmpSchemaName].Value.Properties[tmpSchemaPropertyName].Value.Type.Slice()
					if len(types) > 0 {
						tmpPropertyConf.Type = types[0]
					}

					// NEU - x-ui Attribute aus der YAML lesen
					propValue := spec.Components.Schemas[tmpSchemaName].Value.Properties[tmpSchemaPropertyName].Value
					propJSON, _ := propValue.MarshalJSON()
					propStr := string(propJSON)

					// x-ui-control lesen (default: "input")
					if strings.Contains(propStr, "\"x-ui-control\":") {
						start := strings.Index(propStr, "\"x-ui-control\":\"") + len("\"x-ui-control\":\"")
						end := strings.Index(propStr[start:], "\"") + start
						tmpPropertyConf.UIControl = propStr[start:end]
					} else {
						tmpPropertyConf.UIControl = "input"
					}

					// x-ui-required lesen
					tmpPropertyConf.UIRequired = strings.Contains(propStr, "\"x-ui-required\":true")

					// x-ui-group lesen
					if strings.Contains(propStr, "\"x-ui-group\":") {
						start := strings.Index(propStr, "\"x-ui-group\":\"") + len("\"x-ui-group\":\"")
						end := strings.Index(propStr[start:], "\"") + start
						tmpPropertyConf.UIGroup = propStr[start:end]
					}

					if strings.Contains(propStr, "\"x-ui-order\":") {
						start := strings.Index(propStr, "\"x-ui-order\":") + len("\"x-ui-order\":")
						end := start
						for end < len(propStr) && propStr[end] != ',' && propStr[end] != '}' {
							end++
						}
						orderStr := strings.TrimSpace(propStr[start:end])
						if order, err := strconv.Atoi(orderStr); err == nil {
							tmpPropertyConf.UIOrder = order
						}
					}
					// x-ui-options lesen (für Dropdown-Optionen)

					if strings.Contains(propStr, "\"x-ui-options\":") {
						start := strings.Index(propStr, "\"x-ui-options\":[") + len("\"x-ui-options\":[")
						end := strings.Index(propStr[start:], "]") + start
						optionsStr := propStr[start:end]
						options := strings.Split(optionsStr, ",")
						for _, opt := range options {
							opt = strings.TrimSpace(opt)
							opt = strings.Trim(opt, "\"")
							if opt != "" {
								tmpPropertyConf.UIOptions = append(tmpPropertyConf.UIOptions, opt)
							}
						}
					}

					// x-ui-placeholder lesen
					if strings.Contains(propStr, "\"x-ui-placeholder\":") {
						start := strings.Index(propStr, "\"x-ui-placeholder\":\"") + len("\"x-ui-placeholder\":\"")
						end := strings.Index(propStr[start:], "\"") + start
						tmpPropertyConf.UIPlaceholder = propStr[start:end]
					}

					// x-ui-label lesen (fallback: automatisch aus Feldname)
					if strings.Contains(propStr, "\"x-ui-label\":") {
						start := strings.Index(propStr, "\"x-ui-label\":\"") + len("\"x-ui-label\":\"")
						end := strings.Index(propStr[start:], "\"") + start
						tmpPropertyConf.UILabel = propStr[start:end]
					} else {
						tmpPropertyConf.UILabel = tmpPropertyConf.LabelName
					}

					schema.Properties = append(schema.Properties, tmpPropertyConf)
				}

				sort.Slice(schema.Properties, func(i, j int) bool {
					return schema.Properties[i].UIOrder < schema.Properties[j].UIOrder
				})

				// NEU - Felder nach Gruppen gruppieren
				groupMap := make(map[string]*GroupConf)
				groupOrder := []string{}
				ungrouped := GroupConf{Name: "", Properties: []PropertyConf{}}

				for _, prop := range schema.Properties {
					if prop.UIGroup == "" {
						ungrouped.Properties = append(ungrouped.Properties, prop)
					} else {
						if _, exists := groupMap[prop.UIGroup]; !exists {
							groupMap[prop.UIGroup] = &GroupConf{
								Name:       prop.UIGroup,
								Properties: []PropertyConf{},
							}
							groupOrder = append(groupOrder, prop.UIGroup)
						}
						groupMap[prop.UIGroup].Properties = append(groupMap[prop.UIGroup].Properties, prop)
					}
				}

				// Gruppen in richtiger Reihenfolge hinzufügen
				for _, groupName := range groupOrder {
					schema.Groups = append(schema.Groups, *groupMap[groupName])
				}

				// Felder ohne Gruppe am Ende hinzufügen
				if len(ungrouped.Properties) > 0 {
					schema.Groups = append(schema.Groups, ungrouped)
				}

				schema.HasGroups = len(schema.Groups) > 0
				schemas.List = append(schemas.List, schema)
				schemas.IsNotEmpty = true
			}

		}
	}

	return schemas

}

// function to convert an []reflect.Value to []string
func toString(inputArray []reflect.Value) (resultArray []string) {
	for i := range inputArray {
		resultArray = append(resultArray, inputArray[i].Interface().(string))
	}

	return resultArray
}

func generateOpenAPIDoc(conf GeneratorConfig) {

	// create folder
	type templateConfig struct {
		GeneratorConfig
		OpenAPIFile string
	}
	docPath := filepath.Join(conf.OutputPath, "web", "doc")
	fs.GenerateFolder(docPath)

	template := templateConfig{
		GeneratorConfig: conf,
		OpenAPIFile:     fs.GetFileNameWithEnding(conf.OpenAPIPath),
	}

	// create static html files
	createFileFromTemplate(filepath.Join(docPath, "rapidoc.html"), "templates/openapi/rapidoc/index.html.tmpl", template)
	createFileFromTemplate(filepath.Join(docPath, "elements.html"), "templates/openapi/elements/index.html.tmpl", template)

	// copy OpenAPI Specification in this directory
	fs.CopyFile(conf.OpenAPIPath, docPath, template.OpenAPIFile)
	// add symlink to project root
	specPath := filepath.Join(docPath, template.OpenAPIFile)
	linkFilename := "OpenAPI" + path.Ext(template.OpenAPIFile) // static filename for project root
	linkPath := filepath.Join(Config.Path, linkFilename)
	if !fs.CheckIfFileExists(linkPath) { // skip it file (symlink) already exists
		// symlink target must be relative to the link's own directory, not to the CWD
		relTarget, err := filepath.Rel(filepath.Dir(linkPath), specPath)
		if err != nil {
			relTarget = specPath
		}
		if err := os.Symlink(relTarget, linkPath); err != nil {
			log.Warn().Err(err).Str("source", relTarget).Str("target", linkPath).Msg("Failed to create Symlink for OpenAPI specification file")
		}
	}

	log.Info().Msg("Created OpenAPI Documentation successfully.")
}
