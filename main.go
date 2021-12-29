package main

import (
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"kmodules.xyz/client-go/apiextensions"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	catalogapi "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctrl.SetLogger(klogr.New())
	cfg := ctrl.GetConfigOrDie()

	crdClient, err := crd_cs.NewForConfig(cfg)
	if err != nil {
		return err
	}
	return EnsureCustomResourceDefinitions(crdClient)
}

// EnsureCustomResourceDefinitions ensures CRD for MySQl, DormantDatabase and Snapshot
func EnsureCustomResourceDefinitions(crdClient crd_cs.Interface) error {
	klog.Infoln("Ensuring CustomResourceDefinition...")
	crds := []*apiextensions.CustomResourceDefinition{
		dbapi.Elasticsearch{}.CustomResourceDefinition(),
		dbapi.Etcd{}.CustomResourceDefinition(),
		dbapi.MariaDB{}.CustomResourceDefinition(),
		dbapi.Memcached{}.CustomResourceDefinition(),
		dbapi.MongoDB{}.CustomResourceDefinition(),
		dbapi.MySQL{}.CustomResourceDefinition(),
		dbapi.PerconaXtraDB{}.CustomResourceDefinition(),
		dbapi.PgBouncer{}.CustomResourceDefinition(),
		dbapi.Postgres{}.CustomResourceDefinition(),
		dbapi.ProxySQL{}.CustomResourceDefinition(),
		dbapi.Redis{}.CustomResourceDefinition(),
		dbapi.RedisSentinel{}.CustomResourceDefinition(),

		catalogapi.ElasticsearchVersion{}.CustomResourceDefinition(),
		catalogapi.EtcdVersion{}.CustomResourceDefinition(),
		catalogapi.MariaDBVersion{}.CustomResourceDefinition(),
		catalogapi.MemcachedVersion{}.CustomResourceDefinition(),
		catalogapi.MongoDBVersion{}.CustomResourceDefinition(),
		catalogapi.MySQLVersion{}.CustomResourceDefinition(),
		catalogapi.PerconaXtraDBVersion{}.CustomResourceDefinition(),
		catalogapi.PgBouncerVersion{}.CustomResourceDefinition(),
		catalogapi.PostgresVersion{}.CustomResourceDefinition(),
		catalogapi.ProxySQLVersion{}.CustomResourceDefinition(),
		catalogapi.RedisVersion{}.CustomResourceDefinition(),

		appcat.AppBinding{}.CustomResourceDefinition(),
	}
	return apiextensions.RegisterCRDs(crdClient, crds)
}
