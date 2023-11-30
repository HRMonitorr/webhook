package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HRMonitorr/githubwrapper"
	"github.com/HRMonitorr/webhook/functions"
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func PostBalasan(w http.ResponseWriter, r *http.Request, githubToken string) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	link := "https://medium.com/@rofinafiisr/whatsauth-free-2fa-otp-notif-whatsapp-gateway-api-gratis-f540249cd050"
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if msg.Message == "loc" || msg.Message == "Loc" || msg.Message == "lokasi" || msg.LiveLoc {
			location, err := functions.ReverseGeocode(msg.Latitude, msg.Longitude)
			if err != nil {
				// Handle the error (e.g., log it) and set a default location name
				location = "Unknown Location"
			}

			reply := fmt.Sprintf("Hai hai haiii kamu pasti lagi di %s "+
				"\n Koordinatenya : %s - %s"+
				"\n Cara Penggunaan WhatsAuth Ada di link dibawah ini"+
				"yaa kak %s\n", location,
				strconv.Itoa(int(msg.Longitude)), strconv.Itoa(int(msg.Latitude)), link)
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: reply,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if msg.Message == "Babi" || msg.Message == "Anjing" || msg.Message == "goblok" {
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: fmt.Sprintf("Ihh kakak %s kamu kasar bangett, aku jadi takut tauuu", msg.Alias_name),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if msg.Message == "cantik" || msg.Message == "ganteng" || msg.Message == "cakep" {
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: fmt.Sprintf("makasiihh kakak %s kamu jugaa cakep kooo", msg.Alias_name),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if strings.Contains(msg.Message, "Push file") || msg.Filedata != "" {
			//push file ke repo namarepo pesan test
			messages := strings.Split(msg.Message, " ")
			reponame := messages[4]
			pesan := strings.Split(msg.Message, "pesan")
			datas := githubwrapper.PushRepositories{
				Context:       context.Background(),
				PersonalToken: githubToken,
				Reponame:      reponame,
				OwnerName:     "HRMonitorr",
				Path:          msg.Filedata,
				Username:      "rofinafiin",
				Email:         "rofinafiisr@gmail.com",
				Message:       pesan[1],
				Branch:        "master",
			}
			responses, err := githubwrapper.UploadFileToRepository(datas)
			if err != nil {
				dt := &wa.TextMessage{
					To:       msg.Phone_number,
					IsGroup:  false,
					Messages: fmt.Sprintf("aduh kak %s kayaknya kakak gabisa upload file nih soalnya %s", msg.Alias_name, err.Error()),
				}
				resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
			}
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: fmt.Sprintf("Hai hai kakakkk sudah berhasil upload nih ke reponyaa %s", responses),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		} else {
			randm := []string{
				"Hai Hai Haiii kamuuuui " + msg.Alias_name + "\nrofinya lagi gaadaa \n aku giseuubott salam kenall yaaaa \n Cara penggunaan WhatsAuth ada di link berikut ini ya kak...\n" + link,
				"IHHH jangan SPAAM berisik tau giseu lagi tidur",
				"Kamu ganteng tau",
				"Ihhh kamu cantik banget",
				"bro, mending beliin aku nasgor",
				"Jangan galak galak dong kak, aku takut tauu",
				"Mawar Indah hanya akan muncul dipagi hari, MAKANYA BANGUN PAGI KAK",
				"Cihuyyyy hari ini giseuu bahagiaaa banget",
				"Bercandyaaa berrcandyaaaa",
			}
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: functions.GetRandomString(randm),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
