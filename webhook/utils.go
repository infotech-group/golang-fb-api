package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"io/ioutil"
	"net/http"
)

func calculateHash(r *http.Request, secret []byte) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	hash := hmac.New(sha1.New, secret)
	hash.Write(body)
	return hash.Sum(nil), nil
}
