package packet

import (
	"fmt"
	"strings"
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

func TestPacket_LayerType(t *testing.T) {
	gp, _ := GetPacket()
	p := NewPacket(gp)
	if fmt.Sprintf("%v", p.LayerType()) != "TCP" {
		t.Errorf("p.LayerType() should be TCP, but actually %v\n", p.LayerType())
	} else {
		t.Log("p.LayerType() is TCP as expected\n")
	}

}
func TestPacket_Src(t *testing.T) {
	gp, _ := GetPacket()
	p := NewPacket(gp)
	if fmt.Sprintf("%v", p.Src()) != "10.1.1.2" {
		t.Errorf("p.Src() should be 10.1.1.2, but actually %v\n", p.Src())
	} else {
		t.Log("p.Src() is 10.1.1.2 as expected\n")
	}
}

func TestPacket_Dst(t *testing.T) {
	gp, _ := GetPacket()
	p := NewPacket(gp)
	if fmt.Sprintf("%v", p.Dst()) != "10.1.1.1" {
		t.Errorf("p.Dst() should be 10.1.1.1, but actually %v\n", p.Dst())
	} else {
		t.Log("p.Dst() is 10.1.1.1 as expected\n")
	}
}

func TestPacket_Oneline(t *testing.T) {
	gp, _ := GetPacket()
	p := NewPacket(gp)
	if !strings.HasPrefix(p.Oneline(), "SrcPort=44644") {
		t.Errorf("p.Oneline() should be start with \"SrcPort=44644\", but actually \"%s\"\n", p.Oneline())
	} else {
		t.Logf("p.Oneline() is %s\n", p.Oneline())
	}
}
