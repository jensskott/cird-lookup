# cird-lookup

## Build 

```shell
make all
make install
```

## Usage

```shell
cidrlookup -r "aws-region-1" -t "name tag to search for"
```

## Explaination

Lookup up existing VPC and finds ips and itterates with one on the second octet.

Works only on 10.X.0.0 as of this moment