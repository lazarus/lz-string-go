package LZString

import (
	"fmt"
	"testing"
)

func Test_CompressToBase64(t *testing.T) {
	var dataStr = `{"cNounce":"80309","cRq":{"d":"xeoGCEv+djKzc6ntAdAgZxSfmJcB5E1kLTEFsN1CGukpZkAstuTm8157S94AAYL83gT4cm0q2Hlspb2WKLfanSTXiHu9jWmtI8/NSAUULpdy9R5XxJlz+7mojw42EqGNC+loqPsGHZEbp0+b9lrOGhML8CBcEduUNq2mSTOUjgusEd8ug6+JlwzRC2kABHJDGLFX1LDRaT9WbMEEC/TeQ2bykPoBEYNrG/qxTcYwgshpAwkvHYlCXH0v1NbyBG+DI7E1dYMHQMAPNvxK450Wqh3W9Rq3/iqLrAOTMYXwuXuui82ch47qW5XKTCjLoQ5+TM+gM/EYQ6A+vWw7dkBAAJcekaohmAItymVJ/v234YzavDTna0vRDoeZZyf/v2dLPuz25Rc4ZsUX3M2L9Bne11xI2KWnoCCgsrfwCZ5d6NFRnH1Ld/jO67mVJj1SamexITk9WjwVhSetyjfw3o7Tw4eQ/4WaTe1UcQa2MMZgxO1HZEuNjR3k4MORJATbfwdxsBxhvFn2MXTvDSQMwS6LPKmR3nz3JHhrDgRj2KangIJdKLpFCP+lIK6SPghKJvNM7jXyFLqiucpYMeMqyjyOOn5cy2ANIRiO5HppjaCs36n8xM/qlaSrmbaK7jGMkD1CokK8VHpf4rMkF5VSLhhYHwXZXdSAq2GPzxY+0AOi80Hi3+6HF8e0rQ1nYGOTWrplrzeW/vCRUEiBAq+3SdKeoA==","hh":"tEtDmj9NO12CQujtDAf6znTptSdPuOFPcXitzqFb5a4=","i1":"MmEhOUHiGsRJZjij1zCWjQ==","i2":"ukacN6jPVDqrPE8x7lSyqg==","m":"P76F7Om9PMh31Wrui42tfu7uvpVw9MsguMLcoozrHaY=","ra":"TW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzg3LjAuNDI4MC4xNDEgU2FmYXJpLzUzNy4zNg==","rm":"R0VU","ru":"aHR0cHM6Ly93d3cuc25lYWtlcnNuc3R1ZmYuY29tLw==","t":"MTYxODM3NDcyNy4zODkwMDA=","uh":"sfgiHR/bGP0DFXSp4KJXpJnmQXuNHfWxwUhpX9SCMz4="},"chC":0,"chCAS":0,"chLog":{"0":{"start":1618374728142},"c":1},"chReq":"non-interactive","cvId":"2","oV":1}`
	var d = `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	fmt.Println(CompressToBase64(dataStr, d))
}

func TestCompressFromString(t *testing.T) {
	type args struct {
		uncompressed string
		chars        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressToBase64(tt.args.uncompressed, tt.args.chars); got != tt.want {
				t.Errorf("CompressFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompressFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressToBase64("1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer", _defaultKeyStrBase64)
	}
}

func BenchmarkCompressFromString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressToBase64("3123123", _defaultKeyStrBase64)
	}
}

func BenchmarkCompressToBase64(b *testing.B) {
	str := `{"cNounce":"36728","cRq":{"d":"oraaF6KB66uNJPX9UPeeiXActri35Lwu5mv+YI1bqHV1rotfmyjcn9iBX5LZto7WyVo6iQlpqpGTRZ3WeOFTCGsJBdry5WLcXcxGqhuJixjpDcv2W9h4geRwf+agasGzInsjMYBCg6GnzYmxxpdVp4U16pzjPJoikxflHq0XG+GkxMnMIAPWgpKeyNAy0Aj7WdTo0pmP3Cn6B4NrNFXhzekJG7IBR8KkfUXlteieFKS3k0WRH09OsmA+GxMz8Pag0KiXN8+tIs19PqDMHNoeJBkLMtmLKTzExtxP2wgaVjWh1UPUim7M7yLGkmdv5bU3Zuz64ASvr8IEWmiZVJYenz9hFA4MhYtNRZVyR1AYx41Fb0hvRcrQG5M5vsZXymJykAYWITIthMw6rxHnTI+Vbx6E46Yu9/QcJ4xT5MWmW8zXRIM+jp5fow4/3335tRsMSdyIL4bE+ulw52cTVXrpNqhqCatH/wHHyUshuhjsFPMMxs6WoRFW4ceOX8vi7R8TtprsOtSmWg1OK7hYrl99ZXfIcbcJwRjRzfdsYrTN5FrpCqdk2DjD5UXG7EDcI+LodnsK4PyDAMzzAFyzubmQFICT8spvRrnUoDqMzCq7ViWt2JpwqDt5gmGMHZJg5VWkGYG1EjwkLr8FZ9gh2ezHeg7e0siiPDf6AbkvU4KCmR1Ht7qgOe0l+LDM1t/m7o1tUwkoY3J1poq0U2juD+Itcg==","hh":"QSlXmfS2NeZCSV7a6l5l5OS3E6mq9iZtckO4Lj4aWug=","i1":"bSoctx5zCBRkkKOF5bXUCQ==","i2":"XPiyF0pHvrNO7OakWEMoOg==","m":"eqV80Sw4dkgMfJScOzRO4cIGos2UPT/g3Fn2WVKKC5Y=","ra":"TW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzg3LjAuNDI4MC4xNDEgU2FmYXJpLzUzNy4zNg==","rm":"R0VU","ru":"aHR0cHM6Ly93d3cud29vZHdvb2QuY29tLw==","t":"MTYxODM4NDAzMi4wNzAwMDA=","uh":"sfgiHR/bGP0DFXSp4KJXpJnmQXuNHfWxwUhpX9SCMz4="},"chC":0,"chCAS":0,"chLog":{"0":{"start":1618384032901},"c":1},"chReq":"non-interactive","cvId":"2","oV":1}`
	for i := 0; i < b.N; i++ {
		CompressToBase64(str, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=")
	}
}
