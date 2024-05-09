package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Character struct {
	Name                 string    `json:"name"`
	Race                 string    `json:"race"`
	Class                string    `json:"class"`
	ActiveSpecName       string    `json:"active_spec_name"`
	ActiveSpecRole       string    `json:"active_spec_role"`
	Gender               string    `json:"gender"`
	Faction              string    `json:"faction"`
	AchievementPoints    int       `json:"achievement_points"`
	HonorableKills       int       `json:"honorable_kills"`
	ThumbnailURL         string    `json:"thumbnail_url"`
	Region               string    `json:"region"`
	Realm                string    `json:"realm"`
	LastCrawledAt        time.Time `json:"last_crawled_at"`
	ProfileURL           string    `json:"profile_url"`
	ProfileBanner        string    `json:"profile_banner"`
	MythicPlusRecentRuns []struct {
		Dungeon             string    `json:"dungeon"`
		ShortName           string    `json:"short_name"`
		MythicLevel         int       `json:"mythic_level"`
		CompletedAt         time.Time `json:"completed_at"`
		ClearTimeMs         int       `json:"clear_time_ms"`
		ParTimeMs           int       `json:"par_time_ms"`
		NumKeystoneUpgrades int       `json:"num_keystone_upgrades"`
		MapChallengeModeID  int       `json:"map_challenge_mode_id"`
		ZoneID              int       `json:"zone_id"`
		Score               float64   `json:"score"`
		Affixes             []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
			WowheadURL  string `json:"wowhead_url"`
		} `json:"affixes"`
		URL string `json:"url"`
	} `json:"mythic_plus_recent_runs"`
}

func main() {
	resp, err := http.Get("https://raider.io/api/v1/characters/profile?region=eu&realm=Shadowsong&name=Brenith&fields=mythic_plus_recent_runs")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	var character Character
	charByte := []byte(body)

	if err := json.Unmarshal(charByte, &character); err != nil {
		panic(err)
	}
	fmt.Println(character.Name)
	fmt.Println(character.Class)
	fmt.Println(character.ThumbnailURL)
	fmt.Println(character.MythicPlusRecentRuns[0].Dungeon)

	defer resp.Body.Close()

	log.Print("Server running, listerning on port 5173")

	mainpage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, character)
	}

	http.HandleFunc("/", mainpage)

	log.Fatal(http.ListenAndServe(":5173", nil))
}
