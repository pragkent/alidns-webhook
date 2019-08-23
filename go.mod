module github.com/pragkent/alidns-webhook

go 1.12

require (
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190822073329-cd5cf285f2a3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jetstack/cert-manager v0.8.1
	github.com/pkg/errors v0.8.0
	k8s.io/apiextensions-apiserver v0.0.0-20190413053546-d0acb7a76918
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.2.0
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190413052642-108c485f896e

replace github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190203023257-5858425f7550
