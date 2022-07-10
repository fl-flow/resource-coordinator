package etc

import (
  "log"
  "fmt"
  "flag"
)


func init() {
  // server
  ip := flag.String("ip", IP, "ip")
	port := flag.Int("port", PORT, "port")

	flag.Parse()

  IP = *ip
  PORT = *port

  log.Println(fmt.Sprintf(
      `
      // server
      IP: %v
      PORT: %v
      `,
      // server
      IP,
      PORT,
  ),)
}
