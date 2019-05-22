/*
 * @author    Emmanuel Kofi Bessah
 * @email     bekinsoft@gmail.com
 */

package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// MarshalRequestBody ...
// Used LimitReader to protect against malicious attacks on your server.
// Imagine if someone wanted to send you 500GBs of json!
func MarshalRequestBody(r *http.Request, model interface{}, limit string) error {
	// size, err := ToBytes(viper.GetString("Server.Rest.Remoting.JSONLimit"))
	size, err := ToBytes(limit)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(io.LimitReader(r.Body, int64(size))).Decode(&model); err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") {
			return ErrRequestBodyLarge
		} else if strings.Contains(err.Error(), "EOF") {
			return ErrRequestBodyMalformed
		}

		return err
	}

	// bodyLength := r.ContentLength

	// fmt.Println(bodyLength)

	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, int64(size+1)))

	// if err != nil {
	// 	return err
	// }

	// // Check if the request body is not empty
	// if len(body) > 0 {
	// 	if err := json.Unmarshal(body, &model); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// ExtractHeaderToken gets or extracts token from header
func ExtractHeaderToken(header http.Header) (token string) {
	val := header.Get("Authorization")
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) == 2 && strings.ToLower(authHeaderParts[0]) == "bearer" {
		token = authHeaderParts[1]
	}

	return
}

// RemoveDuplicatesFromSlice removes duplicates from string array
func RemoveDuplicatesFromSlice(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// TokenGenerator generates token or session for mobile and desktop client
func TokenGenerator() string {
	b := make([]byte, 42)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// LogDatabaseQueries prints to console queries of db
func LogDatabaseQueries(tx *gorm.DB) {
	if viper.Get("profile") != "prod" {
		tx.LogMode(true)
	}
}
