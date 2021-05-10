package cas

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/caddyauth"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(CASAuthenticator{})
	httpcaddyfile.RegisterHandlerDirective("cas", parseCaddyfile)
}

// CASAuthenticator is an example; put your own type here.
type CASAuthenticator struct {
	CASVersion int    `json:"version,omitempty"`
	CASBaseURL string `json:"base_url,omitempty"`
	// CASLoginURL         string `json:"login_url,omitempty"`
	// CASValidateURL      string `json:"validate_url,omitempty"`
	// CASProxyValidateURL string `json:"proxy_validate_url,omitempty"`
	// CASCookiePath       string `json:"cookie_path,omitempty"`
	// CASTimeout          int    `json:"timeout,omitempty"`
	// CASIdleTimeout      int    `json:"idle_timeout,omitempty"`
	// CASCookie         string `json:"cookie,omitempty"`
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (CASAuthenticator) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.authentication.providers.cas",
		New: func() caddy.Module { return new(CASAuthenticator) },
	}
}

func (g *CASAuthenticator) Provision(ctx caddy.Context) error {
	g.logger = ctx.Logger(g) // g.logger is a *zap.Logger
	return nil
}

// Validate validates that the module has a usable config.
func (g *CASAuthenticator) Validate() error {
	// TODO: validate the module's setup
	return nil
}

func (g *CASAuthenticator) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for nesting := d.Nesting(); d.NextBlock(nesting); {
		if !d.Args(&g.CASBaseURL) {
			return d.ArgErr()
		}
		// if d.NextArg() {
		// 	// optional arg
		// 	g.CASVersion = d.Val()
		// }
		// if d.NextArg() {
		// 	// optional arg
		// 	g.CASValidateURL = d.Val()
		// }
		// if d.NextArg() {
		// 	// optional arg
		// 	g.Option = d.Val()
		// }
		// if d.NextArg() {
		// 	// optional arg
		// 	g.Option = d.Val()
		// }
		// if d.NextArg() {
		// 	// optional arg
		// 	g.Option = d.Val()
		// }
	}
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m CASAuthenticator
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return caddyauth.Authentication{
		ProvidersRaw: caddy.ModuleMap{
			"cas": caddyconfig.JSON(&m, nil),
		},
	}, err
}

func (m CASAuthenticator) Authenticate(w http.ResponseWriter, r *http.Request) (caddyauth.User, bool, error) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
	w.Write([]byte(`User Unauthorized`))
	return caddyauth.User{}, false, nil
	// return caddyauth.User{}, true, nil
}

// Interface guards
var (
	_ caddy.Provisioner       = (*CASAuthenticator)(nil)
	_ caddy.Validator         = (*CASAuthenticator)(nil)
	_ caddyauth.Authenticator = (*CASAuthenticator)(nil)
)
