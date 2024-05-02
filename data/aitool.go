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
	Tags             []Tag   `json:"tags"`
	ProfilePicture   *string `json:"profile_picture"`
	SiteUrl          *string `json:"site_url"`
	InstagramUrl     *string `json:"instagram_url"`
	DiscordUrl       *string `json:"discord_url"`
	LinkedinUrl      *string `json:"linkedin_url"`
	GithubUrl        *string `json:"github_url"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
	Favorited        bool    `json:"favorited"`
}

func GetAll(page int, size int) (aitools []*AITool, count int, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	aitools = make([]*AITool, 0)

	sql := `SELECT a.id, a.name, a.description, a.short_description, a.site_url,
	    		a.instagram_url, a.discord_url, a.linkedin_url, a.github_url, a.profile_picture, 
				a.tags, a.created_at, a.updated_at, favorited, at.tags_id, t.name as tag_name, t.color
			FROM (SELECT * FROM aitool ORDER BY id LIMIT $1 OFFSET $2) a
			LEFT JOIN aitool_tags at ON a.id = at.aitool_id 
			LEFT JOIN tag t ON at.tags_id = t.id`

	rows, err := conn.Query(sql, size, page*size)
	if err != nil {
		return
	}

	aitoolMap := make(map[int]*AITool)

	for rows.Next() {
		var aitoolID int
		var aitoolName, shortDesc, desc,
			profilePicture, siteUrl,
			instagramUrl, discordUrl, linkedinUrl,
			githubUrl, tags, createdAt, updatedAt, color *string
		var tagId *int
		var tagName *string
		var favorited bool

		err = rows.Scan(&aitoolID, &aitoolName, &desc, &shortDesc,
			&siteUrl, &instagramUrl, &discordUrl, &linkedinUrl,
			&githubUrl, &profilePicture, &tags, &createdAt, &updatedAt, &favorited, &tagId, &tagName, &color)
		if err != nil {
			return
		}

		if _, ok := aitoolMap[aitoolID]; !ok {
			aitoolMap[aitoolID] = &AITool{
				ID:               aitoolID,
				Name:             aitoolName,
				ShortDescription: shortDesc,
				Description:      desc,
				ProfilePicture:   profilePicture,
				SiteUrl:          siteUrl,
				InstagramUrl:     instagramUrl,
				DiscordUrl:       discordUrl,
				LinkedinUrl:      linkedinUrl,
				GithubUrl:        githubUrl,
				CreatedAt:        createdAt,
				UpdatedAt:        updatedAt,
				Favorited:        false,
				Tags:             make([]Tag, 0),
			}
		}

		if tagId != nil {
			aitoolMap[aitoolID].Tags = append(aitoolMap[aitoolID].Tags, Tag{
				ID:    *tagId,
				Name:  *tagName,
				Color: *color,
			})
		}
	}

	for _, v := range aitoolMap {
		aitools = append(aitools, v)
	}

	row := conn.QueryRow("SELECT COUNT(*) FROM aitool")
	err = row.Scan(&count)

	return
}

func Get(id int) (aitool AITool, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT id, name, description, short_Description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		tags, created_at, updated_at, favorited FROM aitool WHERE id = $1`

	row := conn.QueryRow(sql, id)

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

	//remover
	var date string = "12-10-1989"
	aitool.CreatedAt = &date
	aitool.UpdatedAt = &date

	sql := `INSERT INTO aitool (name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		created_at, updated_at, favorited) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 
		$9, $10, $11, $12) RETURNING id, name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		created_at, updated_at, favorited`

	err = conn.QueryRow(sql, aitool.Name, aitool.Description, aitool.ShortDescription,
		aitool.SiteUrl, aitool.InstagramUrl, aitool.DiscordUrl, aitool.LinkedinUrl,
		aitool.GithubUrl, aitool.ProfilePicture, aitool.CreatedAt,
		aitool.UpdatedAt).Scan(&returned.ID, &returned.Name, &returned.Description, &returned.ShortDescription,
		&returned.SiteUrl, &returned.InstagramUrl, &returned.DiscordUrl,
		&returned.LinkedinUrl, &returned.GithubUrl, &returned.ProfilePicture,
		&returned.CreatedAt, &returned.UpdatedAt, &returned.Favorited)

	sql = `INSERT INTO aitool_tags (tags_id, aitool_id) VALUES ($1, $2)`

	for _, tag := range aitool.Tags {
		conn.QueryRow(sql, tag.ID, returned.ID).Scan()
	}

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
				github_url = $8, profile_picture = $9, created_at = $10,
				updated_at = $11, favorited = $12
			WHERE id = $12 RETURNING id, name, description, short_description, site_url,
	    instagram_url, discord_url, linkedin_url, github_url, profile_picture, 
		created_at, updated_at, favorited`

	err = conn.QueryRow(sql, aitool.Name, aitool.Description, aitool.ShortDescription,
		aitool.SiteUrl, aitool.InstagramUrl, aitool.DiscordUrl, aitool.LinkedinUrl,
		aitool.GithubUrl, aitool.ProfilePicture, aitool.CreatedAt,
		aitool.UpdatedAt, aitool.Favorited, aitool.ID).Scan(&returned.ID, &returned.Name, &returned.Description,
		&returned.ShortDescription, &returned.SiteUrl, &returned.InstagramUrl,
		&returned.DiscordUrl, &returned.LinkedinUrl, &returned.GithubUrl,
		&returned.ProfilePicture, &returned.CreatedAt,
		&returned.UpdatedAt, &returned.Favorited)

	sql = `DELETE FROM aitool_tags WHERE aitool_id = $1`
	conn.Exec(sql, aitool.ID)

	sql = `INSERT INTO aitool_tags (tags_id, aitool_id) VALUES ($1, $2)`

	for _, tag := range aitool.Tags {
		conn.QueryRow(sql, tag.ID, returned.ID).Scan()
	}

	return
}

func Delete(id int) (res sql.Result, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	res, err = conn.Exec(`DELETE FROM aitool_tags WHERE aitool_id = $1`, id)

	if err != nil {
		return
	}

	res, err = conn.Exec(`DELETE FROM aitool WHERE id = $1`, id)

	return
}
