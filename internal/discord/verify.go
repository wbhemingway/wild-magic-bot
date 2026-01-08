package discord

import (
	"crypto/ed25519"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)

// VerifySignature checks the signature of the incoming request from Discord.
func VerifySignature(r *http.Request, publicKey ed25519.PublicKey) bool {
	signature := r.Header.Get("X-Signature-Ed25519")
	timestamp := r.Header.Get("X-Signature-Timestamp")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false
	}

	// Re-add the body so it can be read again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	decodedSignature, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	message := []byte(timestamp + string(body))

	return ed25519.Verify(publicKey, message, decodedSignature)
}
