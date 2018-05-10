FROM sangyunhan/ubuntu-for-network-test
MAINTAINER Sangyun Han <sangyun628@gmail.com>

# run update
RUN apt-get update

RUN apt-get install strongswan -y
RUN apt-get install strongswan-swanctl

