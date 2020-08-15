## myip

myip is a command line interface (CLI) for doing client ip address retrieval, written in Go and utilizing https://fast.com/ and [https://nord](https://nordvpn.com/what-is-my-ip/). It is memory efficient, performant, and descriptive in the information that it provides. 

### Installation

Assuming you already have Golang installed on your machiine, simply run
```
go get github.com/claytonblythe/myip
```


### Usage

```
myip
```


### Visual Output
```
~ $ myip

Using fast.com...
Cincinnati, US, 8.21.101.100

Using nordvpn.com...
Cincinnati, Ohio, United States, 8.21.101.100, Zillow,
https://www.google.com/maps/search/?api=1&query=39.152600,-84.386900

```



### Next Steps

Add Nord VPN support for more granular ISP information