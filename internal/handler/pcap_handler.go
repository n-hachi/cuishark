package handler

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type PcapHandler struct {
	handle *pcap.Handle
	source *gopacket.PacketSource
}

func NewPcapHandler(path string) (handler *PcapHandler, err error) {
	handle, err := pcap.OpenOffline(path)
	if err != nil {
		return nil, err
	}
	handler = new(PcapHandler)
	handler.handle = handle
	handler.source = gopacket.NewPacketSource(handle, handle.LinkType())

	return handler, nil
}

func (pc *PcapHandler) OpenChan() chan gopacket.Packet {
	return pc.source.Packets()
}
