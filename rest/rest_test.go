package rest

import (
	"net/http"
	"testing"
)

func Test_signature(t *testing.T) {
	secret := "chNOOS4KvNXR_Xq4k4c9qsfoKWvnDecLATCRlcBwyKDYnWgO"

	type args struct {
		secret     string
		method     string
		requestURI string
		expires    string
		dataPost   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get instrument",
			args: args{
				secret:     secret,
				method:     http.MethodGet,
				requestURI: "/api/v1/instrument",
				expires:    "1518064236",
			},
			want: "c7682d435d0cfe87c16098df34ef2eb5a549d4c5a3c2b1f0f77b8af73423bf00",
		},
		{
			name: "get instrument with query params",
			args: args{
				secret:     secret,
				method:     http.MethodGet,
				requestURI: "/api/v1/instrument?filter=%7B%22symbol%22%3A+%22XBTM15%22%7D",
				expires:    "1518064237",
			},
			want: "e2f422547eecb5b3cb29ade2127e21b858b235b386bfa45e1c1756eb3383919f",
		},
		{
			name: "post ",
			args: args{
				secret:     secret,
				method:     http.MethodPost,
				requestURI: "/api/v1/order",
				expires:    "1518064238",
				dataPost:   `{"symbol":"XBTM15","price":219.0,"clOrdID":"mm_bitmex_1a/oemUeQ4CAJZgP3fjHsA","orderQty":98}`,
			},
			want: "1749cd2ccae4aa49048ae09f0b95110cee706e0944e6a14ad0b3a8cb45bd336b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signature(tt.args.secret, tt.args.method, tt.args.requestURI, tt.args.expires, tt.args.dataPost); got != tt.want {
				t.Errorf("signature() = %v, want %v", got, tt.want)
			}
		})
	}
}
