#!/bin/bash

# Copyright 2014 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

REGION=us-east-1

INSTANCETYPE=m3.medium
NUM_MINIONS=1

CLUSTER_KEY=kubernetes

# Ubuntu AMI
AMI=ami-9eaa1cf6
SSH_USER=ubuntu

# AWS AMI... doesn't work - sudo needs a tty?
#AMI=ami-0268d56a
#SSH_USER=ec2-user

#MINION_NAMES=($(eval echo ${INSTANCE_PREFIX}-minion-{1..${NUM_MINIONS}}))
#MINION_IP_RANGES=($(eval echo "10.244.{1..${NUM_MINIONS}}.0/24"))
#MINION_SCOPES="compute-rw"
# Increase the sleep interval value if concerned about API rate limits. 3, in seconds, is the default.
#POLL_SLEEP_INTERVAL=3

# Once we have IPv6 merged...
  # This is the 0a64 => 10.100.0.0
  #  10.100 is hard coded in the VPC setup
  #PORTAL_NET="2002:0a64:0000::/48"
# Until IPv6...
PORTAL_NET="10.100.0.0/16"
