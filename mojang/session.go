package mojang

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/spookspace/plasma/sha1"
	"net/http"
)

type profile struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties []struct {
		Name      string `json:"name"`
		Value     string `json:"value"`
		Signature string `json:"signature"`
	} `json:"properties"`
}

func URLSessionServerHasJoined(username, sessionHash string) string {
	return fmt.Sprintf(
		"https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%s&serverId=%s",
		username,
		sessionHash,
	)
}

func URLSessionServerHasJoinedWithIP(username, sessionHash, ip string) string {
	return fmt.Sprintf("%s&ip=%s",
		URLSessionServerHasJoined(username, sessionHash),
		ip,
	)
}

func SessionHash(serverID string, sharedSecret, publicKey []byte) string {
	notchHash := sha1.New()
	notchHash.Update([]byte(serverID))
	notchHash.Update(sharedSecret)
	notchHash.Update(publicKey)
	return notchHash.HexDigest()
}

func AuthenticateSession(username, sessionHash string) ([16]byte, string, string, error) {
	return AuthenticateSessionPreventProxy(username, sessionHash, "")
}

func AuthenticateSessionPreventProxy(username, sessionHash, ip string) ([16]byte, string, string, error) {
	var url string
	if ip == "" {
		url = URLSessionServerHasJoined(username, sessionHash)
	} else {
		url = URLSessionServerHasJoinedWithIP(username, sessionHash, ip)
	}

	resp, err := http.Get(url)
	if err != nil {
		return [16]byte{}, "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return [16]byte{}, "", "", fmt.Errorf("unable to authenticate session (%s)", resp.Status)
	}

	var p profile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return [16]byte{}, "", "", err
	}
	_ = resp.Body.Close()

	id, err := uuid.FromString(p.ID)
	if err != nil {
		return [16]byte{}, "", "", err
	}

	var skin string
	var signature string
	for _, property := range p.Properties {
		if property.Name == "textures" {
			skin = property.Value
			signature = property.Signature
			break
		}
	}

	if skin == "" {
		return [16]byte{}, "", "", errors.New("no skin in request")
	}

	return id, skin, signature, nil
}
