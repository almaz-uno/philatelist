package google

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"bitbucket.org/CuredPlumbum/philatelist/imagesearch"

	"github.com/kr/pretty"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	kovrovKey  = "AIzaSyBZeYbJ6pMNUy-VdLVxnhBCwYcWxSrZZAE"
	testAPIKey = kovrovKey
)

func TestTypes(t *testing.T) {

	t.Run("text-search-google-sample", func(t *testing.T) {

		textSearchResult := `{
   "html_attributions" : [],
   "results" : [
      {
         "geometry" : {
            "location" : {
               "lat" : -33.870775,
               "lng" : 151.199025
            }
         },
         "icon" : "http://maps.gstatic.com/mapfiles/place_api/icons/travel_agent-71.png",
         "id" : "21a0b251c9b8392186142c798263e289fe45b4aa",
         "name" : "Rhythmboat Cruises",
         "opening_hours" : {
            "open_now" : true
         },
         "photos" : [
            {
               "height" : 270,
               "html_attributions" : [],
               "photo_reference" : "CnRnAAAAF-LjFR1ZV93eawe1cU_3QNMCNmaGkowY7CnOf-kcNmPhNnPEG9W979jOuJJ1sGr75rhD5hqKzjD8vbMbSsRnq_Ni3ZIGfY6hKWmsOf3qHKJInkm4h55lzvLAXJVc-Rr4kI9O1tmIblblUpg2oqoq8RIQRMQJhFsTr5s9haxQ07EQHxoUO0ICubVFGYfJiMUPor1GnIWb5i8",
               "width" : 519
            }
         ],
         "place_id" : "ChIJyWEHuEmuEmsRm9hTkapTCrk",
         "scope" : "GOOGLE",
         "alt_ids" : [
            {
               "place_id" : "D9iJyWEHuEmuEmsRm9hTkapTCrk",
               "scope" : "APP"
            }
         ],
         "reference" : "CoQBdQAAAFSiijw5-cAV68xdf2O18pKIZ0seJh03u9h9wk_lEdG-cP1dWvp_QGS4SNCBMk_fB06YRsfMrNkINtPez22p5lRIlj5ty_HmcNwcl6GZXbD2RdXsVfLYlQwnZQcnu7ihkjZp_2gk1-fWXql3GQ8-1BEGwgCxG-eaSnIJIBPuIpihEhAY1WYdxPvOWsPnb2-nGb6QGhTipN0lgaLpQTnkcMeAIEvCsSa0Ww",
         "types" : [ "travel_agency", "restaurant", "food", "establishment" ],
         "vicinity" : "Pyrmont Bay Wharf Darling Dr, Sydney"
      },
      {
         "geometry" : {
            "location" : {
               "lat" : -33.866891,
               "lng" : 151.200814
            }
         },
         "icon" : "http://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
         "id" : "45a27fd8d56c56dc62afc9b49e1d850440d5c403",
         "name" : "Private Charter Sydney Habour Cruise",
         "photos" : [
            {
               "height" : 426,
               "html_attributions" : [],
               "photo_reference" : "CnRnAAAAL3n0Zu3U6fseyPl8URGKD49aGB2Wka7CKDZfamoGX2ZTLMBYgTUshjr-MXc0_O2BbvlUAZWtQTBHUVZ-5Sxb1-P-VX2Fx0sZF87q-9vUt19VDwQQmAX_mjQe7UWmU5lJGCOXSgxp2fu1b5VR_PF31RIQTKZLfqm8TA1eynnN4M1XShoU8adzJCcOWK0er14h8SqOIDZctvU",
               "width" : 640
            }
         ],
         "place_id" : "ChIJqwS6fjiuEmsRJAMiOY9MSms",
         "scope" : "GOOGLE",
         "reference" : "CpQBhgAAAFN27qR_t5oSDKPUzjQIeQa3lrRpFTm5alW3ZYbMFm8k10ETbISfK9S1nwcJVfrP-bjra7NSPuhaRulxoonSPQklDyB-xGvcJncq6qDXIUQ3hlI-bx4AxYckAOX74LkupHq7bcaREgrSBE-U6GbA1C3U7I-HnweO4IPtztSEcgW09y03v1hgHzL8xSDElmkQtRIQzLbyBfj3e0FhJzABXjM2QBoUE2EnL-DzWrzpgmMEulUBLGrtu2Y",
         "types" : [ "restaurant", "food", "establishment" ],
         "vicinity" : "Australia"
      },
      {
         "geometry" : {
            "location" : {
               "lat" : -33.870943,
               "lng" : 151.190311
            }
         },
         "icon" : "http://maps.gstatic.com/mapfiles/place_api/icons/restaurant-71.png",
         "id" : "30bee58f819b6c47bd24151802f25ecf11df8943",
         "name" : "Bucks Party Cruise",
         "opening_hours" : {
            "open_now" : true
         },
         "photos" : [
            {
               "height" : 600,
               "html_attributions" : [],
               "photo_reference" : "CnRnAAAA48AX5MsHIMiuipON_Lgh97hPiYDFkxx_vnaZQMOcvcQwYN92o33t5RwjRpOue5R47AjfMltntoz71hto40zqo7vFyxhDuuqhAChKGRQ5mdO5jv5CKWlzi182PICiOb37PiBtiFt7lSLe1SedoyrD-xIQD8xqSOaejWejYHCN4Ye2XBoUT3q2IXJQpMkmffJiBNftv8QSwF4",
               "width" : 800
            }
         ],
         "place_id" : "ChIJLfySpTOuEmsRsc_JfJtljdc",
         "scope" : "GOOGLE",
         "reference" : "CoQBdQAAANQSThnTekt-UokiTiX3oUFT6YDfdQJIG0ljlQnkLfWefcKmjxax0xmUpWjmpWdOsScl9zSyBNImmrTO9AE9DnWTdQ2hY7n-OOU4UgCfX7U0TE1Vf7jyODRISbK-u86TBJij0b2i7oUWq2bGr0cQSj8CV97U5q8SJR3AFDYi3ogqEhCMXjNLR1k8fiXTkG2BxGJmGhTqwE8C4grdjvJ0w5UsAVoOH7v8HQ",
         "types" : [ "restaurant", "food", "establishment" ],
         "vicinity" : "37 Bank St, Pyrmont"
      },
      {
         "geometry" : {
            "location" : {
               "lat" : -33.867591,
               "lng" : 151.201196
            }
         },
         "icon" : "http://maps.gstatic.com/mapfiles/place_api/icons/travel_agent-71.png",
         "id" : "a97f9fb468bcd26b68a23072a55af82d4b325e0d",
         "name" : "Australian Cruise Group",
         "opening_hours" : {
            "open_now" : true
         },
         "photos" : [
            {
               "height" : 242,
               "html_attributions" : [],
               "photo_reference" : "CnRnAAAABjeoPQ7NUU3pDitV4Vs0BgP1FLhf_iCgStUZUr4ZuNqQnc5k43jbvjKC2hTGM8SrmdJYyOyxRO3D2yutoJwVC4Vp_dzckkjG35L6LfMm5sjrOr6uyOtr2PNCp1xQylx6vhdcpW8yZjBZCvVsjNajLBIQ-z4ttAMIc8EjEZV7LsoFgRoU6OrqxvKCnkJGb9F16W57iIV4LuM",
               "width" : 200
            }
         ],
         "place_id" : "ChIJrTLr-GyuEmsRBfy61i59si0",
         "scope" : "GOOGLE",
         "reference" : "CoQBeQAAAFvf12y8veSQMdIMmAXQmus1zqkgKQ-O2KEX0Kr47rIRTy6HNsyosVl0CjvEBulIu_cujrSOgICdcxNioFDHtAxXBhqeR-8xXtm52Bp0lVwnO3LzLFY3jeo8WrsyIwNE1kQlGuWA4xklpOknHJuRXSQJVheRlYijOHSgsBQ35mOcEhC5IpbpqCMe82yR136087wZGhSziPEbooYkHLn9e5njOTuBprcfVw",
         "types" : [ "travel_agency", "restaurant", "food", "establishment" ],
         "vicinity" : "32 The Promenade, King Street Wharf 5, Sydney"
      }
   ],
   "status" : "OK"
}`

		res := new(SearchResponse)

		err := json.Unmarshal([]byte(textSearchResult), res)
		require.NoError(t, err)

		assert.Len(t, res.Results, 4)

		assert.Equal(t, "OK", res.Status)

		// it's enough to check only the first result
		r := res.Results[0]
		assert.Len(t, r.Photos, 1)
		assert.Equal(t, 270, r.Photos[0].Height)
		assert.Equal(t, 519, r.Photos[0].Width)
		assert.Equal(t, "CnRnAAAAF-LjFR1ZV93eawe1cU_3QNMCNmaGkowY7CnOf-kcNmPhNnPEG9W979jOuJJ1sGr75rhD5hqKzjD8vbMbSsRnq_Ni3ZIGfY6hKWmsOf3qHKJInkm4h55lzvLAXJVc-Rr4kI9O1tmIblblUpg2oqoq8RIQRMQJhFsTr5s9haxQ07EQHxoUO0ICubVFGYfJiMUPor1GnIWb5i8", r.Photos[0].PhotoReference)

	})

	t.Run("text-search-nt", func(t *testing.T) {
		textSearchResult := `{
   "html_attributions" : [],
   "results" : [
      {
         "formatted_address" : "6, 5, 4, ulitsa Novotushinskaya, 3, 2 Москва, Moscow, Russia, 143441",
         "geometry" : {
            "location" : {
               "lat" : 55.8697078,
               "lng" : 37.4007989
            },
            "viewport" : {
               "northeast" : {
                  "lat" : 55.8710567802915,
                  "lng" : 37.40214788029149
               },
               "southwest" : {
                  "lat" : 55.8683588197085,
                  "lng" : 37.39944991970849
               }
            }
         },
         "icon" : "https://maps.gstatic.com/mapfiles/place_api/icons/generic_business-71.png",
         "id" : "d513b5f4fb5c2a6e7bc4836d22bb9e3a8ef13090",
         "name" : "UP-Kvartal Novoye Tushino",
         "photos" : [
            {
               "height" : 960,
               "html_attributions" : [
                  "\u003ca href=\"https://maps.google.com/maps/contrib/100414973130479106069/photos\"\u003eАнастасия Хачатурян\u003c/a\u003e"
               ],
               "photo_reference" : "CmRaAAAA8fFng-XXjD6AdLGk9sa03g2XKb-_mFt4TSWMk_NlCEAMfpz-wCiBeAwk2y8AfzU4uhndnbzqudSE8XJHSszYNHi2e5bFfFocb6ctREcWsz9hKUJZ5Qu_sTr50QTCCQlMEhDi17EY-lQGBLOShNlJMV4KGhRaVh_UouT2eRVhKMmESGpSAXaZ2g",
               "width" : 1280
            }
         ],
         "place_id" : "ChIJhSOnan5HtUYR5dkSIiDaj-U",
         "rating" : 5,
         "reference" : "CmRSAAAAzk-sg_KRH2482UeB0pBvZa03yAThs1RfhbsjBJjMByPSnJKmjjHB22J0nl1Fc-eATgsPF5D3cBarkHcivmPmQBi2sYijANncyn2PkAkRV_EBrR36q31PnPLyhXNtKEgaEhATisS4G52zxIzIcU1eVKrBGhSY3ZwIyWGR2homR540P-cCPuXQjg",
         "types" : [ "point_of_interest", "establishment" ]
      }
   ],
   "status" : "OK"
}`

		res := new(SearchResponse)

		err := json.Unmarshal([]byte(textSearchResult), res)
		require.NoError(t, err)

		assert.Len(t, res.Results, 1)

		assert.Equal(t, "OK", res.Status)

		// it's enough to check only the first result
		r := res.Results[0]
		assert.Len(t, r.Photos, 1)
		assert.Equal(t, 960, r.Photos[0].Height)
		assert.Equal(t, 1280, r.Photos[0].Width)
		assert.Equal(t, "CmRaAAAA8fFng-XXjD6AdLGk9sa03g2XKb-_mFt4TSWMk_NlCEAMfpz-wCiBeAwk2y8AfzU4uhndnbzqudSE8XJHSszYNHi2e5bFfFocb6ctREcWsz9hKUJZ5Qu_sTr50QTCCQlMEhDi17EY-lQGBLOShNlJMV4KGhRaVh_UouT2eRVhKMmESGpSAXaZ2g", r.Photos[0].PhotoReference)

		assert.Equal(t, "6, 5, 4, ulitsa Novotushinskaya, 3, 2 Москва, Moscow, Russia, 143441", r.FormattedAddress)
		assert.Equal(t, GeoPoint{55.8697078, 37.4007989}, r.Geometry.Location)
		assert.Equal(t, &GeoPoint{55.8710567802915, 37.40214788029149}, r.Geometry.Viewport.Northeast)
		assert.Equal(t, &GeoPoint{55.8683588197085, 37.39944991970849}, r.Geometry.Viewport.Southwest)
	})

	t.Run("details-auckland", func(t *testing.T) {
		answer := `{
   "html_attributions" : [],
   "result" : {
      "address_components" : [
         {
            "long_name" : "Auckland",
            "short_name" : "Auckland",
            "types" : [ "locality", "political" ]
         },
         {
            "long_name" : "Auckland",
            "short_name" : "Auckland",
            "types" : [ "administrative_area_level_2", "political" ]
         },
         {
            "long_name" : "Auckland",
            "short_name" : "Auckland",
            "types" : [ "administrative_area_level_1", "political" ]
         },
         {
            "long_name" : "New Zealand",
            "short_name" : "NZ",
            "types" : [ "country", "political" ]
         }
      ],
      "adr_address" : "\u003cspan class=\"locality\"\u003eAuckland\u003c/span\u003e, \u003cspan class=\"country-name\"\u003eNew Zealand\u003c/span\u003e",
      "formatted_address" : "Auckland, New Zealand",
      "geometry" : {
         "location" : {
            "lat" : -36.8484597,
            "lng" : 174.7633315
         },
         "viewport" : {
            "northeast" : {
               "lat" : -36.660571,
               "lng" : 175.2871371
            },
            "southwest" : {
               "lat" : -37.0654751,
               "lng" : 174.4438016
            }
         }
      },
      "icon" : "https://maps.gstatic.com/mapfiles/place_api/icons/geocode-71.png",
      "id" : "088418ddc17fef2513462d92dbee1355929b35ed",
      "name" : "Auckland",
      "photos" : [
         {
            "height" : 3888,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/118418761811197795026/photos\"\u003eLeen Groenendijk\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAzSIOIi67C1GF9iqNGTuH-DrMZN8dKaBttDFv8GLyACqbkN0s4Cs0joOwb9eHSquGSodOcJcob7vFQ9cMfKONNQ1dyWBnphBBTDE02PCrxwpzIHVhYusswVqDPxa93b6rEhD1WnHLc8WLujat4EZBKmgYGhS2s3GCNbnQrIuu9UZgPZd55T5K4g",
            "width" : 5184
         },
         {
            "height" : 2823,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/105593113606540041698/photos\"\u003eSonia Alex\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAydkHpth6DcqevGwgneRipzCuelPxLPTKTOXGk1NsUN5Gqzc_Ktqaps9Dowx_yHwunsRp5XnYU6KBBZYa5MRlIzSeg7eTUh5bFj7Y_90yiHu3wdEUJilIjbtyEcXTEwMjEhDe1B2VqdQxE0G0hc1rOiKcGhQ5NQgYGtRZEE3c9MOPNE_vMJzOWA",
            "width" : 3816
         },
         {
            "height" : 1067,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/100578995542168242521/photos\"\u003eAndrew Burns\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAA70KdNXt1cSFiHr6kALiNecBSs2OEYS1K1O-yWjbIysC7E1yXT4udg9CxYY_8crRtL3kr9SPu-h9EbVCG2IRNqPmnc9AWfHiafFGUNaM3Qtz_EL05GTZS4eb9AK3LU_8TEhAmQZemBc67WuIw3rgmP68rGhSOfY17rKLWIlVVZWxJi4YFpuSWGQ",
            "width" : 1600
         },
         {
            "height" : 2988,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/104449230804367883642/photos\"\u003eMichal Panek\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAV4nJHjPvuLlEY9BszlEQBSmr1rUEqVs49iHiouzxGQBPOtYXidxgEnmpOOAypjPBb5IGLCCJR-HxxW_JCNnLEpeoUKz8lwz3YIE_RVm8UN_DT7vDH2-8nbdXLtLe_4VLEhDAtO_2CcFN0DQM7qJAJSKiGhQ52UJCb_Br4nAcuGAw7iOIAnz0Og",
            "width" : 5312
         },
         {
            "height" : 2250,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/116101852626406728300/photos\"\u003eRyan Resurreccion\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAWEefvIKfENV7CDJOdSsZM3_w9esseXp3qr_PYpSaIQ8qCAt5LQIK427ovA8VZl5p3sIDI9Pp4uA0VLc_9unYxi0cJfJAGN2SHKsOLJnffOkL12QrBn-WwJUk3eikv-QyEhC26MY_WkZRnHhslVsRwAGTGhSGdACqNT8qw889DPiKZ8yUtS9GwQ",
            "width" : 4000
         },
         {
            "height" : 2160,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/103332394033221586739/photos\"\u003ePathum Chathuranga\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAi18ugdM24_hlgHBJk6rEEAEZK6lrU71A1OET_erTgUMGA8rPNWf6FnkuVdh906mgi95uK67gh0Mrs4G9yEkwcho61nNT4eqS0EEue-qt_Bg0TaQhHYJJuyJImtBI7hMTEhABnl_vUiucAu_TkF1N1UpKGhQrVR109UGufNClBMI59hd7EZl7cw",
            "width" : 3840
         },
         {
            "height" : 3341,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/102080148977768116246/photos\"\u003eDillon Mahoney\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAfFuOHbHRiethpC-zevF_xVhzUiaob0ggAB4oOmRabU-2gc-lA9q8yBPPdisy15MCZb7XlOe8vAGjaeIls0HDSTrKSI0txmdgTxT2Le1_dx2e1pdUPOJItZ-z7U382-FgEhDMrZkOR8N7REvKABN2nmhbGhRValUjPdc9fZAJIxVhFuZfxC1hjg",
            "width" : 5012
         },
         {
            "height" : 1152,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/103146681571927272099/photos\"\u003eChrissy Mitchell\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAz7EzePIX7otreVvf_D15tRkcHtSa4yTYlrMi7gPIbSRdZactKSuQwrsBP7ewIHLQ8NTyGQB1b1H-xSC4m82iVSdYlwHWbnObzOmuJdTr_98W7001X0l1nmnJgybCNAUREhBDck8o5YivIFVkQyQP2RuRGhR4JThdad9OEnRuTi0zArilLyc2pw",
            "width" : 2048
         },
         {
            "height" : 1836,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/110468464466037605553/photos\"\u003eIan Bryant\u003c/a\u003e"
            ],
            "photo_reference" : "CmRZAAAABfV2IGOul0trOBT04U-N3wVJbB4h8C9CPWTZXMjhD27N0-dwj99eTkEMdegLndtlbbLHXBWAZ0t3EmrknoBKUhRzcroToS_l5ZPTKqssxlqwS1CwuHqEfUG7esJ01WrsEhDGi-n5778rlrj-K0tCXK4BGhSlbbaCKxA5ro0Z9m0dYX1EbNzXfA",
            "width" : 3264
         },
         {
            "height" : 3888,
            "html_attributions" : [
               "\u003ca href=\"https://maps.google.com/maps/contrib/118418761811197795026/photos\"\u003eLeen Groenendijk\u003c/a\u003e"
            ],
            "photo_reference" : "CmRaAAAAn_M2khG1ZRrb9PtDL3evPfcF0MJtIbtCq-poUFhKhGy4Zar9T0_9VPgh2fkBe4QtBXvS7aM9WAJZX330Ud2c9_gtvDzSkWg8A0uun1SpnVDWScW-34Lh5SEc6rh3ZXa0EhDb7IMlI-bTY_2asgdGpj_1GhRbe01nXWbrh6PH_MootP7iUAzcHg",
            "width" : 5184
         }
      ],
      "place_id" : "ChIJ--acWvtHDW0RF5miQ2HvAAU",
      "reference" : "CmRbAAAA_LUPT9OG-PQDHw5yItyrzVx3gVcwrVxWm6Rf-AaZmfEz-Fc0CRiGRiwGUSPwrzatn-2n1p9tfnwqD4P1cfzJK-LNFNb2bz05yM-GH0pJv9GDY5MNcQJ3GbQ1J1TdRQ00EhD4soZprqJQfxlJj2ag8vGlGhRpbqAWd2HE-LUOJYECzapaNuU3LA",
      "scope" : "GOOGLE",
      "types" : [ "locality", "political" ],
      "url" : "https://maps.google.com/?q=Auckland,+New+Zealand&ftid=0x6d0d47fb5a9ce6fb:0x500ef6143a29917",
      "utc_offset" : 720,
      "vicinity" : "Auckland"
   },
   "status" : "OK"
}`
		res := new(DetailsResponse)

		err := json.Unmarshal([]byte(answer), res)
		require.NoError(t, err)

		assert.Equal(t, "OK", res.Status)

		assert.Equal(t, GeoPoint{Lat: -36.84846, Lng: 174.76334}, res.Result.Geometry.Location)
		assert.Equal(t, &GeoPoint{Lat: -36.660572, Lng: 175.28714}, res.Result.Geometry.Viewport.Northeast)
		assert.Equal(t, &GeoPoint{Lat: -37.065475, Lng: 174.4438}, res.Result.Geometry.Viewport.Southwest)

		assert.Equal(t, "Auckland, New Zealand", res.Result.FormattedAddress)

		assert.Len(t, res.Result.Photos, 10)

		// it's enough to check only the first photo
		p := res.Result.Photos[0]
		assert.Len(t, p.HTMLAttributions, 1)
		assert.Equal(t, "<a href=\"https://maps.google.com/maps/contrib/118418761811197795026/photos\">Leen Groenendijk</a>", p.HTMLAttributions[0])
		assert.Equal(t, 3888, p.Height)
		assert.Equal(t, 5184, p.Width)
		assert.Equal(t, "CmRaAAAAzSIOIi67C1GF9iqNGTuH-DrMZN8dKaBttDFv8GLyACqbkN0s4Cs0joOwb9eHSquGSodOcJcob7vFQ9cMfKONNQ1dyWBnphBBTDE02PCrxwpzIHVhYusswVqDPxa93b6rEhD1WnHLc8WLujat4EZBKmgYGhS2s3GCNbnQrIuu9UZgPZd55T5K4g", p.PhotoReference)

	})

}

func TestGoogleApi(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	t.Run("details-auckland", func(t *testing.T) {
		// Auckland
		const auckland = "ChIJ--acWvtHDW0RF5miQ2HvAAU"

		subj := New(testAPIKey)
		subj.Timeout = time.Minute

		res, err := subj.Details(auckland)
		require.NoError(t, err)
		t.Log(pretty.Sprint("auckland details:", res))

		for _, ph := range res.Result.Photos {
			t.Log(subj.Photo(ph))
		}

	})

	t.Run("details-by-placeid-custom", func(t *testing.T) {
		// ulitsa Krzhizhanovskogo, Moscow
		const krzhizhanovskogoStreet = "ChIJIY8uUsBMtUYRdCGdlBVZxlU" // nolint: megacheck
		// sanatoriya Podmoskovie
		const sanatoriyaPodmoskovie = "ChIJVVVVVdigSkERIUYqip8z7OU" // nolint: megacheck
		// Novotushinskaya ulitsa
		const novotushinskaya = "ChIJhSOnan5HtUYR5dkSIiDaj-U"

		subj := New(testAPIKey)
		subj.Timeout = time.Minute

		res, err := subj.Details(novotushinskaya)
		require.NoError(t, err)
		t.Log(pretty.Sprint("novotushinskaya details:", res))

		for _, ph := range res.Result.Photos {
			t.Log(subj.Photo(ph))
		}

	})

	t.Run("photos-by-textsearch-auckland", func(t *testing.T) {

		subj := New(testAPIKey)
		subj.Timeout = time.Minute

		query := "New Zeland, Auckland"

		res, err := subj.TextSearch(query)
		assert.NoError(t, err)

		assert.NotNil(t, res)
		require.Len(t, res.Results, 1)
		assert.Equal(t, "Auckland, New Zealand", res.Results[0].FormattedAddress)
		t.Log(pretty.Sprint("res", res))

		require.True(t, len(res.Results) > 0 && len(res.Results[0].Photos) > 0, "There is no photograph for the place")

		t.Logf("%v?key=%v&photoreference=%v&maxwidth=%v", photoServiceURL, testAPIKey, res.Results[0].Photos[0].PhotoReference, res.Results[0].Photos[0].Width)

	})

	t.Run("photos-by-textsearch-custom", func(t *testing.T) {

		subj := New(testAPIKey)
		subj.Timeout = time.Minute

		query := "Новое Тушино, Путилково"

		res, err := subj.TextSearch(query)
		assert.NoError(t, err)

		assert.NotNil(t, res)

		t.Log(pretty.Sprint("res", res))

		require.True(t, len(res.Results) > 0 && len(res.Results[0].Photos) > 0, "There is no photograph for `"+query+"`")

		expectedPhotoURL := fmt.Sprintf("%v?key=%v&photoreference=%v&maxwidth=%v", photoServiceURL, testAPIKey, res.Results[0].Photos[0].PhotoReference, res.Results[0].Photos[0].Width)
		assert.Equal(t, expectedPhotoURL, subj.Photo(res.Results[0].Photos[0]))

		t.Logf(subj.Photo(res.Results[0].Photos[0]))

	})
}

func TestSearcher(t *testing.T) {

	t.Run("interface", func(t *testing.T) {
		assert.Implements(t, (*imagesearch.Searcher)(nil), new(API))
	})

	t.Run("SearchByQuery-auckland", func(t *testing.T) {
		subj := New(testAPIKey)
		urls, err := subj.SearchByQuery("Auckland, New Zeland")
		require.NoError(t, err)

		assert.Len(t, urls, 10)

		for _, u := range urls {
			t.Log(u)
		}

	})
	t.Run("SearchByQuery-domodedovo", func(t *testing.T) {
		subj := New(testAPIKey)
		urls, err := subj.SearchByQuery("Domodedovo town, Russia")
		require.NoError(t, err)

		assert.Len(t, urls, 10)

		for _, u := range urls {
			t.Log(u)
		}

	})
}
