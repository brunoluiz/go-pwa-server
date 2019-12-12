package envjs

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func EnvToJS(prefix string, key string) string {
	values := []string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := pair[0]
		value := pair[1]

		if !strings.HasPrefix(key, prefix) {
			continue
		}

		if value == "true" || value == "false" {
			values = append(values, key+":"+value)
			continue
		}

		values = append(values, key+":\""+value+"\"")
	}

	return "window." + key + "={" + strings.Join(values, ",") + "}"
}

type EnvJS struct {
	prefix string
	key    string
}

func Handler(prefix string, key string) *EnvJS {
	return &EnvJS{prefix, key}
}

func (h *EnvJS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprint(w, EnvToJS(h.prefix, h.key))
}