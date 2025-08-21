# all-language

kubectl create namespace datadog
kubectl create secret generic datadog-secret -n datadog --from-literal api-key=<DATADOG_API_KEY> --from-literal app-key=<DATADOG_APP_KEY>
helm repo add datadog https://helm.datadoghq.com
helm repo update
helm install datadog datadog/datadog -n datadog -f values.yaml
helm upgrade datadog datadog/datadog -n datadog -f values.yaml

k apply -f all.yaml 


