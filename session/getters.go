package session

import (
	"net/http"
	"net/url"

	"github.com/kotahu/go-librespot/ap"
	"github.com/kotahu/go-librespot/audio"
	"github.com/kotahu/go-librespot/dealer"
	"github.com/kotahu/go-librespot/spclient"
)

func (s *Session) Username() string {
	return s.ap.Username()
}

func (s *Session) StoredCredentials() []byte {
	return s.ap.StoredCredentials()
}

func (s *Session) Spclient() *spclient.Spclient {
	return s.sp
}

func (s *Session) AudioKey() *audio.KeyProvider {
	return s.audioKey
}

func (s *Session) Dealer() *dealer.Dealer {
	return s.dealer
}

func (s *Session) Accesspoint() *ap.Accesspoint {
	return s.ap
}

func (s *Session) WebApi(method string, path string, query url.Values, header http.Header, body []byte) (*http.Response, error) {
	return s.sp.WebApiRequest(method, path, query, header, body)
}
