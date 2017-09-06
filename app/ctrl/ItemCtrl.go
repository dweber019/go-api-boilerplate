package ctrl

import (
	"net/http"

	"github.com/dweber019/go-api-boilerplate/app/lib"
	"github.com/dweber019/go-api-boilerplate/app/models"
)

// GetAllUsersHandler ...
func GetAllItemsHandler(w http.ResponseWriter, req *http.Request) {
	res := lib.Response{ResponseWriter: w}
	item := new(models.Item)
	items := item.FetchAll()
	res.SendOK(items)
}
