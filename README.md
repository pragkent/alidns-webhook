# alidns-webhook

Cert-manager ACME DNS webhook provider for alidns.

## Install

1. Install alidns-webhook

  ```bash
  # Install alidns-webhook to cert-manager namespace
  kubectl apply -f https://raw.githubusercontent.com/pragkent/alidns-webhook/master/deploy/bundle.yaml
  ```

2. Create secret contains alidns credentials
  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: alidns-secret
  data:
    access-key: YOUR_ACCESS_KEY
    secret-key: YOUR_SECRET_KEY

  ```

3. Example Issuer
  ```yaml
  apiVersion: certmanager.k8s.io/v1alpha1
  kind: Issuer
  metadata:
    name: letsencrypt-staging
  spec:
    acme:
      email: certmaster@example.com
      server: https://acme-staging-v02.api.letsencrypt.org/directory
      privateKeySecretRef:
        name: letsencrypt-staging-account-key
      solvers:
      - dns01:
          webhook:
            groupName: acme.yourcompany.com
            solverName: alidns
            config:
              region: ""
              accessKeySecretRef:
                name: alidns-secret
                key: access-key
              secretKeySecretRef:
                name: alidns-secret
                key: secret-key
  ```

4. Issue a certificate
```yaml
apiVersion: certmanager.k8s.io/v1alpha1
kind: Certificate
metadata:
  name: example-tls
spec:
  secretName: example-com-tls
  commonName: example.com
  dnsNames:
  - example.com
  - "*.example.com"
  issuerRef:
    name: letsencrypt-staging
    kind: Issuer
```

## Development
### Running the test suite

1. Edit `testdata/alidns/alidns-secret.yaml` and `testdata/alidns/config.json`.

2. Run test suites:

```bash
$ ./scripts/fetch-test-binaries.sh
$ TEST_ZONE_NAME=example.com go test .
```
