# httpclient lib

Package httpclient is used as client to call 3rd api

httpclient is based on https://github.com/valyala/fasthttp

httpclient provide:

+ standard request/response format

+ standard http config (e.g http idle, http keep alive, timeout, read timeout, write timeout, retry, connection per host)

+ integrate with circuit breaker (using proxy pattern)

## How to install

1. Add GOPRIVATE in ~/.bashrc or ~/.zshrc

```bash
export GOPRIVATE=dev.azure.com
```
2. Add git config for ssh

```bash
[url "git@ssh.dev.azure.com:v3/acustombot/"]
	  insteadOf = https://dev.azure.com/acustombot/
```

3. Go to your project and install 

```bash
go get dev.azure.com/acustombot/CustomBotPlatform/httpclient.git
```