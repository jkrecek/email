package emim

import (
	"testing"
	"github.com/scorredoira/email"
	"github.com/jkrecek/goconf"
	"reflect"
)

func Test(t *testing.T) {
	m := email.NewMessage("TEST", "This is test")
	em := Create(m)
	implType := reflect.TypeOf((*goconf.EmailMessage)(nil)).Elem()
	msgType := reflect.TypeOf(em)
	if !msgType.Implements(implType) {
		t.Fatal("Email Message test does not have proper goconf.EmailMessage implementation")
	}
}
