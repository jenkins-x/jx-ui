package pipelines

import (
	"context"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	jenkinsxv1 "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/typed/jenkins.io/v1"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/naming"
	pipelineapi "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tknclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/unrolled/render"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type pipelinesRes struct {
	GitOwner      string
	GitRepository string
	GitBranch     string
	Build         string
	Status        string
	StartTime     *metav1.Time
	EndTime       *metav1.Time
	Name          string
}

type Pipeline struct {
	JxClient  jenkinsxv1.PipelineActivityInterface
	TknClient tknclient.Interface
	render    *render.Render
	Namespace string
}

func NewPipeline() *Pipeline {
	p := &Pipeline{}
	p.render = render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	return p
}

// PipelinesHandler function
func (p *Pipeline) PipelinesHandler(w http.ResponseWriter, r *http.Request) {
	pa, err := p.JxClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Pipelines := make([]pipelinesRes, len(pa.Items))

	for k := range pa.Items {
		Pipelines[k].GitOwner = pa.Items[k].Spec.GitOwner
		Pipelines[k].GitRepository = pa.Items[k].Spec.GitRepository
		Pipelines[k].GitBranch = pa.Items[k].Spec.GitBranch
		Pipelines[k].Build = pa.Items[k].Spec.Build
		Pipelines[k].Status = pa.Items[k].Spec.Status.String()
		Pipelines[k].StartTime = pa.Items[k].Spec.StartedTimestamp
		Pipelines[k].EndTime = pa.Items[k].Spec.CompletedTimestamp
		Pipelines[k].Name = pa.Items[k].ObjectMeta.Name
	}

	sort.Slice(Pipelines, func(i, j int) bool {
		return Pipelines[j].StartTime.Before(Pipelines[i].StartTime)
	})

	p.render.JSON(w, http.StatusOK, Pipelines) //nolint:errcheck
}

// PipelineHandler function
func (p *Pipeline) PipelineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]
	branch := vars["branch"]
	build := vars["build"]
	name := naming.ToValidName(owner + "-" + repo + "-" + branch + "-" + build)
	method := r.Method

	switch method {
	case "POST":
		p.render.JSON(w, http.StatusMethodNotAllowed, nil) //nolint:errcheck
	case "PUT":
		pa, err := p.JxClient.
			Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			// Todo: improve error handling!
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		prName := pa.Labels["tekton.dev/pipeline"]

		pr, err := p.TknClient.TektonV1beta1().PipelineRuns(p.Namespace).Get(context.Background(), prName, metav1.GetOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if pr == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if strings.EqualFold(pa.Spec.Status.String(), "pending") || strings.EqualFold(pa.Spec.Status.String(), "running") {
			pr.Spec.Status = pipelineapi.PipelineRunSpecStatusCancelled
		}
		_, err = p.TknClient.TektonV1beta1().PipelineRuns(p.Namespace).Update(context.Background(), pr, metav1.UpdateOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.render.JSON(w, http.StatusOK, "pipeline "+name+" stopped") //nolint:errcheck
	default:
		pa, err := p.JxClient.
			Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			// Todo: Send 404 when pipeline does not exist!
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.render.JSON(w, http.StatusOK, pa) //nolint:errcheck
	}
}
