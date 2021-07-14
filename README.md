# Deploy 
```
kubectl apply -f deploy/monitor_v1_podhealthcheck_crd.yaml
kubectl apply -f deploy/cluster_role.yaml
kubectl apply -f deploy/service_account.yaml
kubectl apply -f deploy/cluster_role_binding.yaml
kubectl apply -f deploy/operator.yaml
```

# Create Pod Health Check
```
cat >> pod_health_check.yaml << EOF
apiVersion: monitor.kphc/v1
kind: PodHealthCheck
metadata:
  name: example-podhealthcheck
spec:
  # Add fields here
  dingTalk:
    webhook: xxxx
    secret: xxxx
  namespace: xxx # resource namespace
  labelSelector: # label selector to selector pods
    key: value
  # only support http now
  healthCheck:
    interval: 60 # unit is second
    port: 80
    path: /path
    timeout: 5 # unit is second
EOF
kubectl apply -f pod_health_check.yaml
```