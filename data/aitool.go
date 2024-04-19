package data

import (
	"aiexplorer/db"
	"encoding/json"
	"strconv"
)

type AITool struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	ProfilePicture   string `json:"profile_picture"`
	SiteUrl          string `json:"site_url"`
	InstagramUrl     string `json:"instagram_url"`
	DiscordUrl       string `json:"discord_url"`
	LinkedinUrl      string `json:"linkedin_url"`
	GithubUrl        string `json:"github_url"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func GetAll() (aitools []AITool, err error) {
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	data, _, err := client.From("aitool").Select("*", "exact", false).Execute()

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &aitools)

	return
}

func Get(id int) (aitool *AITool, err error) {
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	data, _, err := client.From("aitool").Select("*", "", false).Eq("id", strconv.Itoa(id)).Single().Execute()

	if len(data) == 0 {
		err = nil
		return
	}

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &aitool)

	return

}

func Save(aitool *AITool) (returned AITool, err error) {
	var data []AITool
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	q := client.From("aitool").Insert(aitool, false, "do-nothing", "", "")
	_, err = q.ExecuteTo(&data)

	returned = data[0]

	return
}

func Update(aitool *AITool) (returned *AITool, err error) {
	var data []AITool
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	q := client.From("aitool").Update(aitool, "", "").Eq("id", strconv.Itoa(aitool.ID))
	_, err = q.ExecuteTo(&data)

	if err != nil {
		returned = nil
		return
	}

	if len(data) != 0 {
		returned = &data[0]
	}

	return
}
