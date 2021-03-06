package dns_test

import (
	"fmt"
	"log"

	"github.com/shuLhan/share/lib/dns"
	libnet "github.com/shuLhan/share/lib/net"
)

//
// The following example show how to use send and Recv to query domain name
// address.
//
func ExampleUDPClient() {
	cl, err := dns.NewUDPClient("127.0.0.1:53")
	if err != nil {
		log.Println(err)
		return
	}

	ns, err := libnet.ParseUDPAddr("127.0.0.1", 53)
	if err != nil {
		log.Println(err)
		return
	}

	req := &dns.Message{
		Header: &dns.SectionHeader{},
		Question: &dns.SectionQuestion{
			Name:  []byte("kilabit.info"),
			Type:  dns.QueryTypeA,
			Class: dns.QueryClassIN,
		},
	}

	_, err = req.Pack()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = cl.Send(req, ns)
	if err != nil {
		log.Println(err)
		return
	}

	res := dns.NewMessage()

	_, err = cl.Recv(res)
	if err != nil {
		log.Println(err)
		return
	}

	res.Unpack()

	fmt.Printf("Receiving DNS message: %s\n", res)
	for x, answer := range res.Answer {
		fmt.Printf("Answer %d: %s\n", x, answer.RData())
	}
	for x, auth := range res.Authority {
		fmt.Printf("Authority %d: %s\n", x, auth.RData())
	}
	for x, add := range res.Additional {
		fmt.Printf("Additional %d: %s\n", x, add.RData())
	}
}

func ExampleUDPClient_Lookup() {
	cl, err := dns.NewUDPClient("127.0.0.1:53")
	if err != nil {
		log.Println(err)
		return
	}

	msg, err := cl.Lookup(dns.QueryTypeA, dns.QueryClassIN, []byte("kilabit.info"))
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Receiving DNS message: %s\n", msg)
	for x, answer := range msg.Answer {
		fmt.Printf("Answer %d: %s\n", x, answer.RData())
	}
	for x, auth := range msg.Authority {
		fmt.Printf("Authority %d: %s\n", x, auth.RData())
	}
	for x, add := range msg.Additional {
		fmt.Printf("Additional %d: %s\n", x, add.RData())
	}
}
