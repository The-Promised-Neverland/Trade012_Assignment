package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RewardHandler struct {
	service     *services.RewardService
	userchecker *services.UserService
	log         *logger.Logger
}

func NewRewardHandler(s *services.RewardService, userChecker *services.UserService, log *logger.Logger) *RewardHandler {
	return &RewardHandler{
		service:     s,
		userchecker: userChecker,
		log:         log,
	}
}

type rewardRequest struct {
	UserID int64   `json:"user_id" binding:"required"`
	Stock  string  `json:"stock" binding:"required"`
	Shares float64 `json:"shares" binding:"required"`
}

func (h *RewardHandler) RewardUser(c *gin.Context) {
	var body rewardRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		h.log.WithFields(logrus.Fields{
			"user_id": body.UserID,
			"stock":   body.Stock,
			"shares":  body.Shares,
			"error":   err,
		}).Error("Failed to bind reward request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	_, err := h.userchecker.GetUser(context.Background(), body.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	id, err := h.service.RewardUser(c, body.UserID, body.Stock, body.Shares)
	if err != nil {
		h.log.WithFields(logrus.Fields{
			"user_id": body.UserID,
			"stock":   body.Stock,
			"shares":  body.Shares,
			"error":   err,
		}).Error("Failed to reward user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not reward user"})
		return
	}

	h.log.WithFields(logrus.Fields{
		"user_id":   body.UserID,
		"stock":     body.Stock,
		"shares":    body.Shares,
		"reward_id": id,
	}).Info("Reward granted successfully")

	c.JSON(http.StatusOK, gin.H{
		"reward_id": id,
		"message":   "reward granted",
	})
}

func (h *RewardHandler) GetTodayRewards(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID == 0 {
		h.log.WithField("user_id_param", userIDStr).Warn("Invalid userId parameter for today rewards")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	rewards, err := h.service.GetTodayRewards(c, userID)
	if err != nil {
		h.log.WithField("user_id", userID).Errorf("Failed to fetch today's rewards: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch today's rewards"})
		return
	}

	h.log.WithField("user_id", userID).Infof("Fetched today's rewards successfully, count=%d", len(rewards))
	c.JSON(http.StatusOK, rewards)
}

func (h *RewardHandler) GetHistorialRecords(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil || userID == 0 {
		h.log.WithField("user_id_param", c.Param("userId")).Warn("Invalid userId parameter for all rewards")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	rewards, err := h.service.GetHistoricalRecord(c, userID)
	if err != nil {
		h.log.WithField("user_id", userID).Errorf("Failed to fetch historical rewards: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch historical rewards"})
		return
	}

	h.log.WithField("user_id", userID).Infof("Fetched all rewards successfully, count=%d", len(rewards))
	c.JSON(http.StatusOK, rewards)
}

func (h *RewardHandler) GetStats(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		h.log.WithField("user_id_param", c.Param("userId")).Warn("Invalid userId parameter for stats")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	stats, err := h.service.GetStats(c, userID)
	if err != nil {
		h.log.WithField("user_id", userID).Errorf("Failed to fetch stats: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch stats"})
		return
	}

	h.log.WithField("user_id", userID).Info("Fetched stats successfully")
	c.JSON(http.StatusOK, stats)
}

func (h *RewardHandler) GetPortfolio(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		h.log.WithField("user_id_param", c.Param("userId")).Warn("Invalid userId parameter for stats")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	portfolio, err := h.service.GetPortfolio(c, userID)
	if err != nil {
		h.log.WithField("user_id", userID).Errorf("Failed to fetch portfolio: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch portfolio"})
		return
	}

	h.log.WithField("user_id", userID).Info("Fetched portfolio successfully")
	c.JSON(http.StatusOK, portfolio)
}
