# NMap

1. SYN scan

`-sS`: `nmap -sS 10.31.0.10`. It only sends a SYN packet to the target machine WITHOUT any useful information. So `no trace`

 Pretty much, it gives a brief view of open ports

`nmap -sS 10.31.0.10 | grep open`

2. Connect scan

`-sT`: It is TCP scan. 

3. Version scan

`-sV`: It gives versions of ports

4. Brief scan

`-F`: simplest scan

## Example

`sudo nmap -sV -O -T4 10.31.0.180`: is an example that lists detailed infor about an IP
