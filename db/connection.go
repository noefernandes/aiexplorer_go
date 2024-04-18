package db

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

func OpenConnection() (*supabase.Client, error) {

	client, err := supabase.NewClient(os.Getenv("PROJ_URL"), os.Getenv("API_KEY"), nil)

	if err != nil {
		panic(err)
	}

	return client, err
}
