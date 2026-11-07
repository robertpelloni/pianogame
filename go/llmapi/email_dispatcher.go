package llmapi

import (
	"fmt"
	"time"
)

// InboundEmail represents an incoming developer task message
type InboundEmail struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

// OutboundDispatch represents an outbound alert notification
type OutboundDispatch struct {
	Recipient string    `json:"recipient"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// EmailDispatcher manages external message parsing and alert transmissions
type EmailDispatcher struct {
	dispatches []OutboundDispatch
}

// NewEmailDispatcher constructs a new dispatcher
func NewEmailDispatcher() *EmailDispatcher {
	return &EmailDispatcher{}
}

// ParseInbound converts incoming raw text to a structured message
func (e *EmailDispatcher) ParseInbound(raw string) InboundEmail {
	return InboundEmail{
		ID:        fmt.Sprintf("email_%d", time.Now().UnixMilli()),
		From:      "supervisor@robertpelloni.com",
		To:        "agent@harness.internal",
		Subject:   "Autonomous task execution update",
		Body:      raw,
		Timestamp: time.Now(),
	}
}

// DispatchNotification sends an external system alert
func (e *EmailDispatcher) DispatchNotification(recipient, subject, message string) OutboundDispatch {
	item := OutboundDispatch{
		Recipient: recipient,
		Subject:   subject,
		Message:   message,
		Status:    "sent",
		Timestamp: time.Now(),
	}
	e.dispatches = append(e.dispatches, item)
	return item
}

// GetHistory returns dispatched notification history
func (e *EmailDispatcher) GetHistory() []OutboundDispatch {
	return e.dispatches
}
