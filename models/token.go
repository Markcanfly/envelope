package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// Tokens stored in memory, but in real life it should be stored in database
type TokenData struct {
	User *User
	timeout int64
}
const TTL = 5 * 60//seconds
const TokenLength = 64
var tokens map[string]TokenData = make(map[string]TokenData)

func StartTokenCleanup() {
	go func() {
		for {
			time.Sleep(time.Second * TTL)
			for token, tokenData := range tokens {
				if tokenData.timeout < time.Now().Unix() {
					delete(tokens, token)
				}
			}
		}
	}()
}

// See if a user is already authenticated with a token
func GetTokenData(token string) *TokenData {
	if tokenData, ok := tokens[token]; ok {
		if tokenData.timeout > time.Now().Unix() {
			tokenData.timeout = time.Now().Unix() + TTL
			return &tokenData
		} else {
			delete(tokens, token)
		}
	}
	return nil
}

// Only used when user is already authenticated
func CreateTokenFor(u *User) string {
	token := GenerateSecureToken(TokenLength)
	for GetTokenData(token) != nil { // Make sure our our token is unique
		token = GenerateSecureToken(TokenLength)
	}
	tokens[token] = TokenData{u, time.Now().Unix() + TTL}
	return token
}

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}


