import { w3cwebsocket as W3CWebSocket } from 'websocket';
export const SocketClient = new W3CWebSocket('ws://localhost:8080/ws');