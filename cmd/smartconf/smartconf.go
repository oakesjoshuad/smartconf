package main

import (
	"flag"
	"fmt"
	//	"io/ioutil"

	"github.com/oakesjoshuad/smartconf"
)

const (
	ALIAS         string = "ALIAS"
	HOST          string = "HOSTNAME"
	MEMORY        int    = 512
	BRAND         string = "joyent"
	IMAGE         string = "800db35c-5408-11eb-9792-872f658e7911"
	INTERFACE     string = "net0"
	GATEWAY       string = "10.0.0.253"
	NETMASK       string = "255.255.255.0"
	NICTAG        string = "extern"
	DNSDOMAIN     string = "local.net"
	PUBLICKEY     string = "PUBLICKEY"
	ROOTAUTHKEY   string = "root_authorized_keys"
	USERSCRIPTKEY string = "user-script"
	USERSCRIPT    string = "/usr/bin/mdata-get root_authorized_keys > ~root/.ssh/authorized_keys"
	LIMITPRIV     string = "default"
)

var (
	RESOLVERS []string = []string{"10.0.0.250"}
)

func main() {
	var autoboot bool
	flag.BoolVar(&autoboot, "autoboot", true, "automatically boot virtual machine on system startup")
	var brand string
	flag.StringVar(&brand, "b", BRAND, "virtual machine brand: (joyent|lx|hvm)")
	flag.StringVar(&brand, "brand", BRAND, "")
	var memory int
	flag.IntVar(&memory, "m", MEMORY, "max physical memory")
	flag.IntVar(&memory, "memory", MEMORY, "")
	var hostname string
	flag.StringVar(&hostname, "h", HOST, "vm hostname")
	flag.StringVar(&hostname, "host", HOST, "")
	var alias string
	flag.StringVar(&alias, "a", ALIAS, "vm alias")
	flag.StringVar(&alias, "alias", ALIAS, "")
	var nictag string
	flag.StringVar(&nictag, "n", NICTAG, "global interface")
	flag.StringVar(&nictag, "nictag", NICTAG, "")
	var dnsdomain string
	flag.StringVar(&dnsdomain, "d", DNSDOMAIN, "dns domain name")
	flag.StringVar(&dnsdomain, "dns", DNSDOMAIN, "")
	var image string
	flag.StringVar(&image, "i", IMAGE, "virtual machine image uuid")
	flag.StringVar(&image, "image", IMAGE, "")
	var pubkey string
	flag.StringVar(&pubkey, "pk", PUBLICKEY, "ssh public key file")
	flag.StringVar(&pubkey, "publickey", PUBLICKEY, "")
	var firewall bool
	flag.BoolVar(&firewall, "fw", false, "enable/disable global firewall participation")
	flag.BoolVar(&firewall, "firewall", false, "")
	var limitpriv string
	flag.StringVar(&limitpriv, "l", LIMITPRIV, "default privileges")
	flag.StringVar(&limitpriv, "limit", LIMITPRIV, "")

	flag.Parse()
	var zonecfg = smartconf.ZoneConfig{
		Autoboot:          autoboot,
		Brand:             brand,
		ImageUUID:         image,
		LimitPriv:         limitpriv,
		MaxPhysicalMemory: memory,
		Hostname:          hostname,
		Alias:             alias,
		Resolvers:         RESOLVERS,
		Nics: []smartconf.Nic{
			{
				InterfaceName: INTERFACE,
				NicTag:        nictag,
				Ips:           []string{"dhcp", "addrconf"},
				Gateways:      []string{GATEWAY},
				Netmask:       NETMASK,
				Primary:       true,
			},
		},
		DNSDomain:       dnsdomain,
		FirewallEnabled: firewall,
		CustomerMetadata: map[string]string{
			ROOTAUTHKEY:   PUBLICKEY,
			USERSCRIPTKEY: USERSCRIPT,
		},
	}
	payload, err := zonecfg.JSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(payload))
}
