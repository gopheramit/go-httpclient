package gomime

import "testing"

func TestHeadears(t *testing.T) {
	t.Run("test headers", func(t *testing.T) {
		if HeaderContentType != "Content-Type" {
			t.Error("invalid header content type")
		}

		if HeaderUserAgent != "User-Agent" {
			t.Error("invalid header user agent")
		}

		if ContentType != "Content-Type" {
			t.Error("invalid content type")
		}

		if ContentTypeJson != "application/json" {
			t.Error("invalid content type json")
		}

		if ContentTypeXml != "application/xml" {
			t.Error("invalid content type xml")
		}

		if ContentTypeOctet != "application/octet-stream" {
			t.Error("invalid content type octet")
		}
	})
}
