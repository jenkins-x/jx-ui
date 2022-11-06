package pipelines_test

import (
	"io"
	"jx-ui/internal/server"
	"jx-ui/internal/server/pipelines"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	jenkinsv1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"
	fakejx "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/fake"
	faketekton "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/fake"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

const fakeNamespace = "jx-fake"

func TestPipelinesHandler(t *testing.T) {
	s := server.NewServer()
	p := pipelines.NewPipeline()
	s.Pipeline = p
	s.Namespace = fakeNamespace
	jxClient := fakejx.NewSimpleClientset(&jenkinsv1.PipelineActivityList{
		Items: []jenkinsv1.PipelineActivity{{
			ObjectMeta: v1.ObjectMeta{
				Name:      "fake-jx-fake-jx-pr-404-2",
				Namespace: fakeNamespace,
			},
			Spec: jenkinsv1.PipelineActivitySpec{
				Pipeline: "jenkins-x-test",
				Build:    "2",
				Status:   "Completed",
				StartedTimestamp: &v1.Time{
					Time: time.Now(),
				},
				CompletedTimestamp: &v1.Time{
					Time: time.Now(),
				},
				Steps:         []jenkinsv1.PipelineActivityStep{},
				GitURL:        "https://fake-jx/fake-jx.git",
				GitRepository: "fake-jx",
				GitOwner:      "fake-jx",
				GitBranch:     "PR-404",
			},
		}},
	})
	p.JxClient = jxClient.JenkinsV1().PipelineActivities(s.Namespace)
	r := httptest.NewRequest(http.MethodGet, "/pipelines", nil)
	w := httptest.NewRecorder()
	p.PipelinesHandler(w, r)
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

func TestPipelineHandler(t *testing.T) {
	testCases := []struct {
		description  string
		url          string
		expectedCode int
		method       string
	}{
		{
			description:  "POST method returns method not allowed",
			url:          "/api/v1/pipelines/fake-jx/fake-jx/PR-404/2",
			expectedCode: http.StatusMethodNotAllowed,
			method:       http.MethodPost,
		},
		{
			description:  "Passing no parameters to the pipeline url results in 404",
			url:          "/api/v1/pipelines",
			expectedCode: http.StatusNotFound,
			method:       http.MethodGet,
		},
		{
			description:  "Passing wrong parameters to the pipeline url results in 500",
			url:          "/api/v1/pipelines/fake-jx/fake-jx/PR-404/1",
			expectedCode: http.StatusInternalServerError,
			method:       http.MethodGet,
		},
		{
			description:  "Passing parameters to the pipeline url returns OK",
			url:          "/api/v1/pipelines/fake-jx/fake-jx/PR-404/2",
			expectedCode: http.StatusOK,
			method:       http.MethodGet,
		},
		{
			description:  "Stopping the pipeline returns OK",
			url:          "/api/v1/pipelines/fake-jx/fake-jx/PR-404/2",
			expectedCode: http.StatusOK,
			method:       http.MethodPut,
		},
	}

	for k, tt := range testCases {
		s := server.NewServer()
		s.Namespace = fakeNamespace
		p := pipelines.NewPipeline()
		s.Pipeline = p
		jxClient := fakejx.NewSimpleClientset(&jenkinsv1.PipelineActivityList{
			Items: []jenkinsv1.PipelineActivity{{
				ObjectMeta: v1.ObjectMeta{
					Name:      "fake-jx-fake-jx-pr-404-2",
					Namespace: fakeNamespace,
					Labels: map[string]string{
						"tekton.dev/pipeline":"fake-tekton-pr",
					},
				},
				Spec: jenkinsv1.PipelineActivitySpec{
					Pipeline: "fake-jx-fake-jx-pr-404-2",
					Build:    "2",
					Status:   "Running",
					StartedTimestamp: &v1.Time{
						Time: time.Now(),
					},
					CompletedTimestamp: &v1.Time{
						Time: time.Now(),
					},
					Steps:         []jenkinsv1.PipelineActivityStep{},
					GitURL:        "https://fake-jx/fake-jx.git",
					GitRepository: "fake-jx",
					GitOwner:      "fake-jx",
					GitBranch:     "PR-404",
				},
			}},
		})
		t.Logf("Running %d test: %s", k+1, tt.description)
		p.Namespace = fakeNamespace
		p.JxClient = jxClient.JenkinsV1().PipelineActivities(s.Namespace)
		if tt.method == http.MethodPut {
			fakeTektonClient := faketekton.NewSimpleClientset(&v1beta1.PipelineRun{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{
					Name:                       "fake-tekton-pr",
					Namespace:                  fakeNamespace,
				},
				Spec:       v1beta1.PipelineRunSpec{
					Status:              "Running",
				},
			})
			p.TknClient = fakeTektonClient
		}
		r := httptest.NewRequest(tt.method, tt.url, nil)
		w := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/pipelines/{owner}/{repo}/{branch}/{build}", p.PipelineHandler)
		router.ServeHTTP(w, r)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error("Unexpected error when reading response")
		}
		if string(data) == "" {
			t.Error("Did not receive any pipelines")
		}

		if w.Code != tt.expectedCode {
			t.Errorf("Unexpected response Code %d", w.Code)
		}
	}
}
