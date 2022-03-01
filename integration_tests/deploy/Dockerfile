# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

################################################################################
# Image of a clean machine to deploy the new node onto.
################################################################################
FROM jrei/systemd-ubuntu:20.04

RUN apt-get update && apt-get install -y \
        sudo \
        curl \
        netcat \
        iproute2 \
    && rm -rf /var/lib/apt/lists/*

# test user
ARG TEST_USER
ENV TEST_USER=${TEST_USER:-dcl}

#ARG TEST_USER_GROUP
#ENV TEST_USER_GROUP=${TEST_USER_GROUP:-dcl}

ARG TEST_UID
ENV TEST_UID=${TEST_UID:-1000}
#ARG gid=1000
RUN adduser --disabled-password -uid ${TEST_UID} --home /var/lib/${TEST_USER} --gecos 'DCLedger user' ${TEST_USER}

RUN usermod -aG sudo dcl \
    && echo "dcl     ALL=(ALL:ALL) NOPASSWD:ALL" >>/etc/sudoers

EXPOSE 26656 26657 1317 26660

# USER ${TEST_USER}
WORKDIR /var/lib/${TEST_USER}
