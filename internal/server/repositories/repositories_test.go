package repositories_test

import (
	"io"
	"jx-ui/internal/server"
	"jx-ui/internal/server/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	jenkinsv1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"
	fakejx "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/fake"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const fakeNamespace = "jx-fake"

func TestRepositoriesHandler(t *testing.T) {
	s := server.NewServer()
	rs := repositories.NewRepository()
	s.Repositories = rs
	s.Namespace = fakeNamespace
	srClient := fakejx.NewSimpleClientset(&jenkinsv1.SourceRepositoryList{
		TypeMeta: v1.TypeMeta{},
		ListMeta: v1.ListMeta{},
		Items: []jenkinsv1.SourceRepository{
			{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{},
				Spec:       jenkinsv1.SourceRepositorySpec{},
			},
		},
	})
	rs.SrClient = srClient.JenkinsV1().SourceRepositories(s.Namespace)
	r := httptest.NewRequest(http.MethodGet, "/api/v1/repositories", nil)
	w := httptest.NewRecorder()
	rs.RepositoriesHandler(w, r)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error("Unexpected error when reading response")
	}
	if string(data) == "" {
		t.Error("Did not receive any pipelines")
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected response Code %d", w.Code)
	}
}
