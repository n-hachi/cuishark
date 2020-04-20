package packet

import (
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func GetPacket() (p gopacket.Packet, err error) {
	handle, err := pcap.OpenOffline("test_ethernet.pcap")
	if err != nil {
		return nil, err
	}
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	p, err = source.NextPacket()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func TestNewPacket(t *testing.T) {
	gp, err := GetPacket()
	if err != nil {
		t.Errorf("err is %v\n", err)
	} else if gp == nil {
		t.Error("p should not be nil\n")
	}
	p := NewPacket(gp)
	if err != nil {
		t.Errorf("err is %v\n", err)
	} else if p == nil {
		t.Error("p should not be nil\n")
	}
}

func TestPacket_Size(t *testing.T) {
	gp, _ := GetPacket()
	p := NewPacket(gp)
	if s := p.Size(); s != 3 {
		t.Errorf("size should be 3, but actually size is %d\n", s)
	} else {
		t.Log("size value is 3 as expected\n")
	}
}
