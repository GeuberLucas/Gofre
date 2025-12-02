package reverseproxy

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ReverseProxy struct {
	routes map[string][]string
}

func NewProxyRoutes() *ReverseProxy {
	return &ReverseProxy{
		routes: map[string][]string{
			"auth":        {"http://auth:80"},
			"transaction": {"http://transactions:80"},
			"budgets":     {},
			"forecast":    {},
			"goals":       {},
			"investments": {},
			"property":    {},
			"reports":     {},
			"simulator":   {},
		},
	}
}
func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	targetHost, err := rp.selectMicroService(r.URL.Path)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	urlMicroService, err := url.Parse(targetHost)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	proxy := httputil.NewSingleHostReverseProxy(urlMicroService)

	originHost := r.Host
	originPath := r.URL.Path
	r.Host = urlMicroService.Host
	r.URL.Path = prepareCompleteUrlPath(originPath)

	proxy.ServeHTTP(w, r)
	r.Host = originHost
	r.URL.Path = originPath
}

func (rp *ReverseProxy) selectMicroService(path string) (string, error) {

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	if len(parts) < 2 {
		return "", errors.New("Invalid path structure")
	}
	serviceName := parts[1]
	targetUrl, exists := rp.routes[serviceName]
	if !exists || len(targetUrl) == 0 {
		return "", errors.New("Service not found")
	}

	return targetUrl[0], nil

}

func prepareCompleteUrlPath(path string) string {

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	newpath := strings.TrimPrefix(path, "/"+parts[0]+"/"+parts[1])

	return newpath

}

// type UserAuthenticated struct {
// 	UserId int `json:"user_id"`
// }

// func (rp *ReverseProxy) VerifyAuthentication(ctx context.Context, headers http.Header) (int, error) {
// 	microService, err := rp.selectMicroService("/api/auth")
// 	if err != nil {
// 		return 0, err
// 	}

// 	baseURL, err := url.Parse(microService)
// 	if err != nil {
// 		return 0, err
// 	}

// 	remoteUrl := baseURL
// 	remoteUrl.Path = "/isAuthenticated"
// 	println(remoteUrl)
// 	request, err := http.NewRequestWithContext(ctx, "GET", remoteUrl.String(), nil)
// 	if err != nil {
// 		return 0, err
// 	}

// 	request.Header = headers.Clone()

// 	client := &http.Client{
// 		Timeout: 5 * time.Second,
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {

// 		return 0, fmt.Errorf("auth service returned status: %d", resp.StatusCode)
// 	}
// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	var user UserAuthenticated
// 	if err := json.Unmarshal(bodyBytes, &user); err != nil {
// 		return 0, fmt.Errorf("failed to decode user: %w", err)
// 	}

// 	if user.UserId == 0 {
// 		return 0, fmt.Errorf("authentication:received invalid user ID")
// 	}

// 	return int(user.UserId), nil
// }
