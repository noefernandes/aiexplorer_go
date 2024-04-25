package data

import (
	"aiexplorer/db"
)

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func GetAllTags() (tags []Tag, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT id, name, color FROM tag")
	if err != nil {
		return
	}

	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.Color)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}

	return
}

func GetTag(id int) (tag Tag, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := "SELECT id, name, color FROM tag WHERE id = $1"

	row := conn.QueryRow(sql, id)

	err = row.Scan(&tag.ID, &tag.Name, &tag.Color)

	return

}

func SaveTag(tag *Tag) (returned Tag, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO tag (name, color) VALUES ($1, $2) RETURNING id, name, color`

	err = conn.QueryRow(sql, tag.Name, tag.Color).Scan(&returned.ID, &returned.Name, &returned.Color)

	return
}
