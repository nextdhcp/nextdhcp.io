---
title: "Introduction"
date: 2019-09-05T12:26:25+02:00
draft: false
---

<br />
<center><img src="/images/logos/nextdhcp-wo-text.svg" alt="NextDHCP" width="170px"></center>
<br />
<br />

NextDHCP is simle but extensible DHCP server that chains plugins. It's based on the concept of the [Caddy Webserver](https://caddyserver.com) and [CoreDNS](https://coredns.io). It's written in [Go](https://golang.org), runs on lots of different platforms, is open source and licensed under the [MIT License](https://github.com/nextdhcp/nextdhcp/blob/master/LICENSE). Development, feature requests, bug reports or any other kind of issues are managed on [Github](https://github.com/nextdhcp).

---

DHCP stands for Dynamic Host Configuration Protocol and provides network clients with information required to participate in the overall communication. NextDHCP strives to fullfill the needs of home users as well as complicated buisness enviroments. Nearly every single feature NextDHCP offers is implemented as a plugin. Thus, if you're searching for something NextDHCP does not offer out-of-the-box, you may find it easy enough to add the required feature youself. Contributions to NextDHCP and all related projects are always welcome!  
    
<br/>

Regardless of why you need a DHCP server; playing some games on a LAN party, setting up network boot, some kind of home-automation (user-presence tracking) or even more complicated DHCP setups, NextDHCP will have you covered. If not, it will be easy to extend NextDHCP to do what you want, just a few [Go](https://golang.org) skills are required. Checkout the list of [plugins](https://github.com/nextdhcp/nextdhcp/blob/master/README.md#plugins) for more information on the current feature set!

---

<center><a href="/blog/install/" class="button signup-button rounded raised">Get started</a></center>