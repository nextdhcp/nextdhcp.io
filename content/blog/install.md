---
title: "Getting Started"
date: 2019-09-15T18:37:01+02:00
draft: false
---

# Installation

There are currently 2 different ways of installing NextDHCP. You can either download one of the latest pre-built releases from Github or just build it from source. We'll explain them both in order:

## Release binaries

Official release binaries are the most simple way to use and install NextDHCP. Just download the archive that suits your platform, extract it and you should be ready to go.
You can find a list of official releases on the [Github project page](https://github.com/nextdhcp/nextdhcp/releases). We recommend sticking to the latest available release:

<br />
<center>
    <p>
        <a href="https://github.com/nextdhcp/nextdhcp/releases/latest" class="button signup-button rounded raised secondary-btn" target="_blank">Latest Version</a>
    </p>
</center>
<br />

Once you downloaded the correct archive just unpack it and you should end up with a `nextdhcp` binary. 

For the "terminal" lovers, here're some commands to install it:
<br />
<br />

{{<highlight bash>}}
# Watch out, the next command does NOT automatically download
# the latest version of NextDHCP. Copy the actual link from
# the release page
curl -o nextdhcp.tar.gz https://github.com/nextdhcp/nextdhcp/releases/latest/download/nextdhcp_0.1.0_Linux_x86_64.tar.gz

# Unpack it
tar xfz nextdhcp.tar.gz
{{</highlight>}}


## Source Code

If you want to built NextDHCP from source code you'll need the following dependencies installed:

 * [Git](https://git-scm.com/)
 * [Go](https://golang.org/)

Once you got the dependencies installed (check your distribution's package manager), go ahead and clone the repository:

{{<highlight bash>}}
git clone https://github.com/nextdhcp/nextdhcp
{{</highlight>}}

Finally, we can install all dependencies and build NextDHCP. The binary will end up in the same folder.

{{<highlight bash>}}
# Download all dependencies
go get ./...

# Generate any required code
go generate ./...

# Build NextDHCP
go build .
{{</highlight>}}

# Configuration

> Although NextDHCP provides a lot of functionallity and features this section will only cover the most basic setup. For more information please refer to the list
of builtin plugins and their documentation.

NextDHCP expects a configuration file with the name Dhcpfile in the same directory where you execute the command. The syntax of the Dhcpfile is meant to be easy and
hard to get wrong. It is based on the awesome [Caddy server framework](https://caddyserver.com) and thus looks similar to the Caddyfile or Corefile (for CoreDNS). 
A simple Dhcpfile may look like the following:

{{<highlight plain>}}
192.168.0.1/24 {
    lease 1h
    range 192.168.0.100 192.168.0.200
    option router 192.168.0.1
    option nameserver 8.8.8.8 1.1.1.1
}
{{</highlight>}}

