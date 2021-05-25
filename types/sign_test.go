package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestRSS3_SetSign(t *testing.T) {
	demo := []byte(`
{
    "id": "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944",
    "@version": "rss3.io/version/v0.1.0-rc.0",
    "date_created": "2009-05-01T00:00:00.000Z",
    "date_updated": "2021-05-08T16:56:35.529Z",

    "profile": {
        "name": "DIYgod",
        "avatar": ["dweb://diygod.jpg", "https://example.com/diygod.jpg"],
        "bio": "写代码是热爱，写到世界充满爱！",
        "tags": ["demo", "lovely", "technology"]
    },

    "items": [{
        "id": "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944-item-1",
        "authors": ["0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944"],
        "summary": "Yes!!",
        "date_published": "2021-05-09T16:56:35.529Z",
        "date_modified": "2021-05-09T16:56:35.529Z",

        "type": "comment",
        "upstream": "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944-item-0"
    }, {
        "id": "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944-item-0",
        "authors": ["0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944"],
        "title": "Hello World",
        "summary": "Hello, this is the first item of RSS3.",
        "date_published": "2021-05-08T16:56:35.529Z",
        "date_modified": "2021-05-08T16:56:35.529Z",

        "contents": [{
            "address": ["dweb://never.html", "https://example.com/never.html"],
            "mime_type": "text/html"
        }, {
            "address": ["dweb://never.jpg"],
            "mime_type": "image/jpeg"
        }],

        "@contexts": [{
            "type": "comment",
            "list": ["0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944-item-1"]
        }, {
            "type": "like",
            "list": ["0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
        }]
    }],
    "items_next": "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944-items-0",

    "links": [{
        "type": "follow",
        "list": ["0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
    }, {
        "type": "superfollow",
        "list": ["0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
    }],
    "@backlinks": [{
        "type": "follow",
        "list": ["0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"]
    }],

    "assets": [{
        "type": "some experience point",
        "content": "100"
    }]
}

`)

	var ret RSS3

	if err := json.Unmarshal(demo, &ret); err != nil {
		t.Error(err)
	}

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Using private key: %x", crypto.FromECDSA(key))

	t.Logf("Using public key: %x", crypto.FromECDSAPub(key.Public().(*ecdsa.PublicKey)))

	t.Log("Using Address: ", crypto.PubkeyToAddress(key.PublicKey).String())

	ret.SetSign(key)

	jsonBytes, err := json.MarshalIndent(ret, "", "\t")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(jsonBytes))

	if success, err := ret.CheckSign(); err != nil {
		t.Error(err)
	} else {
		t.Log("No errors")
		if !success {
			t.Fail()
		}
	}

}
