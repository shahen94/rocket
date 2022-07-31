package rocket

import "net/http"

func (r *Rocket) SessionLoad(next http.Handler) http.Handler {
	r.InfoLog.Println("SessionLoad called")
	return r.Session.LoadAndSave(next)
}
