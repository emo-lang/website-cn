package layouts

import (
	"github.com/a-h/templ"
	"github.com/daqing/airway/lib/utils"
)

// href prefixes a site path with the configured URL_PREFIX, so links keep
// working when the app is mounted under a sub-path.
func href(path string) templ.SafeURL {
	return templ.URL(utils.URLPrefix() + path)
}

// Href is the exported form of href, for page templates outside this package.
func Href(path string) templ.SafeURL {
	return href(path)
}

// PublicURL returns the URL of a committed static file under
// app/assets/public, e.g. PublicURL("logo.png") → "/public/logo.png".
func PublicURL(name string) string {
	return utils.URLPrefix() + "/public/" + name
}
