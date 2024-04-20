package data

import (
	"aiexplorer/db"
	"encoding/json"
	"strconv"
)

type AITool struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description      string `json:"description"`
	ProfilePicture   string `json:"profilePicture"`
	SiteUrl          string `json:"siteUrl"`
	InstagramUrl     string `json:"instagramUrl"`
	DiscordUrl       string `json:"discordUrl"`
	LinkedinUrl      string `json:"linkedinUrl"`
	GithubUrl        string `json:"githubUrl"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

func GetAll(page int, size int) (aitools []AITool, count int64, err error) {
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	data, count, err := client.From("aitool").Select("*", "exact", false).Range(page*size, (page+1)*size-1, "").Execute()

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

	if len(data) != 0 {
		returned = data[0]
	}

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
