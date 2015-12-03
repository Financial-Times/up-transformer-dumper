#up-transformer-dumper

A utility to dump a stream of json objects from a transformer, given it's list endpoint.

Installation:
```
go get github.com/Financial-Times/up-transformer-dumper
```

Examples:

```
up-transformer-dumper -listEndpoint http://ftaps39395-law1b-eu-t/transformers/organisations/ >/tmp/orgsdump.json
up-transformer-dumper -listEndpoint http://ftaps50665-law1a-eu-t/transformers/memberships/ >/tmp/membershipsdump.json
up-transformer-dumper -listEndpoint http://ftaps50665-law1a-eu-t/transformers/roles/ >/tmp/rolesdump.json
up-transformer-dumper -listEndpoint http://ftaps35629-law1a-eu-t/transformers/people/ >/tmp/peopledump.json

```
