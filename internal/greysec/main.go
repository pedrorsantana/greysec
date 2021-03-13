package main

import (
	"fmt"
	"log"

	"github.com/pedrorsantana/greysec/internal/greynoise"
	"github.com/pedrorsantana/greysec/internal/iptable"
	"github.com/pedrorsantana/greysec/internal/pkCap"
)

func main() {
	grey := greynoise.New("32yjv1OZOFPq6cYCeOVNgJsVT7w7ESxWCjealRHWQTYrwf4QdahdsAtk7huBNc9B", 120)
	pkCap.New(func(header []uint8) {
		if header != nil {
			ip := fmt.Sprintf("%d.%d.%d.%d", header[0], header[1], header[2], header[3])
			if found, _ := grey.LookUpIP(ip); found == true {
				if blocked, _ := iptable.BlockIP(ip); blocked == true {
					log.Printf("The ip %s has been blocked because they are be founded in greynoise.io database.", ip)
				}
			}
		}
	})

}
