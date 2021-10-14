package main

type network struct {
	base        string    // 10.0.0.0
	cidr        int       // ip/cidr [10.0.0.0/24]
	vlanid      int       // 0-4094
	services    []service // an array of network services i.e. kerberos, ldap, dhcp, dns, etc...
	description string
}

type service struct {
	alias       string
	hostname    string
	description string
	image       string
	autoboot    bool
	memory      int
	publickey   string
}

type site map[string]network

func (s site) generate() error {
	for desc, network := range s {
		fmt.Prinln(desc)
	}
	return nil
}

const (
	minimal int = iota
	base
	build
)

var images = map[int]string{
	minimal: "",
	base:    "",
	build:   "",
}

var vlanids = map[string]int{
	"admin":   10,
	"netsvcs": 100,
	"websvcs": 200,
}

var localio = site{
	"netsvcs": {
		description: "network services, authentication, dhcp, dns, etc...",
		services: []service{
			{
				description: "kerberos v5",
				alias:       "krb",
				hostname:    "krb",
				image:       images[minimal],
				memory:      1024,
			},
			{
				description: "dhcp provider",
				alias:       "dhcp",
				hostname:    "dhcp",
				image:       images[minimal],
				memory:      512,
			},
			{
				description: "ldap provider",
				alias:       "ldap",
				hostname:    "ldap",
				image:       images[minimal],
				memory:      1024,
			},
			{
				description: "radius provider",
				alias:       "rad",
				hostname:    "rad",
				image:       images[minimal],
				memory:      512,
			},
			{
				description: "unbound caching dns server",
				alias:       "unbound",
				hostname:    "dns",
				image:       images[minimal],
				memory:      512,
			},
			{
				description: "certificate authority",
				alias:       "vault",
				hostname:    "vault",
				autoboot:    false,
				image:       images[minimal],
				memory:      2048,
			},
			{
				description: "intermediate certificat authority",
				alias:       "secure",
				hostname:    "secure",
				memory:      2048,
			},
		},
	},
}
