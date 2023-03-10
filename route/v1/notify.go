package v1

import (
	"net/http"

	"github.com/IceWhaleTech/CasaOS-Common/model/notify"
	"github.com/IceWhaleTech/CasaOS/model"
	"github.com/IceWhaleTech/CasaOS/pkg/utils/common_err"
	"github.com/IceWhaleTech/CasaOS/service"
	"github.com/gin-gonic/gin"
)

func PostNotifyMessage(c *gin.Context) {
	path := c.Param("path")
	message := make(map[string]interface{})
	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusBadRequest, model.Result{Success: common_err.INVALID_PARAMS, Message: err.Error()})
		return
	}

	service.MyService.Notify().SendNotify(path, message)
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS)})
}

func PostSystemStatusNotify(c *gin.Context) {
	message := make(map[string]interface{})
	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusBadRequest, model.Result{Success: common_err.INVALID_PARAMS, Message: err.Error()})
		return
	}

	service.MyService.Notify().SettingSystemTempData(message)
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS)})
}

func PostInstallAppNotify(c *gin.Context) {
	app := notify.Application{}
	if err := c.ShouldBind(&app); err != nil {
		c.JSON(http.StatusBadRequest, model.Result{Success: common_err.INVALID_PARAMS, Message: err.Error()})
		return
	}

	service.MyService.Notify().SendInstallAppBySocket(app)
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS)})
}

func PostUninstallAppNotify(c *gin.Context) {
	app := notify.Application{}
	if err := c.ShouldBind(&app); err != nil {
		c.JSON(http.StatusBadRequest, model.Result{Success: common_err.INVALID_PARAMS, Message: err.Error()})
		return
	}

	service.MyService.Notify().SendUninstallAppBySocket(app)
	c.JSON(common_err.SUCCESS, model.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS)})
}
