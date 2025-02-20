WebRTC (Web Real-Time Communication) is a technology that enables peer-to-peer communication directly in the browser. It allows for real-time audio, video, and data transfer without the need for plugins or additional software.

### Key Components of WebRTC:
#### Signaling Server:
- WebRTC itself does not define a signaling protocol. Signaling is used to exchange metadata between peers to establish a connection. This metadata includes SDP (Session Description Protocol) and ICE candidates.
- Signaling can be done using WebSockets, HTTP, or any other communication protocol.

#### SDP (Session Description Protocol):
- SDP is a format for describing multimedia communication sessions. It includes information about the media streams, codecs, and network information.

#### ICE (Interactive Connectivity Establishment):
- ICE is a framework used to establish a connection between peers. It involves gathering candidate IP addresses and ports (using STUN/TURN servers) and then using these candidates to find the best path for communication.

#### STUN/TURN Servers:
- STUN (Session Traversal Utilities for NAT) servers help peers discover their public IP addresses and ports.
- TURN (Traversal Using Relays around NAT) servers act as relays when direct peer-to-peer communication is not possible due to NAT or firewall restrictions.
