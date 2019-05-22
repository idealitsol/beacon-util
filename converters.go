/*
 * @author    Emmanuel Kofi Bessah
 * @email     bekinsoft@gmail.com
 */

package util

import (
	"bytes"
	"fmt"
)

// CreateKeyValuePairs converts map to string
func CreateKeyValuePairs(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		switch value.(type) {
		case string:
			fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
		case int64:
			fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
		case bool:
			fmt.Fprintf(b, "%s=\"%b\"\n", key, value)
		}

	}
	return b.String()
}
