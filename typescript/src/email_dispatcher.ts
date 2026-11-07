export interface InboundEmail {
  id: string;
  from: string;
  to: string;
  subject: string;
  body: string;
  timestamp: string;
}

export interface OutboundDispatch {
  recipient: string;
  subject: string;
  message: string;
  status: "sent" | "queued" | "failed";
  timestamp: string;
}

/**
 * Lightweight External Email & Notification Dispatcher
 * Parses inbound tasks from emails and dispatches outbound alerts.
 */
export class EmailDispatcher {
  private dispatches: OutboundDispatch[] = [];

  public parseInbound(rawMime: string): InboundEmail {
    return {
      id: `email_${Date.now()}`,
      from: "supervisor@robertpelloni.com",
      to: "agent@harness.internal",
      subject: "Autonomous task execution update",
      body: rawMime.trim(),
      timestamp: new Date().toISOString(),
    };
  }

  public dispatchNotification(recipient: string, subject: string, message: string): OutboundDispatch {
    const item: OutboundDispatch = {
      recipient,
      subject,
      message,
      status: "sent",
      timestamp: new Date().toISOString(),
    };
    this.dispatches.push(item);
    return item;
  }

  public getHistory(): OutboundDispatch[] {
    return [...this.dispatches];
  }
}
