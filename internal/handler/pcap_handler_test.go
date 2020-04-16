package handler

import "testing"

func TestNewPcapHanddlerWithNonexistenceFile(t *testing.T) {
	handler, err := NewPcapHandler("path/to/noexistent/file")
	if err == nil {
		t.Error("No error returned for nonexistent file open")
	} else {
		t.Logf("Error return for nonexistent file: %v", err)
	}
	if handler != nil {
		t.Error("Non-nil handler returned for nonexistent file open")
	}
}

func TestNewPcapHanddlerWithExistenceFile(t *testing.T) {
	handler, err := NewPcapHandler("test_ethernet.pcap")
	if err != nil {
		t.Errorf("An error returned for existent file open: %v", err)
	} else {
		t.Log("No error return for existent file")
	}
	if handler == nil {
		t.Error("Nil handler returned for existent file open")
	}
}

func TestOpenChan(t *testing.T) {
	handler, _ := NewPcapHandler("test_ethernet.pcap")
	ch := handler.OpenChan()
	_, ok := <-ch
	if ok == false {
		t.Error("Channel should have some packet informations.")
	}
}
