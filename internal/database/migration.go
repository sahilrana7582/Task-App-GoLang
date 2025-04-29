package database

import (
	"database/sql"
	"fmt"
	"log"
)

func MigrateDB(DB *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		status VARCHAR(50) DEFAULT 'Pending',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Failed to create tasks table: %v", err)
	}

	fmt.Println("✅ Tasks table created or already exists")

}
