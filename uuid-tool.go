package main

import (
	"fmt"
	"flag"
	"code.google.com/p/go-uuid/uuid"
)

func main() {
	flag.Parse()
	args := flag.Args()

	theid := uuid.Parse(args[0])
	fmt.Printf("STRING=%s\n", theid.String())
	fmt.Printf("VARIANT=%s\n", theid.Variant())
	version, valid := theid.Version()
	fmt.Printf("VERSION=%s\n", version)
	fmt.Printf("VALID=%t\n", valid)
	b := []byte(theid)
	fmt.Printf("DCE_FORMAT={0x%08x,0x%04x,0x%04x,0x%01x,0x%01x,{0x%01x,0x%01x,0x%01x,0x%01x,0x%01x,0x%01x}}\n",
		b[:4], b[4:6], b[6:8], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
	return
}
