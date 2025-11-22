package handlers

import (
	"net/http"
	"strconv"

	"github.com/ThePromisedNeverland/021trade/internal/logger"
	"github.com/ThePromisedNeverland/021trade/internal/services"
	"github.com/gin-gonic/gin"
)

type LedgerHandler struct {
	ledgerService *services.LedgerService
	log           *logger.Logger
}

func NewLedgerHandler(ls *services.LedgerService, log *logger.Logger) *LedgerHandler {
	return &LedgerHandler{
		ledgerService: ls,
		log:           log,
	}
}

func (h *LedgerHandler) GetUserLedger(c *gin.Context) {
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		h.log.WithField("user_id_param", userIDStr).Warn("Invalid userId parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}

	entries, err := h.ledgerService.GetUserEntries(c.Request.Context(), userID)
	if err != nil {
		h.log.WithField("user_id", userID).Errorf("Failed to fetch ledger: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ledger"})
		return
	}

	h.log.WithField("user_id", userID).Infof("Fetched ledger successfully, entries=%d", len(entries))

	c.JSON(http.StatusOK, gin.H{
		"userId": userID,
		"ledger": entries,
	})
}
