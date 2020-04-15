package handler

import (
	"github.com/google/gopacket/pcap"
)

type PcapHandler struct {
	handle *pcap.Handle
}

func NewPcapHandler(path string) (handler *PcapHandler, err error) {
	handle, err := pcap.OpenOffline(path)
	if err != nil {
		return nil, err
	}
	handler = new(PcapHandler)
	handler.handle = handle

	return handler, nil
}
