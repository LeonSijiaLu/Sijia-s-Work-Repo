#!/bin/bash

wireless_adapter=$1

echo "Shutting down wireless adapter ... ..."
echo
ifconfig $wireless_adapter down

echo "Airmon-ng kill conflicting processes ... ..."
echo
airmon-ng check kill

echo "Turning $wireless_adapter into monitor mode ... ..."
echo 
iwconfig $wireless_adapter mode monitor

echo "Starting up $wireless_adapter ... ..."
echo
ifconfig $wireless_adapter up
 
echo "Done !"
echo
