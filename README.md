# Check Security

This tool will scan one or more URL's for HTTP Headers provided in the headers.txt file and show you the values.

I built this tool to quickly scan over endpoints and view if HSTS, Secure Cookies and X-Content-Type-Options were being set properly.

## Instructions
### Build
```
git clone https://github.com/medoix/checksec.git
cd checksec
go build
```

### Run
If you do not pass any paramaters you will scan and get the results for https://google.com
```
checksec
```

To scan one URL
```
checksec -url domaintoscan.com
```

To scan a list of URL's
> you will need to create a scope.txt file with one URL per line
```
domain1.com
domain2.com
domain3.com
```
```
checksec -scope /path/to/scope.txt
```

## Planned Features
- TLS Checks
- Report Output
