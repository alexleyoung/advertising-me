package game

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitSaves() {
	db, err := sql.Open("sqlite3", "./game_data.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    createTableSQL := `CREATE TABLE IF NOT EXISTS players (
        name TEXT PRIMARY KEY
    );`
    
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

	createTableSQL = `CREATE TABLE IF NOT EXISTS scores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player TEXT,
		score INTEGER,
		nearMisses INTEGER,
		timestamp INTEGER,
		foreign key (player) references player(name)
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePlayer(name string) {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO players (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: players.name" {
			return
		}
		log.Fatal(err)
	}
}

func SavePlayerData(name string, score int, nearMisses int, timestamp int64) {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO scores (player, score, nearMisses, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, score, nearMisses, timestamp)
	if err != nil {
		log.Fatal(err)
	}
}

type Score struct {
	id int
	Player string
	Score int
	NearMisses int
}

func GetHighScores() []*Score {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM scores ORDER BY score DESC LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	scores := make([]*Score, 0)
	for rows.Next() {
		var score Score
		err = rows.Scan(&score.id, &score.Player, &score.Score, &score.NearMisses)
		if err != nil {
			log.Fatal(err)
		}
		scores = append(scores, &score)
	}

	return scores
}