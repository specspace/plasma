package plasma

import (
	"testing"
)

func TestServer_ListenAndServe(t *testing.T) {
	/*tt := []struct {
		srv  Server
		addr string
		err  string
	}{
		{
			srv:  Server{},
			addr: "[::]:25565",
		},
		{
			srv:  Server{Addr: ":25565"},
			addr: "[::]:25565",
		},
		{
			srv:  Server{Addr: ":1337"},
			addr: "[::]:1337",
		},
		{
			srv:  Server{Addr: "1337"},
			addr: "[::]:1337",
		},
		{
			srv:  Server{Addr: "abc"},
			addr: "abc",
			err:  "listen tcp: address abc: missing port in address",
		},
	}

	for _, tc := range tt {
		go func() {
			if err := tc.srv.ListenAndServe(); err != nil {
				if err.Error() != tc.err {
					t.Error(err)
				}
			}
		}()

		waitForListener := func() <-chan bool {
			b := make(chan bool, 1)
			for tc.srv.listener == nil {
			}
			b <- true
			return b
		}

		select {
		case <-time.After(time.Second):
			t.Fail()
		case <-waitForListener():
		}

		if tc.srv.listener.Addr().Network() != "tcp" {
			_ = tc.srv.Close()
			t.Fail()
		}

		if tc.srv.listener.Addr().String() != tc.addr {
			_ = tc.srv.Close()
			t.Fail()
		}

		_ = tc.srv.Close()
	}*/
}
