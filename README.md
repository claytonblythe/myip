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
Fast.com result:
Seattle, US, 104.200.129.213
```



### Next Steps

Add Nord VPN support for more granular ISP information