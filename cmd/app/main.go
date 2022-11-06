package main

import (
	"fmt"
	"os"

	"jx-ui/internal/server"

	"github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/jxclient"
	"github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"
	tknclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := server.NewServer()
	config, err := kubeclient.NewFactory().CreateKubeConfig()
	if err != nil {
		return err
	}

	if s.Namespace == "" {
		s.Namespace = "jx"
	}

	if s.Pipeline.Namespace == "" {
		s.Pipeline.Namespace = "jx"
	}

	if s.JxClient == nil {
		jxClient, err := versioned.NewForConfig(config)
		if err != nil {
			return err
		}
		s.JxClient = jxClient.JenkinsV1().PipelineActivities(s.Namespace)
	}

	if s.Pipeline.JxClient == nil {
		jxClient, err := versioned.NewForConfig(config)
		if err != nil {
			return err
		}
		s.Pipeline.JxClient = jxClient.JenkinsV1().PipelineActivities(s.Namespace)
	}

	if s.Repositories.SrClient == nil {
		jxClient, err := versioned.NewForConfig(config)
		if err != nil {
			return err
		}
		s.Repositories.SrClient = jxClient.JenkinsV1().SourceRepositories(s.Namespace)
	}

	if s.TknClient == nil {
		tknClient, err := tknclient.NewForConfig(config)
		if err != nil {
			return err
		}
		s.TknClient = tknClient
	}

	if s.Pipeline.TknClient == nil {
		tknClient, err := tknclient.NewForConfig(config)
		if err != nil {
			return err
		}
		s.Pipeline.TknClient = tknClient
	}

	if s.KubeClient == nil {
		kubeClient, err := kubernetes.NewForConfig(config)
		if err != nil {
			return err
		}
		s.KubeClient = kubeClient
	}

	s.JxIface, err = jxclient.LazyCreateJXClient(s.JxIface)
	if err != nil {
		return err
	}

	err = s.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
