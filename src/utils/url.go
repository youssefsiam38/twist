package utils

import (
	"fmt"
	"net/url"
	"path"
)

func CleanUrl(u *url.URL) string {
	if emptyUrl(*u) {
		return ""
	}
	if u.Path == "/" {
		u.Path = ""
	}
	return fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.Path, u.Query().Encode())
}

func CleanUrlHost(u *url.URL) string {
	if emptyUrl(*u) {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}


func JoinUrl(host, urlPath string) (*string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, urlPath)
	s := u.String()
	return &s, nil
}

// isValidUrl tests a string to determine if it is a well-structured url or not.
func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func emptyUrl(u url.URL) bool {
	if u.Host == "" || u.Scheme == "" {
		return true
	}
	return false
}