package middleware

import (
	"github.com/AliasYermukanov/AlfaAuth/src/dbo"
	"net/http"
)

func AuthMiddleware(h http.Handler, scope string, action string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("UserUID") == "" {

			uid := r.Header.Get("UserUID")
			scopesMap, err := dbo.GetScopesByUID(uid)
			if err!= nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if scopesMap[scope] != "" {
				switch action {
				case "CREATE":
					if []rune(scopesMap[scope])[0] == '1' {
						return
					} else {
						http.Error(w,"Access Denied", http.StatusForbidden)
						return
					}
				case "READ":
					if []rune(scopesMap[scope])[1] == '1' {
						return
					} else {
						http.Error(w, "Access Denied", http.StatusForbidden)
						return
					}
				case "UPDATE":
					if []rune(scopesMap[scope])[2] == '1' {
						return
					} else {
						http.Error(w, "Access Denied", http.StatusForbidden)
						return
					}
				case "DELETE":
					if []rune(scopesMap[scope])[3] == '1' {
						return
					} else {
						http.Error(w, "Access Denied", http.StatusForbidden)
						return
					}
				}
			}else {
				http.Error(w, "Access Denied", http.StatusForbidden)
				return
			}
			return
		}
	})
}
