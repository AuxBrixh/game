package server

import (
	"fmt"

	"github.com/codecat/go-enet"
)

type Listener interface {
	Destroy()
	Service(timeout uint32) enet.Event

	Connect(addr enet.Address, channelCount int, data uint32) (enet.Peer, error)

	CompressWithRangeCoder() error
	BroadcastBytes(data []byte, channel uint8, flags enet.PacketFlags) error
	BroadcastPacket(packet enet.Packet, channel uint8) error
	BroadcastString(str string, channel uint8, flags enet.PacketFlags) error
}

func(uc *UserConfig) Listeners(conf ServerConfig) (Listener, error) {
	l, err := EnetStart(uc.Server.Port)
	if err != nil {
		return nil, err
	}
	conf.Log.Info("Listener running.", "addr", uc.Server.Port)
	return l, nil
}

func EnetStart(port uint16) (Listener, error) {
	enet.Initialize()
	host, err := enet.NewHost(enet.NewListenAddress(port), 32, 1, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create host: %s", err.Error())
	}

	return host, nil
}