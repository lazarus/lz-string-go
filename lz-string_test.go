package lzstring

import (
	"fmt"
	"sync"
	"testing"
)

const largeTestString = `lDXyqyYytyxyfJLedLH0LWLzy8V+OiyL2DL1yXcbyLJbXLW7ALMX-0L7yFD-49L9ylXLr8L-+DLui0091oW8LZKo-jL8i0aLjSLLPLjqULhLFUL8yz4aL8o-M6LVV+ihLzCAWOg3S-lx2zCo-U1KyLEOj1LO7JcldV5duyTePQL6iL6yLagUDLkdQLcriJjy0FY6eJLz7LAJQisEnmTRXLQRyLtzzjQy0bI1GM9L0ss01cAHoTF+OYLZeorL-RX1AU56ioEui54FoL0p58TBOUMYc+H-ADFozMLIO8Lz+LUp6UaJzU9-Lt6jAbuSRI9DUPygdvczyUP6H0Jd26Oz+jHYah1FyLXQiSO5jy+SMd104nWP3hAKJia0qaJ56LzuuXmX5FI5oGkU0XoFm8feQMPJ8LkkTVTVI5V1k5F-Pe0X5ud9lj+-yF008Nk8Eh7A7zoq40EIlj2mc1I+o-P2ByPH1+BOQkN+Bjh8h+fx98yJFc1MmJ5Dk1bEF09FLiTx7uOykDPLOQBEhBPJud36qoZJcEhsxdV5NpfJ2Q11EyYBjELO4XnTYOTkzF6eokmb1opYk-loLlo2NoDjZiXodyPjWWk5hyO49OWk46o3KQoq4W519z8sqyyP1RPX+z03JUI4baFsjk8sF2Y6ozy8fzTjHDvSy4z+EFulq4dq+3A29z666UO-yTSq9WTHyNP6M+agDLiAzjlk6Lr8ihzOO+Z2sOk60GHPKaijLLh85+cHT6O3dO43shLZQ1P15mKKQZoc-RLqhh+FLXuD6PU086OKhR6yl6Q8iYJBjycFUGoc6s-PE-y0hjh8LYA6lk+Az6kfv-49dD66GLl-1kDGKkZo1uDj3VAQu7l07XeK6UuytJauevz6OotAQkmvL2-yXFsLOs6bX5Ez88PbPy8zK-gkzKhXX5s0FMDNhyHoi0FuOOLyu71FbMmf-15qKuD67ZFUGLqWoMyUnX-X1d4FDgzV6ev03ePKd8fuN0yXqvWouUUe46XGQTVqfzGOoS6skDXh1k5jKQ17bFc9XJ-aoocLdEP86t6l8-1TeU0D6o2QUfz66iFmTo8MLXd+Eg3PR5qGQ6kz3yNuy1zclPbKE8yRFiOkn8+1oLuD$bX6Lkp7YoaW6f6guDILX-7AXZueFzsLPVo9z6ooD-85qqFveYNQUI+izbYkz-X1iYA267oQzDPby8ZPKh4lPWdzSTJFUGhmuDTei86uP3ADZQLTzc74A7Oi1Qy6ks-aup2SXFDth++bLyVOoPhXZkPsXlkT-aLXk6ULXNh4ILtJsM2-0RD89K6uo+L0FLMFYupuhi5mxuz9Di0yjzzzRlX1ydFJozR6eT-8z2kzjBXPXdOo$z8HymWD1ibuV91M07IYDyVZA8n81o+FsfuI5GuX8d+iZLPTlO5z1S+huQZAjJ2k+Yd+upOrV-oSzj1eTHYNgU0S$YZP+kDS04G3XWhX4MtRsoAuksk3QmkQJ$ekDPhVLi8P8ok5zWmMJ8g-dZz0MyXWklP5zcuOE-4H2lFYLvsQ60LPP0-vp5-bxTzXOyuUIoy+CzyEo67TkLuy-183O1XqU-5nL6u9YFoXhmW6zQ8LYLXzzZi0oLi3yKDiLd0QILDyS0EFgdLYo-8nLku7Rd5DazxDO4$p$MDcXh0EzErtXF8zWd2FU07Zopo1-pLNKUSLcXYkjFY069VAkKLT-yF7uRS62ybuh7UW77ejPD040i5DyPb+yLLEN6W-0i9LJzEFIz+OXOoldNuOAQOo6LJzCFszj-YZPjL-LdL8-6LXQ+4PvzZz+3LbL-0kI-0okLuL20ChyLLZJbzXkL+Yi-dFbzi5LAdMLpz7Fy8ZqL7LpJdZ+7LL`

//goland:noinspection ALL
func Test_Compress(t *testing.T) {
	var dataStr = `{"cNounce":"80309","cRq":{"d":"xeoGCEv+djKzc6ntAdAgZxSfmJcB5E1kLTEFsN1CGukpZkAstuTm8157S94AAYL83gT4cm0q2Hlspb2WKLfanSTXiHu9jWmtI8/NSAUULpdy9R5XxJlz+7mojw42EqGNC+loqPsGHZEbp0+b9lrOGhML8CBcEduUNq2mSTOUjgusEd8ug6+JlwzRC2kABHJDGLFX1LDRaT9WbMEEC/TeQ2bykPoBEYNrG/qxTcYwgshpAwkvHYlCXH0v1NbyBG+DI7E1dYMHQMAPNvxK450Wqh3W9Rq3/iqLrAOTMYXwuXuui82ch47qW5XKTCjLoQ5+TM+gM/EYQ6A+vWw7dkBAAJcekaohmAItymVJ/v234YzavDTna0vRDoeZZyf/v2dLPuz25Rc4ZsUX3M2L9Bne11xI2KWnoCCgsrfwCZ5d6NFRnH1Ld/jO67mVJj1SamexITk9WjwVhSetyjfw3o7Tw4eQ/4WaTe1UcQa2MMZgxO1HZEuNjR3k4MORJATbfwdxsBxhvFn2MXTvDSQMwS6LPKmR3nz3JHhrDgRj2KangIJdKLpFCP+lIK6SPghKJvNM7jXyFLqiucpYMeMqyjyOOn5cy2ANIRiO5HppjaCs36n8xM/qlaSrmbaK7jGMkD1CokK8VHpf4rMkF5VSLhhYHwXZXdSAq2GPzxY+0AOi80Hi3+6HF8e0rQ1nYGOTWrplrzeW/vCRUEiBAq+3SdKeoA==","hh":"tEtDmj9NO12CQujtDAf6znTptSdPuOFPcXitzqFb5a4=","i1":"MmEhOUHiGsRJZjij1zCWjQ==","i2":"ukacN6jPVDqrPE8x7lSyqg==","m":"P76F7Om9PMh31Wrui42tfu7uvpVw9MsguMLcoozrHaY=","ra":"TW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzg3LjAuNDI4MC4xNDEgU2FmYXJpLzUzNy4zNg==","rm":"R0VU","ru":"aHR0cHM6Ly93d3cuc25lYWtlcnNuc3R1ZmYuY29tLw==","t":"MTYxODM3NDcyNy4zODkwMDA=","uh":"sfgiHR/bGP0DFXSp4KJXpJnmQXuNHfWxwUhpX9SCMz4="},"chC":0,"chCAS":0,"chLog":{"0":{"start":1618374728142},"c":1},"chReq":"non-interactive","cvId":"2","oV":1}`
	var d = `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	fmt.Println(Compress(dataStr, d))
}

//goland:noinspection ALL
func Test_DeCompress(t *testing.T) {
	var dataStr = largeTestString
	var d = `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	fmt.Println(Decompress(dataStr, d))
}

//goland:noinspection ALL
func Test_DeCompressToBase641(t *testing.T) {
	var dataStr = `lDXyqyYytyxyfJLedLH0LWLzy8V+Oiy0voL=`
	str := Compress(dataStr, "")
	res, err := Decompress(str, "")
	fmt.Println(res == dataStr, err)
}

//goland:noinspection ALL
func Test_DeCompressToBase642(t *testing.T) {
	var dataStr = `lDXyqyYytyxyfJLedLH0LWLzy8V+Oiy0voL=`
	var d = `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	str := Compress(dataStr, d)
	res, err := Decompress(str, d)
	fmt.Println(res == dataStr, res, err)
}
func TestCompressFromString(t *testing.T) {
	type args struct {
		uncompressed string
		chars        string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress(tt.args.uncompressed, tt.args.chars); got != tt.want {
				t.Errorf("CompressFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Integrity(t *testing.T) {
	startString := "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives."
	compressed := Compress(startString, _defaultKeyStrBase64)
	decompress, err := Decompress(compressed, _defaultKeyStrBase64)
	if err != nil {
		t.Fatal(err)
	}
	if decompress != startString {
		t.Fatalf("Decompress() = %s; want %s", decompress, startString)
	}
	shouldBe := "OIewBAlgzmCGYDsCmB3MAbWCDmBXW2SAdGAILoAuAFiLtlZBWAEYgBObIKMEAJkrBgAzTgFswSAB7QKEHBix4CSKABpGYKoLC4EuKPnRgADp2NI2slWGqwmo2AGskEoUKQBjWQDcXoE5zYbLCiMLwQbhZICExyYB5awV4WYCIg4qYgQSEwKGwQFBTRkAiMMGxImD4qREA==="
	shouldBeDecompressed, err := Decompress(shouldBe, _defaultKeyStrBase64)
	if err != nil {
		t.Fatal(err)
	}
	if shouldBeDecompressed != startString {
		t.Fatalf("Decompress() = %s; want %s", shouldBeDecompressed, startString)
	}
}

//goland:noinspection ALL
func BenchmarkCompressFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compress("1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer1233333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333eeeeeeeeeeeeeeeeee2e23e2qwerqwerqwerqwerqwerqweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerqwerqwer", _defaultKeyStrBase64)
	}
}

func BenchmarkCompressFromString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compress("3123123", _defaultKeyStrBase64)
	}
}

//goland:noinspection ALL
func BenchmarkCompressToBase64(b *testing.B) {
	str := `{"cNounce":"36728","cRq":{"d":"oraaF6KB66uNJPX9UPeeiXActri35Lwu5mv+YI1bqHV1rotfmyjcn9iBX5LZto7WyVo6iQlpqpGTRZ3WeOFTCGsJBdry5WLcXcxGqhuJixjpDcv2W9h4geRwf+agasGzInsjMYBCg6GnzYmxxpdVp4U16pzjPJoikxflHq0XG+GkxMnMIAPWgpKeyNAy0Aj7WdTo0pmP3Cn6B4NrNFXhzekJG7IBR8KkfUXlteieFKS3k0WRH09OsmA+GxMz8Pag0KiXN8+tIs19PqDMHNoeJBkLMtmLKTzExtxP2wgaVjWh1UPUim7M7yLGkmdv5bU3Zuz64ASvr8IEWmiZVJYenz9hFA4MhYtNRZVyR1AYx41Fb0hvRcrQG5M5vsZXymJykAYWITIthMw6rxHnTI+Vbx6E46Yu9/QcJ4xT5MWmW8zXRIM+jp5fow4/3335tRsMSdyIL4bE+ulw52cTVXrpNqhqCatH/wHHyUshuhjsFPMMxs6WoRFW4ceOX8vi7R8TtprsOtSmWg1OK7hYrl99ZXfIcbcJwRjRzfdsYrTN5FrpCqdk2DjD5UXG7EDcI+LodnsK4PyDAMzzAFyzubmQFICT8spvRrnUoDqMzCq7ViWt2JpwqDt5gmGMHZJg5VWkGYG1EjwkLr8FZ9gh2ezHeg7e0siiPDf6AbkvU4KCmR1Ht7qgOe0l+LDM1t/m7o1tUwkoY3J1poq0U2juD+Itcg==","hh":"QSlXmfS2NeZCSV7a6l5l5OS3E6mq9iZtckO4Lj4aWug=","i1":"bSoctx5zCBRkkKOF5bXUCQ==","i2":"XPiyF0pHvrNO7OakWEMoOg==","m":"eqV80Sw4dkgMfJScOzRO4cIGos2UPT/g3Fn2WVKKC5Y=","ra":"TW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzg3LjAuNDI4MC4xNDEgU2FmYXJpLzUzNy4zNg==","rm":"R0VU","ru":"aHR0cHM6Ly93d3cud29vZHdvb2QuY29tLw==","t":"MTYxODM4NDAzMi4wNzAwMDA=","uh":"sfgiHR/bGP0DFXSp4KJXpJnmQXuNHfWxwUhpX9SCMz4="},"chC":0,"chCAS":0,"chLog":{"0":{"start":1618384032901},"c":1},"chReq":"non-interactive","cvId":"2","oV":1}`
	for i := 0; i < b.N; i++ {
		Compress(str, `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`)
	}
}

//goland:noinspection ALL
func BenchmarkDeCompressToBase64(b *testing.B) {
	str := largeTestString
	for i := 0; i < b.N; i++ {
		_, _ = Decompress(str, `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`)
	}
}

//goland:noinspection ALL
func BenchmarkDeCompressCon(b *testing.B) {
	str := largeTestString
	mapper := `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			_, _ = Decompress(str, mapper)
			wg.Done()
		}()
	}
	wg.Wait()
}

//goland:noinspection ALL
func BenchmarkDecompressFromBase64old(b *testing.B) {
	//goland:noinspection ALL
	str := largeTestString
	mapper := `Lz0-6uFWXOjE8lMIok1S49ZH+5TGYfsByPdhQAJK7ebxUpRriqVn2gaCDmcN3tv$w`
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			_, _ = Decompress(str, mapper)
			wg.Done()
		}()
	}
	wg.Wait()
}
