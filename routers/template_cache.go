package routers

import (
	"fmt"
	"html/template"
)

// TemplateCache holds all compiled templates
var TemplateCache = make(map[string]*template.Template)

// InitializeTemplates loads and caches all templates at startup
func InitializeTemplates() error {
	templates := map[string]string{
		"home":               "templates/home",
		"login":              "templates/log_in",
		"signup":             "templates/create_account",
		"account":            "templates/account",
		"account_redirect":   "templates/account_redirect",
		"admin":              "templates/admin",
		"createsheetfailed":  "templates/createsheetfailed",
		"createsheetsuccess": "templates/createsheetsuccess",
	}

	for name, path := range templates {
		t, err := template.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s at %s: %w", name, path, err)
		}
		TemplateCache[name] = t
	}

	return nil
}

// GetTemplate retrieves a cached template by name
func GetTemplate(name string) (*template.Template, error) {
	t, exists := TemplateCache[name]
	if !exists {
		return nil, fmt.Errorf("template %s not found in cache", name)
	}
	return t, nil
}
