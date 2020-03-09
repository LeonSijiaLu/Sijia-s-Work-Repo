# Deauthentication Attack

* It only disconnect clients from AP
* SO they need to connect again
* Then we can capture useful information like `HANDSHAKE`
* `HANDSHAKE` contains no password, it can only check if the password is correct or not

`aireplay-ng --deauth -a AP_MAC -c Client_MAC wlan0` It stops Client_MAC from accessing AP_MAC
