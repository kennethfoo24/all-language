
# all-language

NodeJS > Python > Java > Golang > Dotnet > Ruby > PHP

## Usage

Install my-project with npm

```bash
kubectl create namespace datadog
```

```bash
kubectl create secret generic datadog-secret -n datadog 
--from-literal api-key=<DATADOG_API_KEY> --from-literal app-key=<DATADOG_APP_KEY>
```

```bash
helm repo add datadog https://helm.datadoghq.com
```

```bash
helm repo update
```

```bash
helm install datadog datadog/datadog -n datadog -f values.yaml
```

```bash
helm upgrade datadog datadog/datadog -n datadog -f values.yaml
```
```bash
k apply -f all.yaml 
```
```bash
k get svc
```
```bash
curl http://<LOAD_BALANCER_IP>/nodejs
```
    
