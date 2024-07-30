# `logfmt2json`

## Installation

```
go install github.com/mircodz/logfmt2json@latest
```

## Usage

```
k logs -l -f app=service | logfmt2json | jq '.field' 
```
