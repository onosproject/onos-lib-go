<!--
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Golang SCTP Library 

## Overview

SCTP library provides the following features and functionalities:

* one-to-one and one-to-many SCTP interface models
  
* Supports SCTP event notification and provides abstractions to determine the 
  type of  event. SCTP suggests two ways to enable the notifications using 
  SCTP_EVENT and SCTP_EVENTS socket options. At the moment, in client side we support SCTP_EVENTS but SCTP_EVENT 
  is the more generic one and will be added soon.
  
* Peeling Off an Association


## Additional Documents   

* [Quick Start](quick_start.md)
* [Sockets API Extensions for the SCTP](https://tools.ietf.org/html/rfc6458)