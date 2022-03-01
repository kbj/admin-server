package v1

import (
	"admin-server/utils"
	"github.com/gofiber/fiber/v2"
)

func InitUserRoute(route *fiber.Router) {
	router := *route
	router.Get("/:id", getUser)
}

// getUser 查询某个用户信息
func getUser(context *fiber.Ctx) error {
	id := context.Params("id")
	return context.JSON(utils.ResponseSuccess("查询成功", id))
}
