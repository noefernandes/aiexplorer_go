package data

import (
	"aiexplorer/db"
	"encoding/json"
	"strconv"
)

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TagInput struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func GetAllTags() (tags []Tag, err error) {
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	data, _, err := client.From("tag").Select("*", "exact", false).Execute()

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &tags)
	return
}

func GetTag(id int) (tag *Tag, err error) {
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	data, _, err := client.From("tag").Select("*", "", false).Eq("id", strconv.Itoa(id)).Single().Execute()

	if len(data) == 0 {
		err = nil
		return
	}

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &tag)

	return

}

func SaveTag(tag *TagInput) (returned Tag, err error) {
	var data []Tag
	client, err := db.OpenConnection()

	if err != nil {
		return
	}

	q := client.From("tag").Insert(tag, false, "do-nothing", "", "")
	_, err = q.ExecuteTo(&data)

	if len(data) != 0 {
		returned = data[0]
	}

	return
}
