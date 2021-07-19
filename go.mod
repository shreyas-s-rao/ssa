module github.com/gardener/ssa

go 1.16

require (
	github.com/gardener/etcd-druid/api v0.5.1
	k8s.io/api v0.21.2 // indirect
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
  sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	github.com/gardener/etcd-druid/api => ./../etcd-druid/api
	k8s.io/api => k8s.io/api v0.21.2
	k8s.io/client-go => k8s.io/client-go v0.21.2
)
