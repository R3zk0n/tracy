package rest

import (
	"net/http"
	"tracy/api/common"
)

/*GetConfig gets the global configuration for the application. */
func GetConfig(w http.ResponseWriter, r *http.Request) {
	ret := []byte("{}")
	status := http.StatusInternalServerError

	if config, err := common.GetConfig(); err == nil {
		status = http.StatusOK
		ret = config
	}

	w.WriteHeader(status)
	w.Write(ret)
}
