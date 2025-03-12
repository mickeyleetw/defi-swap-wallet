package handlers

import (
	"defi-swap-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

func SendTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ts := services.NewTransactionService()
	txHash, err := ts.SendTransaction(req.To, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
}
