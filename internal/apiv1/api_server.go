package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
)

type serverInfoType struct {
	Version     string   `json:"version"`
	TitleMaxLen int      `json:"titleMaxlength"`
	BodyMaxLen  int      `json:"bodyMaxlength"`
	MaxLifeTime int64    `json:"maxLifeTime"`
	ServerAbout string   `json:"serverAbout"`
	ServerRules string   `json:"serverRules"`
	AdminName   string   `json:"adminName"`
	AdminMail   string   `json:"adminMail"`
	Syntaxes    []string `json:"syntaxes"`
}

// GET /api/v1/getServerInfo
func (data Data) GetServerInfoHand(rw http.ResponseWriter, req *http.Request) {
	// Log request
	data.Log.HttpRequest(req)

	// Check method
	if req.Method != "GET" {
		data.writeError(rw, req, netshare.ErrMethodNotAllowed)
		return
	}

	// Prepare data
	serverInfo := serverInfoType{
		TitleMaxLen: *data.TitleMaxLen,
		BodyMaxLen:  *data.BodyMaxLen,
		Version:     *data.Version,
		MaxLifeTime: *data.MaxLifeTime,
		ServerAbout: *data.ServerAbout,
		ServerRules: *data.ServerRules,
		AdminName:   *data.AdminName,
		AdminMail:   *data.AdminMail,
		Syntaxes:    *data.Lexers,
	}

	// Return response
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(serverInfo)
	if err != nil {
		data.Log.HttpError(req, err)
		return
	}
}
