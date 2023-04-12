#!/bin/bash

# Java11
sudo add-apt-repository ppa:openjdk-r/ppa
sudo apt-get update -q

sudo apt install -y openjdk-11-jdk
# sudo yum install java-11-openjdk

sudo alternatives --config java
java -version
