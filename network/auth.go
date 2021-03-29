package network

import (
	"encoding/json"

	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
)

func PostNewAuth(name, pub string) bool {
	m := map[string]string{"username": name, "pub": pub}
	asBytes, _ := json.Marshal(m)
	return DoPost("auth", asBytes) != ""
}

func CreateUserKey(c *gin.Context) {
	m := mapJsonBody(c)
	name := m["username"]
	pub := m["pub"]
	sDec, _ := b64.StdEncoding.DecodeString(pub)
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	if len(universes[uids[uidIndex]].UsernameKeys[name]) == 0 {
		universes[uids[uidIndex]].UsernameKeys[name] = sDec
		c.JSON(200, gin.H{"ok": true})
	} else {
		c.JSON(422, gin.H{"ok": false})
	}
}