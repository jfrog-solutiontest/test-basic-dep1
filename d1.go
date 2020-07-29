package d1

import (
	"fmt"
	"github.com/jfrog-solutiontest/test-basic-dep2"
	"github.com/miekg/dns"
)

func PackageName () {
	fmt.Println ("Package Name: d1");
	test-basic-dep2.PackageName()
}

func updateDomainWithCName(r *dns.Msg, fqdn string) string {
	for _, rr := range r.Answer {
		if cn, ok := rr.(*dns.CNAME); ok {
			if cn.Hdr.Name == fqdn {
				return cn.Target
			}
		}
	}

	return fqdn
}
