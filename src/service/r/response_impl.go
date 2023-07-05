package r

import (
	"encoding/json"
	"golang-test/common"
	"net/http"
)

func (g *responseHandlerImpl) ToJson(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	// all response finish here
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		g.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	value, ok := r.Context().Value(common.RequestKey).(*common.RequestType)
	if !ok {
		g.logger.Error("Invalid request key")
		return
	}
	value.Code = code
	reqInfo, _ := json.Marshal(value)
	g.logger.Info(string(reqInfo))
}
