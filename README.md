# GO Security Headers

This tool will scan one or more URL's for HTTP Headers provided in the headers.txt file and show you the values.

I built this tool to quickly scan over endpoints and view if HSTS, Secure Cookies and X-Content-Type-Options were being set properly.

## Instructions
### Build
```
git clone https://gitlab.com/medoix/gosecheaders.git
cd gosecheaders
go build
```

### Run
If you do not pass any paramaters you will scan and get the results for https://google.com
```
gosecheaders
```

To scan one URL
```
gosecheaders -url domaintoscan.com
```

To scan a list of URL's
> you will need to create a scope.txt file with one URL per line

> blah.com

> blah2.com

> bluh.com

```
gosecheaders -scope /path/to/scope.txt
```