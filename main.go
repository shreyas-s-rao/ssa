package main

import (
	"context"
	"fmt"
	"os"
	"time"

	druidv1alpha1 "github.com/gardener/etcd-druid/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	err := druidv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		fmt.Printf("%v", err)
	}

	kubeConfigPath, err := getEnvOrError("KUBECONFIG")
	if err != nil {
		fmt.Printf("%v", err)
	}

	etcdClient, err := getKubernetesClient(kubeConfigPath)
	if err != nil {
		fmt.Printf("%v", err)
	}

	etcd := &druidv1alpha1.Etcd{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "druid.gardener.cloud/v1alpha1",
			Kind:       "Etcd",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "etcd-main",
			Namespace: "shoot",
		},
		Status: druidv1alpha1.EtcdStatus{
			// Members: []druidv1alpha1.EtcdMemberStatus{
			// 	{
			// 		ID:                 "1",
			// 		Name:               "member1",
			// 		Role:               druidv1alpha1.EtcdRoleMember,
			// 		Status:             druidv1alpha1.EtcdMemberStatusReady,
			// 		Reason:             "up and running",
			// 		LastUpdateTime:     metav1.NewTime(time.Now()),
			// 		LastTransitionTime: metav1.NewTime(time.Now()),
			// 	},
			// },
			Members: []druidv1alpha1.EtcdMemberStatus{
				{
					ID:                 "2",
					Name:               "member2",
					Role:               druidv1alpha1.EtcdRoleMember,
					Status:             druidv1alpha1.EtcdMemberStatusReady,
					Reason:             "up and running",
					LastUpdateTime:     metav1.NewTime(time.Now()),
					LastTransitionTime: metav1.NewTime(time.Now()),
				},
			},
		},
	}

	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Minute)
	defer cancelFunc()

	err = etcdClient.Status().Patch(ctx, etcd, client.Apply, client.FieldOwner("shreyas2"), client.ForceOwnership)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("\nDone.")
}

func getEnvOrError(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", fmt.Errorf("environment variable not found: %s", key)
}

func getKubeconfig(kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

func getKubernetesClient(kubeconfigPath string) (client.Client, error) {
	config, err := getKubeconfig(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return client.New(config, client.Options{})
}
