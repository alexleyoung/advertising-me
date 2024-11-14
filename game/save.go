package game

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitSaves() {
	db, err := sql.Open("sqlite3", "./game_data.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    createTableSQL := `
	CREATE TABLE IF NOT EXISTS players (
        name TEXT PRIMARY KEY
    );
	
	CREATE TABLE IF NOT EXISTS scores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		player TEXT,
		score INTEGER,
		nearMisses INTEGER,
		timestamp INTEGER,
		foreign key (player) references player(name)
	);

	CREATE TABLE IF NOT EXISTS inventory (
		player TEXT,
		item TEXT,
		count INTEGER,
		foreign key (player) references player(name)
		foreign key (item) references items(name)
		primary key (player, item)
	);

	CREATE TABLE IF NOT EXISTS items (
		name TEXT PRIMARY KEY
	);
	`
    
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func CreatePlayer(name string) {
	AddCoins(name, 0)	

	name = strings.ToLower(name)
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
	Timestamp int
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
		err = rows.Scan(&score.id, &score.Player, &score.Score, &score.NearMisses, &score.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		scores = append(scores, &score)
	}

	return scores
}

func AddCoins(name string, count int) {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
	INSERT INTO inventory (player, item, count) 
	VALUES (?, ?, ?) 
	ON CONFLICT (player, item) 
	DO UPDATE SET count = inventory.count + excluded.count
	`) 
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, "coins", count)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCoins(name string) int {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT count FROM inventory WHERE player = ? AND item = 'coins'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(name).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}