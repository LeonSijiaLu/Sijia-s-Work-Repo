#!/bin/bash

echo "deb https://mirrors.tuna.tsinghua.edu.cn/kali kali-rolling main contrib non-free" >> /etc/apt/sources.list
echo "deb-src https://mirrors.tuna.tsinghua.edu.cn/kali kali-rolling main contrib non-free" >> /etc/apt/sources.list

sudo apt-get clean
sudo apt-get update -y
sudo apt-get install ttf-wqy-microhei ttf-wqy-zenhei xfonts-intl-chinese

# dpkg-reconfigure locales
