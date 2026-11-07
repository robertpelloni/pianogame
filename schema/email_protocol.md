# External Email & Notification Dispatch Protocol

## Overview
Defines communication protocols for reading inbound developer alerts, task assignments, and sending system notifications across external notification transports (SMTP, Webhook, SMS).

## Schema

### 1. Inbound Email Message
```json
{
  "id": "msg_001",
  "from": "developer@robertpelloni.com",
  "to": "agent@harness.internal",
  "subject": "[Task] Optimize Multi-Agent Racing latency",
  "body": "Please review AST parsing benchmarks and optimize tree traversals.",
  "timestamp": "2026-11-06T12:00:00Z"
}
```

### 2. Outbound Dispatch Notification
```json
{
  "recipient": "admin@robertpelloni.com",
  "level": "info",
  "title": "Harness Build v5.85.0 Succeeded",
  "payload": {
    "version": "v5.85.0",
    "passedEngines": 5
  }
}
```
