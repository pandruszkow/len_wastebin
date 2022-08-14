package apiv1

import (
	"net/http"

	"github.com/coolguy1771/wastebin/internal/netshare"
)

// GET /api/v1/
func (data Data) MainHand(rw http.ResponseWriter, req *http.Request) {
	data.writeError(rw, req, netshare.ErrNotFound)
}
