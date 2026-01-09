package discord

import (
	"crypto/ed25519"
	"encoding/hex"
)

// VerifySignature checks the signature of the incoming request from Discord.
func VerifySignature(signature, timestamp string, body []byte, publicKey ed25519.PublicKey) bool {
	decodedSignature, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	message := []byte(timestamp)
	message = append(message, body...)

	return ed25519.Verify(publicKey, message, decodedSignature)
}
