package authenticators

import (
	"net/http"
	"strings"
	"text/template"

	api_utils "github.com/Cyarun/CyFir/api/utils"
	utils "github.com/Cyarun/CyFir/api/utils"
	config_proto "github.com/Cyarun/CyFir/config/proto"
	"github.com/Cyarun/CyFir/gui/velociraptor"
	gui_assets "github.com/Cyarun/CyFir/gui/velociraptor"
	"github.com/Cyarun/CyFir/json"
)

func renderRejectionMessage(
	config_obj *config_proto.Config,
	r *http.Request, w http.ResponseWriter, err error,
	username string, authenticators []velociraptor.AuthenticatorInfo) {

	// For API calls we render the error as JSON
	base_path := api_utils.GetBasePath(config_obj, "/api/")
	if strings.HasPrefix(r.URL.Path, base_path) {
		_, _ = w.Write([]byte(json.Format(`{"message": %q}`, err.Error())))
		return
	}

	data, err := gui_assets.ReadFile("/index.html")
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
		return
	}

	tmpl, err := template.New("").Parse(string(data))
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
		return
	}

	err = tmpl.Execute(w, velociraptor.HTMLtemplateArgs{
		BasePath: utils.GetBasePath(config_obj),
		ErrState: json.MustMarshalString(velociraptor.ErrState{
			Type:           "Login",
			Username:       username,
			Authenticators: authenticators,
			BasePath:       utils.GetBasePath(config_obj),
		}),
	})
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
	}
}

func renderLogoffMessage(
	config_obj *config_proto.Config,
	w http.ResponseWriter, username string) {
	data, err := gui_assets.ReadFile("/index.html")
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
		return
	}

	tmpl, err := template.New("").Parse(string(data))
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
		return
	}

	err = tmpl.Execute(w, velociraptor.HTMLtemplateArgs{
		BasePath: utils.GetBasePath(config_obj),
		ErrState: json.MustMarshalString(velociraptor.ErrState{
			Type:           "Logoff",
			Username:       username,
			BasePath:       utils.GetBaseDirectory(config_obj),
			Authenticators: []velociraptor.AuthenticatorInfo{},
		}),
	})
	if err != nil {
		//utils.Debug(err)
		w.WriteHeader(500)
	}
}
