package types

import (
	"encoding/json"
	"testing"
)

func TestRSS3_CheckSign(t *testing.T) {
	demo := []byte(`
{
	"@version": "rss3.io/version/v0.1.0",
	"date_created": "2021-05-25T15:32:18.099Z",
	"date_updated": "2021-05-25T15:32:18.099Z",
	"id": "0x2C5E0d168D1dc1C5c606C1D74A0C79FeCE8b5c8e",
	"items": [
		{
			"authors": [
				"0x2C5E0d168D1dc1C5c606C1D74A0C79FeCE8b5c8e"
			],
			"date_modified": "2021-05-25T15:32:18.099Z",
			"date_published": "2021-05-25T15:32:18.099Z",
			"id": "0x2C5E0d168D1dc1C5c606C1D74A0C79FeCE8b5c8e-item-0",
			"signature": "0x586b749d9facb8144fe087816209e288556d34c202aacc9c99ad8e1074add16a4f7a7c1427517db147e1a427f950084510b446e08ef9b66f970da8d6738162d91b",
			"title": "Hello"
		}
	],
	"profile": {
		"name": "RSS3"
	},
	"signature": "0x37cf04404520171c20427089c00075e08b31bcbc1639e77402d8fbabb155b67e753bc6f1e7ce74c514fd0a2c160d6c2d3b262794f3383edf472d062ba99a25231b"
}
`)

	var ret RSS3

	if err := json.Unmarshal(demo, &ret); err != nil {
		t.Error(err)
	}

	//privateKeyHex := "0x35972403612cce8aef98e429b3af64d572b298a8b2f4331913d0a8ad1dc93322"

	if success, err := ret.CheckSign(); err != nil {
		t.Error(err)
	} else {
		t.Log("No errors")
		if !success {
			t.Fail()
		}
	}

}
