# cert-manager

cert-manager is a Kubernetes addon to automate the management and issuance of
TLS certificates from various issuing sources.

It will ensure certificates are valid and up to date periodically, and attempt
to renew certificates at an appropriate time before expiry.

## Prerequisites

- Kubernetes 1.7+

## Installing the Chart

Full installation instructions, including details on how to configure extra
functionality in cert-manager can be found in the [getting started docs](https://docs.cert-manager.io/en/latest/getting-started/).

To install the chart with the release name `my-release`:

```console
## IMPORTANT: you MUST install the cert-manager CRDs **before** installing the
## cert-manager Helm chart
$ kubectl apply \
    -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.11/deploy/manifests/00-crds.yaml

## If you are installing on openshift :
$ oc create \
    -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.11/deploy/manifests/00-crds.yaml

## IMPORTANT: if the cert-manager namespace **already exists**, you MUST ensure
## it has an additional label on it in order for the deployment to succeed
$ kubectl label namespace cert-manager cert-manager.io/disable-validation="true"

## For openshift:
$ oc label namespace cert-manager cert-manager.io/disable-validation=true

## Add the Jetstack Helm repository
$ helm repo add jetstack https://charts.jetstack.io


## Install the cert-manager helm chart
$ helm install --name my-release --namespace cert-manager jetstack/cert-manager
```

In order to begin issuing certificates, you will need to set up a ClusterIssuer
or Issuer resource (for example, by creating a 'letsencrypt-staging' issuer).

More information on the different types of issuers and how to configure them
can be found in our documentation:

https://docs.cert-manager.io/en/latest/tasks/issuers/index.html

For information on how to configure cert-manager to automatically provision
Certificates for Ingress resources, take a look at the `ingress-shim`
documentation:

https://docs.cert-manager.io/en/latest/tasks/issuing-certificates/ingress-shim.html

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Special considerations may be required when upgrading the Helm chart, and these
are documented in our full [upgrading guide](https://docs.cert-manager.io/en/latest/tasks/upgrading/index.html).
Please check here before perform upgrades!

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the cert-manager chart and their default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imagePullSecrets` | Reference to one or more secrets to be used when pulling images | `[]` |
| `global.rbac.create` | If `true`, create and use RBAC resources (includes sub-charts) | `true` |
| `image.repository` | Image repository | `quay.io/jetstack/cert-manager-controller` |
| `image.tag` | Image tag | `v0.11.0-alpha.0` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `replicaCount`  | Number of cert-manager replicas  | `1` |
| `clusterResourceNamespace` | Override the namespace used to store DNS provider credentials etc. for ClusterIssuer resources | Same namespace as cert-manager pod
| `leaderElection.Namespace` | Override the namespace used to store the ConfigMap for leader election | Same namespace as cert-manager pod
| `extraArgs` | Optional flags for cert-manager | `[]` |
| `extraEnv` | Optional environment variables for cert-manager | `[]` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template |  |
| `serviceAccount.annotations` | Annotations to add to the service account |  |
| `resources` | CPU/memory resource requests/limits | |
| `securityContext.enabled` | Enable security context | `false` |
| `securityContext.fsGroup` | Group ID for the container | `1001` |
| `securityContext.runAsUser` | User ID for the container | `1001` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |
| `ingressShim.defaultIssuerName` | Optional default issuer to use for ingress resources |  |
| `ingressShim.defaultIssuerKind` | Optional default issuer kind to use for ingress resources |  |
| `ingressShim.defaultACMEChallengeType` | Optional default challenge type to use for ingresses using ACME issuers |  |
| `ingressShim.defaultACMEDNS01ChallengeProvider` | Optional default DNS01 challenge provider to use for ingresses using ACME issuers with DNS01 |  |
| `prometheus.enabled` | Enable Prometheus monitoring | `true` |
| `prometheus.servicemonitor.enabled` | Enable Prometheus Operator ServiceMonitor monitoring | `false`
| `prometheus.servicemonitor.namespace` | Define namespace where to deploy the ServiceMonitor resource | (namespace where you are deploying) |
| `prometheus.servicemonitor.prometheusInstance` | Prometheus Instance definition | `default` |
| `prometheus.servicemonitor.targetPort` | Prometheus scrape port | `9402` |
| `prometheus.servicemonitor.path` | Prometheus scrape path | `/metrics` |
| `prometheus.servicemonitor.interval` | Prometheus scrape interval | `60s` |
| `prometheus.servicemonitor.labels` | Add custom labels to ServiceMonitor | |
| `prometheus.servicemonitor.scrapeTimeout` | Prometheus scrape timeout | `30s` |
| `podAnnotations` | Annotations to add to the cert-manager pod | `{}` |
| `podDnsPolicy` | Optional cert-manager pod [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-policy) |  |
| `podDnsConfig` | Optional cert-manager pod [DNS configurations](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pods-dns-config) |  |
| `podLabels` | Labels to add to the cert-manager pod | `{}` |
| `priorityClassName`| Priority class name for cert-manager and webhook pods | `""` |
| `http_proxy` | Value of the `HTTP_PROXY` environment variable in the cert-manager pod | |
| `https_proxy` | Value of the `HTTPS_PROXY` environment variable in the cert-manager pod | |
| `no_proxy` | Value of the `NO_PROXY` environment variable in the cert-manager pod | |
| `webhook.enabled` | Toggles whether the validating webhook component should be installed | `true` |
| `webhook.replicaCount` | Number of cert-manager webhook replicas | `1` |
| `webhook.podAnnotations` | Annotations to add to the webhook pods | `{}` |
| `webhook.extraArgs` | Optional flags for cert-manager webhook component | `[]` |
| `webhook.resources` | CPU/memory resource requests/limits for the webhook pods | |
| `webhook.nodeSelector` | Node labels for webhook pod assignment | `{}` |
| `webhook.image.repository` | Webhook image repository | `quay.io/jetstack/cert-manager-webhook` |
| `webhook.image.tag` | Webhook image tag | `v0.11.0-alpha.0` |
| `webhook.image.pullPolicy` | Webhook image pull policy | `IfNotPresent` |
| `webhook.injectAPIServerCA` | if true, the apiserver's CABundle will be automatically injected into the ValidatingWebhookConfiguration resource | `true` |
| `cainjector.enabled` | Toggles whether the cainjector component should be installed (required for the webhook component to work) | `true` |
| `cainjector.replicaCount` | Number of cert-manager cainjector replicas | `1` |
| `cainjector.podAnnotations` | Annotations to add to the cainjector pods | `{}` |
| `cainjector.extraArgs` | Optional flags for cert-manager cainjector component | `[]` |
| `cainjector.resources` | CPU/memory resource requests/limits for the cainjector pods | |
| `cainjector.nodeSelector` | Node labels for cainjector pod assignment | `{}` |
| `cainjector.image.repository` | cainjector image repository | `quay.io/jetstack/cert-manager-cainjector` |
| `cainjector.image.tag` | cainjector image tag | `v0.11.0-alpha.0` |
| `cainjector.image.pullPolicy` | cainjector image pull policy | `IfNotPresent` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
$ helm install --name my-release -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](values.yaml)

## Contributing

This chart is maintained at [github.com/jetstack/cert-manager](https://github.com/jetstack/cert-manager/tree/master/deploy/charts/cert-manager).
