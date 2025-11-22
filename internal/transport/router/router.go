package router

import (
	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/ThePromisedNeverland/021trade/internal/transport/handlers"
	"github.com/ThePromisedNeverland/021trade/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	rewardHandler *handlers.RewardHandler,
	ledgerHandler *handlers.LedgerHandler,
	userService *services.UserService,
) *gin.Engine {

	r := gin.Default()

	api := r.Group("/stocky")
	{
		userCheck := middleware.UserExistsMiddleware(userService)

		api.GET("/user/:userId", userCheck, userHandler.GetUser)

		api.POST("/reward", rewardHandler.RewardUser)
		
		api.GET("/today-stocks/:userId", userCheck, rewardHandler.GetTodayRewards)
		api.GET("/historical-inr/:userId", userCheck, rewardHandler.GetHistorialRecords)
		api.GET("/stats/:userId", userCheck, rewardHandler.GetStats)
		api.GET("/portfolio/:userId", userCheck, rewardHandler.GetPortfolio)

		api.GET("/ledger/:userId", ledgerHandler.GetUserLedger)
	}

	return r
}
