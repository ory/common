package path

import (
	"net/url"
	"path"

	log "github.com/sirupsen/logrus"
)

func Join(u string, paths ...string) string {
	ur, err := url.Parse(u)
	if err != nil {
		log.WithError(err).WithField("url", u).Panic("Could not parse url")
	}
	ur.Path = path.Join(append([]string{ur.Path}, paths...)...)
	return ur.String()
}
