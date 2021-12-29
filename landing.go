package main

import "net/http"

var landingTmpl = []byte(`
<!doctype html>
<html>
	<head><title>Adjust Square</title></head>
	<body style="background:url('https://stage.app.adjustsquare.com/assets/img/covers/auth-side-cover.jpg'); background-size: cover;">
	</body>
</html>
`)

func handleLanding(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(200)
	rw.Write(landingTmpl)
}
