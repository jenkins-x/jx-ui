package repositories

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	jenkinsxv1 "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/typed/jenkins.io/v1"
)

type Repositories struct {
	SrClient jenkinsxv1.SourceRepositoryInterface
	render   *render.Render
}

func NewRepository() *Repositories {
	r := &Repositories{}
	r.render = render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	return r
}

// RepositoriesHandler function
func (rs *Repositories) RepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	repo, err := rs.SrClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs.render.JSON(w, http.StatusOK, repo.Items) //nolint:errcheck
}
