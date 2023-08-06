# Run the Demo On Kubernetes

---

We will run this demo in the `demo` namespace:
`kubectl create ns demo`
`wget -O network-policy-demo.yaml https://raw.githubusercontent.com/GoogleCloudPlatform/microservices-demo/c1536ff6e6782bb37e36d2e6eee0fa64a6461216/release/kubernetes-manifests.yaml`
`kubectl apply -f network-policy-demo.yaml -n demo`

## Kuard demo application with Network Policy

- Run the `kuard` demo application in the `demo` namespace:
`kubectl run kuard --image=gcr.io/kuar-demo/kuard-amd64:1 --replicas=1 --port=8080 -n demo`

- Expose the `kuard` application as a service:
`kubectl expose deployment kuard --type=LoadBalancer --name=kuard -n demo`

### Create a Network Policy example kuard

- Create a network policy that allows traffic to the `kuard` application from pods with the label `access: true`:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-access-to-kuard
  namespace: demo
spec:
    podSelector:
        matchLabels:
        app: kuard
    ingress:
    - from:
        - podSelector:
            matchLabels:
            access: "true"
```

- Apply the network policy

```bash
kubectl apply -f allow-access-to-kuard.yaml
```

### Create L7 visibility on the kuard application with Cilium examples

```bash
kubectl annotate pod foo -n BOO policy.cilium.io/proxy-visibility="<Egress/53/UDP/DNS>,<Egress/80/TCP/HTTP>"

kubectl annotate pod frontend-65f74c4c88-mflk7 -n demo policy.cilium.io/proxy-visibility="<Ingress/8080/TCP/HTTP>,<Egress/53/UDP/DNS>,<Egress/8080/TCP/HTTP>" --overwrite

```
