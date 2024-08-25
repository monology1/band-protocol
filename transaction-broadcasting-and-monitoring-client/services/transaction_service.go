package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/band-protocol/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Context interface {
	QueryParam(input string) string
	Json(status int, input interface{})
	Bind(interface{}) error
}

func BroadcastTransaction(c Context) (string, error) {
	var req models.BroadcastRequest
	timeStamp := time.Now().Unix()
	err := c.Bind(&req)
	if err != nil {
		c.Json(400, gin.H{"message": err.Error()})
	}
	txHash, err := broadcastTransaction(req.Symbol, req.Price, timeStamp)
	if err != nil {
		c.Json(500, gin.H{"message": err.Error()})
	}
	c.Json(200, gin.H{"txHash": txHash})
	return txHash, nil
}

func MonitorTransaction(txHash string) (string, error) {
	for {
		status, err := getTransactionStatus(txHash)
		if err != nil {
			return "", err
		}

		fmt.Println("Transaction status: %s\n", status)

		// If transaction confirmed, failed, done just stop monitoring
		if status == "CONFIRMED" || status == "FAILED" || status == "DNE" {
			return status, nil
		}
		// If transaction is pending, wait before checking again
		if status == "PENDING" {
			time.Sleep(10 * time.Second) // Wait 10 seconds before re checking
		}
	}
}

func BroadcastAndMonitorTransaction(c Context) {
	// Broadcast Transaction
	txHash, err := BroadcastTransaction(c)
	if err != nil {
		c.Json(500, gin.H{"message": "Failed to broadcast transaction"})
	}

	// Monitor the transaction status
	status, err := MonitorTransaction(txHash)
	if err != nil {
		c.Json(500, gin.H{"message": "Failed to monitor transaction"})
	}

	// Response with status
	c.Json(200, gin.H{"status": status})
}

func broadcastTransaction(symbol string, price uint64, timeStamp int64) (string, error) {
	payload := models.BroadcastRequest{
		Symbol:    symbol,
		Price:     price,
		Timestamp: timeStamp,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	// Create a new http client
	url := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	client := &http.Client{}
	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to broadcast transaction")
	}

	//Response body
	var response models.BroadcastResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.TxHash, nil
}

func getTransactionStatus(txHash string) (string, error) {
	url := fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%s", txHash)
	// Create a new http client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get transaction status")
	}
	var statusResponse models.TransactionStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&statusResponse)
	if err != nil {
		return "", err
	}
	return statusResponse.TxStatus, nil
}
