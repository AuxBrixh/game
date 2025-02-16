package server

import (
	// "context"
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/codecat/go-enet"
)

type Server struct {
	conf ServerConfig

	once    sync.Once
	listeners []Listener
	started   atomic.Pointer[time.Time]
	wg        sync.WaitGroup
}

func (srv *Server) Listen() {
	t := time.Now()
	if !srv.started.CompareAndSwap(nil, &t) {
		panic("start server: already started")
	}

	info, _ := debug.ReadBuildInfo()
	if info == nil {
		info = &debug.BuildInfo{GoVersion: "N/A", Settings: []debug.BuildSetting{{Key: "vcs.revision", Value: "N/A"}}}
	}

	srv.conf.Log.Info("Server started.", "go-version", info.GoVersion)
	srv.startListening()
	go srv.wait()
}

func(srv *Server) startgame(l Listener) {
	
}

func(srv *Server) wait() {
	srv.wg.Wait()
}

func (srv *Server) startListening() {
	srv.wg.Add(len(srv.listeners))
	for _, l := range srv.listeners {
		go srv.listen(l)
	}
}

func (srv *Server) listen(l Listener) {
	wg := new(sync.WaitGroup)
	// ctx, _ := context.WithCancel(context.Background())

	for {
		ev := l.Service(100)

		switch ev.GetType() {
		case enet.EventConnect:
			peer := ev.GetPeer()
			ev.GetPeer().SendBytes([]byte("DATA_AUTH"), ev.GetChannelID(), enet.PacketFlagReliable)
			fmt.Println("New peer connected:", peer.GetAddress(), peer.GetAddress().GetPort(), peer.GetData())
			wg.Add(1)
		case enet.EventReceive:
			packet := ev.GetPacket()

			defer packet.Destroy()

			packetBytes := packet.GetData()
			fmt.Printf("data yang di terima %s", string(packetBytes))
		}
	}
}
