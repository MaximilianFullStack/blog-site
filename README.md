# MaxFullStack Blog Site

## Stack

Client Side: [HTMX](https://htmx.org)
Server Side: Golang + [Templ](https://templ.guide/)
Backend: Golang
DB: [Turso](https://turso.tech/) (Postges)

## Getting Started

### Manual

1. **Generate Go Templates**: `templ generate`

2. **Run Server**: `go run .`

3. **Kill Port**: `sh -c 'lsof -i :${PORT:-3000} -t | xargs kill'`

### 'Hot Reloading'

`./reload.sh`
