FROM sangyunhan/ubuntu-for-network-test
MAINTAINER Sangyun Han <sangyun628@gmail.com>

# run update
RUN apt-get update

RUN apt-get install strongswan -y
RUN apt-get install git -y

# project build setting
WORKDIR /root
RUN git clone https://github.com/sangyun-han/sidekick
RUN git clone https://github.com/sangyun-han/strongswan-go
#ADD sidekick/vpn/strongswan.py /root/
