# Plasma
Minecraft server framework written in Go. Having an API similar to the standard http package,
this framework allows you to write Minecraft servers from scratch without needing too much
boilerplate.

## Basic Server
```go
package main

import (
	"github.com/specspace/plasma"
	"github.com/specspace/plasma/protocol"
	"github.com/specspace/plasma/protocol/packets/handshaking"
	"github.com/specspace/plasma/protocol/packets/status"
	"log"
)

func main() {
	plasma.HandleFunc(protocol.StateHandshaking, handshaking.ServerBoundHandshakePacketID, handshakeHandler)
	plasma.HandleFunc(protocol.StateStatus, status.ServerBoundPingPacketID, pingHandler)
	plasma.HandleFunc(protocol.StateStatus, status.ServerBoundRequestPacketID, responseHandler)

	log.Fatal(plasma.ListenAndServe(":25565", nil))
}

func handshakeHandler(w plasma.ResponseWriter, r *plasma.Request) {
	// You can access and unmarshal the handshake packet like this.
	hs, err := handshaking.UnmarshalServerBoundHandshake(r.Packet)
	if err != nil {
		return
	}

	if hs.IsStatusRequest() {
		// We can now update the connection state according to the
		// handshake request packet.
		w.SetState(protocol.StateStatus)
	} else if hs.IsLoginRequest() {
		w.SetState(protocol.StateLogin)
	}
}

func pingHandler(w plasma.ResponseWriter, r *plasma.Request) {
	ping, err := status.UnmarshalServerBoundPing(r.Packet)
	if err != nil {
		return
	}

	pong := status.ClientBoundPong{
		Payload: ping.Payload,
	}

	w.WritePacket(pong.Marshal())
}

func responseHandler(w plasma.ResponseWriter, r *plasma.Request) {
	statusResponse := plasma.StatusResponse{
		Version: plasma.Version{
			Name:           "Plasma 1.16.4",
			ProtocolNumber: 754,
		},
		PlayersInfo: r.Server.PlayersInfo(),
		IconPath:    "",
		MOTD:        "Hello World",
	}

	bb, err := statusResponse.JSON()
	if err != nil {
		return
	}

	w.WritePacket(status.ClientBoundResponse{
		JSONResponse: protocol.String(bb),
	}.Marshal())
}
```