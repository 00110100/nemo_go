package portscan

import (
	"testing"
)

func TestNmap_Run(t *testing.T) {
	config := Config{
		Target:        "192.168.3.0/24",
		ExcludeTarget: "",
		Port:          "--top-ports 1000",
		Rate:          1000,
		IsPing:        true,
		Tech:          "-sV",
		CmdBin:        "nmap",
	}
	nmap := NewNmap(config)
	nmap.Do()
	nmap.Result.SaveResult(nmap.Config)

	t.Log(nmap.Result)
	for ip,ipa := range nmap.Result.IPResult{
		t.Log(ip,ipa)
		for port,pa := range ipa.Ports{
			t.Log(port,pa)
		}
	}
}
