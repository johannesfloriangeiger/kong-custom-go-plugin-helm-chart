A Helm chart for running NGINX behind Kong with a custom plugin written in Go
=

# Setup

The following instructions are based on a local Kubernetes cluster provided by Docker Compose with an NGINX Ingress
controller installed and listening on `http://localhost:80` and a Kubernetes registry running
under `http://localhost:5000`.

## 1. Build Kong plugin and image

```
env GOOS=linux GOARCH=amd64 go build -C kong/add-header \
    && (cd kong \
        && docker build . -t localhost:5000/kong-local \
        && docker push localhost:5000/kong-local)
```

## 2. Install Helm chart

```
helm install kong-custom-go-plugin-helm-chart chart
```

# Demo

To capture the custom header added by the Kong plugin, running `tcpdump` on NGINX is the quickest way:

```
kubectl exec $(kubectl get pods -o custom-columns=PodName:.metadata.name,PodUID:.metadata.uid \
    | grep nginx \
    | awk '{print $1}') \
    -- bash -c "apt-get update; apt-get install tcpdump -y; tcpdump -A 'tcp port 80'"
```

Call the API via

```
curl localhost/api
```

and the NGINX logging window should show `My-Custom-Header: Hello NGINX behind Kong!` as one of the many headers
forwarded to NGINX from Kong.

# Uninstall

```
helm uninstall kong-custom-go-plugin-helm-chart
```