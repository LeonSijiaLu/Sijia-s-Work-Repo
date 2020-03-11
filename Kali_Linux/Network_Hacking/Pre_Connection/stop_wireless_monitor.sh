#!/bin/bash

wireless_adapter=$1

echo "Shutting down wireless adapter ... ..."
echo
ifconfig $wireless_adapter down

echo "Airmon-ng kill conflicting processes ... ..."
echo
airmon-ng check kill

echo "Turning $wireless_adapter into managed mode ... ..."
echo 
iwconfig $wireless_adapter mode managed

echo "Starting up $wireless_adapter ... ..."
echo
ifconfig $wireless_adapter up

echo "Restarting NetworkManager ... ..."
echo 
systemctl restart NetworkManager

echo "Done !"
echo
