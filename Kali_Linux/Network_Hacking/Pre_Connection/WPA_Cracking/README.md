# WPA/WPA2 Cracking

## Method 1: WPS (This only works if push-only auth disabled)

* WPS is a feature that can be used with WPA, WPA2. 
* It allows clients to connect without password, but only 8 digit PIN
* Then 8 digit PIN can be used to compute real passwords

WPS requies an 8bit PIN, which is easy to guess. But if PBC enabled, then we have to physically press button on the router, which is impossible. 

`wash --interface wlan0` to check if any WPS around 

## Method 2: If WPA/WPA2 does not have WPS, or PBC enabled

* WPA/WPA2 are very strong
* Packets itself has no useful information like WEP
* ONLY `HANDSHAKE` packets are important 

Then use `crunch wordlist` to crack the password

Packet has an MIC which is the result of the real handshake and real password. As long as our wordlist password can have the same result after encoding with the handshake. Then we crack the password.

`Handshake + Password = MIC`
`If Handshake + Crunched Wordlist = MIC`
`Then Crunched Wordlist = Password`

`airodump-ng --bssid xxx --channel xxx --write some_file`: Put captured packets into a file
`aireplay-ng --deauth 4 -a TARGET_MAC -c CLIENT_MAC wlan0`: Send some packets to cut the connection, so handshake packet will be sent
`aircrack-ng some_file.cap -w word_list`: Use wordlist to crack WPA using brutal force 
