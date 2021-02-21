package namespace

//
//func AccessCheck(repository Repository, getLoginFn GetLoginFunc, getRouteAndMethodFn GetRouteAndMethodFunc) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//
//		fn := func(w http.ResponseWriter, r *http.Request) {
//
//			ctx := r.Context()
//
//			login := getLoginFn(r)
//			route, method := getRouteAndMethodFn(r)
//
//			res, err := service.CheckAccess(ctx, login, route, method)
//			if err != nil {
//				w.WriteHeader(http.StatusForbidden)
//				return
//			}
//
//			if !res.AccessAllowed {
//				w.WriteHeader(http.StatusForbidden)
//				return
//			}
//
//			ctx = setUserGroupsIntoContext(ctx, res.Groups)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		}
//
//		return http.HandlerFunc(fn)
//	}
//}
