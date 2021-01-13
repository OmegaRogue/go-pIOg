package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

var bug = [8]byte{
	0x00, 0x0A, 0x04, 0x1F, 0x0E, 0x1F, 0x0E, 0x11,
}

var notice = [8]byte{
	0x04, 0x0E, 0x0E, 0x0E, 0x1F, 0x00, 0x04, 0x00,
}

var alert = [8]byte{
	0x00, 0x1B, 0x0E, 0x04, 0x0E, 0x1B, 0x00, 0x00,
}

var emergency = [8]byte{
	0x00, 0x00, 0x0A, 0x00, 0x0E, 0x11, 0x00, 0x00,
}

var severities = [8]string{
	fmt.Sprintf("%v", emergency),
	fmt.Sprintf("%v", alert),
	"X",
	"x",
	"!",
	fmt.Sprintf("%v", notice),
	"i",
	fmt.Sprintf("%v", bug),
}

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Lookup a pin by its number:
	p := gpioreg.ByName("GPIO17")
	if p == nil {
		log.Fatal("Failed to find GPIO17")
	}

	fmt.Printf("%s: %s\n", p, p.Function())

	// Set it as input, with an internal pull down resistor:
	if err := p.In(gpio.PullUp, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}

	// Wait for edges as detected by the hardware, and print the value read:
	for {
		p.WaitForEdge(-1)
		fmt.Printf("-> %s\n", p.Read())
	}

}

/*

   Numerical             Facility
             Code

              0             kernel messages
              1             user-level messages
              2             mail system
              3             system daemons
              4             security/authorization messages
              5             messages generated internally by syslogd
              6             line printer subsystem
              7             network news subsystem
              8             UUCP subsystem
              9             clock daemon
             10             security/authorization messages
             11             FTP daemon
             12             NTP subsystem
             13             log audit
             14             log alert
             15             clock daemon (note 2)
             16             local use 0  (local0)
             17             local use 1  (local1)
             18             local use 2  (local2)
             19             local use 3  (local3)
             20             local use 4  (local4)
             21             local use 5  (local5)
             22             local use 6  (local6)
             23             local use 7  (local7)

Numerical         Severity
             Code

              0       Emergency: system is unusable
              1       Alert: action must be taken immediately
              2       Critical: critical conditions
              3       Error: error conditions
              4       Warning: warning conditions
              5       Notice: normal but significant condition
              6       Informational: informational messages
              7       Debug: debug-level messages

*/
