package notifications

import (
	"github.com/icinga/icinga-go-library/notifications/source"
	"net/url"
	"testing"
)

func TestObjectURL(t *testing.T) {
	for _, test := range []struct{ base, want string }{
		{"", ""},
		{"https://monitor.example/", "https://monitor.example/icingadb/service?name=HTTP&host.name=example.com"},
		{"https://monitor.example/icingaweb2/", "https://monitor.example/icingaweb2/icingadb/service?name=HTTP&host.name=example.com"},
	} {
		client := &Client{Config: source.Config{}}
		if test.base != "" {
			client.Icingaweb2UrlParsed, _ = url.Parse(test.base)
		}
		got := client.objectURL("/icingadb/service?name=HTTP&host.name=example.com")
		if got != test.want {
			t.Fatalf("base %q: got %q, want %q", test.base, got, test.want)
		}
	}
}
