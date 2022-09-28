package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCahce      bool
	TemplateCahce map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
