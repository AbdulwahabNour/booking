package render

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AbdulwahabNour/booking/internal/config"
	"github.com/alexedwards/scs"
)





var testapp config.AppConfig
var session *scs.SessionManager


func TestMain(m *testing.M){
	session = scs.New()
	session.Lifetime = 24*time.Minute
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = false
     testapp.Session = session
app = &testapp
	os.Exit(m.Run())
}

