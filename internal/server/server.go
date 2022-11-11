package server

import (
	"jx-ui/internal/server/pipelines"
	"jx-ui/internal/server/repositories"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned"
	jenkinsxv1 "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/typed/jenkins.io/v1"
	"github.com/rs/cors"

	tknclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/unrolled/render"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// Server struct holds the external dependencies requires for running the UI backend
type Server struct {
	router       *mux.Router
	Addr         string
	Server       *http.Server
	JxIface      versioned.Interface
	JxClient     jenkinsxv1.PipelineActivityInterface
	TknClient    tknclient.Interface
	KubeClient   kubernetes.Interface
	render       *render.Render
	Namespace    string
	Pipeline     *pipelines.Pipeline
	Repositories *repositories.Repositories
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func NewServer() *Server {
	s := &Server{}
	if s.Addr == "" {
		s.Addr = (os.Getenv("UI_ADDR"))
		if s.Addr == "" {
			s.Addr = (":9200")
		}
	}
	s.render = render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	s.router = mux.NewRouter()
	s.Pipeline = pipelines.NewPipeline()
	s.Repositories = repositories.NewRepository()
	s.router = registerRoutes(s)
	// This is only required for dev mode
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "PUT"},
		AllowCredentials: true,
	})
	s.Server = &http.Server{
		Handler: c.Handler(s.router),
		Addr:    s.Addr,
	}
	return s
}
