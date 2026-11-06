import http from "http";
import crypto from "crypto";

export interface WsClient {
  id: string;
  socket: any;
  alive: boolean;
}

/**
 * Lightweight native WebSocket Bridge for the Ultimate Agentic Coding Harness.
 * Allows remote browser/CLI agents to stream control messages and AST analysis results
 * without introducing bulky external runtime dependencies.
 */
export class WsBridge {
  private clients: Map<string, any> = new Map();

  public handleUpgrade(req: http.IncomingMessage, socket: any, head: Buffer) {
    const key = req.headers['sec-websocket-key'];
    if (!key) {
      socket.destroy();
      return;
    }

    const acceptKey = crypto
      .createHash('sha1')
      .update(key + '258EAFA5-E914-47DA-95CA-C5AB0DC85B11')
      .digest('base64');

    const headers = [
      'HTTP/1.1 101 Switching Protocols',
      'Upgrade: websocket',
      'Connection: Upgrade',
      `Sec-WebSocket-Accept: ${acceptKey}`
    ];

    socket.write(headers.join('\r\n') + '\r\n\r\n');

    const clientId = crypto.randomUUID();
    this.clients.set(clientId, socket);

    socket.on('close', () => {
      this.clients.delete(clientId);
    });

    socket.on('error', () => {
      this.clients.delete(clientId);
    });
  }

  public getConnectedClientsCount(): number {
    return this.clients.size;
  }
}
