#up-transformer-dumper

A utility to dump a stream of json objects from a transformer, given it's list endpoint.

Installation:
```
go get github.com/Financial-Times/up-transformer-dumper
```

Example:

```
up-transformer-dumper -listEndpoint http://ftaps39395-law1b-eu-t/transformers/organisations/ >/tmp/orgsdump.json^C

```
