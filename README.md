# go-echo
Basic HTTP server in Golang for test Kubernetes

### Usage

```shell script
docker run -p 8080:8080 borosr/go-echo:latest
```

### Response
`Content-Type: plain/text`
#### Format sample
```
Request [instance id] is:
Url: 
Method is: 
Protocol: 
Headers:
Accept: 
User-Agent: 
Body: # if present
```

