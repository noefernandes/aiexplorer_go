package data

import (
	"aiexplorer/db"
	"database/sql"
)

type AITool struct {
	ID               int     `json:"id"`
	Name             *string `json:"name"`
	ShortDescription *string `json:"short_description"`
	Description      *string `json:"description"`
	Tags             *string `json:"tags"`
	ProfilePicture   *string `json:"profile_picture"`
	SiteUrl          *string `json:"site_url"`
	InstagramUrl     *string `json:"instagram_url"`
	DiscordUrl       *string `json:"discord_url"`
	LinkedinUrl      *string `json:"linkedin_url"`
	GithubUrl        *string `json:"github_url"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

func GetAll(page int, size int) (aitools []AITool, count int, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM aitool ORDER BY id LIMIT $1 OFFSET $2`

	rows, err := conn.Query(sql, size, page*size)
	if err != nil {
		return
	}

	for rows.Next() {
		var aitool AITool

		err = rows.Scan(&aitool.ID, &aitool.Name, &aitool.Description, &aitool.ShortDescription,
			&aitool.SiteUrl, &aitool.InstagramUrl, &aitool.DiscordUrl, &aitool.LinkedinUrl,
			&aitool.GithubUrl, &aitool.ProfilePicture, &aitool.Tags, &aitool.CreatedAt, &aitool.UpdatedAt)
		if err != nil {
			continue
		}

		aitools = append(aitools, aitool)
	}

	row := conn.QueryRow("SELECT COUNT(*) FROM aitool")
	err = row.Scan(&count)
	if err != nil {
		return
	}

	return
}

func Get(id int) (aitool AITool, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM aitool WHERE id = $1", id)

	err = row.Scan(&aitool.ID, &aitool.Name, &aitool.Description, &aitool.ShortDescription,
		&aitool.SiteUrl, &aitool.InstagramUrl, &aitool.DiscordUrl, &aitool.LinkedinUrl,
		&aitool.GithubUrl, &aitool.ProfilePicture, &aitool.Tags, &aitool.CreatedAt,
		&aitool.UpdatedAt)

	return
}

func Save(aitool *AITool) (returned AITool, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO aitool (name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 
		$9, $10, $11, $12) RETURNING id, name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		tags, created_at, updated_at`

	err = conn.QueryRow(sql, aitool.Name, aitool.Description, aitool.ShortDescription,
		aitool.SiteUrl, aitool.InstagramUrl, aitool.DiscordUrl, aitool.LinkedinUrl,
		aitool.GithubUrl, aitool.ProfilePicture, aitool.Tags, aitool.CreatedAt,
		aitool.UpdatedAt).Scan(&returned.ID, &returned.Name, &returned.Description, &returned.ShortDescription,
		&returned.SiteUrl, &returned.InstagramUrl, &returned.DiscordUrl,
		&returned.LinkedinUrl, &returned.GithubUrl, &returned.ProfilePicture,
		&returned.Tags, &returned.CreatedAt, &returned.UpdatedAt)

	return
}

func Update(aitool *AITool) (returned AITool, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE aitool SET name = $1, description = $2, short_description = $3,
				site_url = $4, instagram_url = $5, discord_url = $6, linkedin_url = $7, 
				github_url = $8, profile_picture = $9, tags = $10, created_at = $11,
				updated_at = $12
			WHERE id = $13 RETURNING id, name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		tags, created_at, updated_at`

	err = conn.QueryRow(sql, aitool.Name, aitool.Description, aitool.ShortDescription,
		aitool.SiteUrl, aitool.InstagramUrl, aitool.DiscordUrl, aitool.LinkedinUrl,
		aitool.GithubUrl, aitool.ProfilePicture, aitool.Tags, aitool.CreatedAt,
		aitool.UpdatedAt, aitool.ID).Scan(&returned.ID, &returned.Name, &returned.Description,
		&returned.ShortDescription, &returned.SiteUrl, &returned.InstagramUrl,
		&returned.DiscordUrl, &returned.LinkedinUrl, &returned.GithubUrl,
		&returned.ProfilePicture, &returned.Tags, &returned.CreatedAt,
		&returned.UpdatedAt)

	return
}

func Delete(id int) (res sql.Result, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	res, err = conn.Exec(`DELETE FROM aitool WHERE id = $1`, id)

	return
}
