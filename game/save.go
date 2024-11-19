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
	AddItem(name, "coin", 0)	
	AddItem(name, "background", 0)
	AddItem(name, "childhood", 0)
	AddItem(name, "now", 0)
	AddItem(name, "future", 0)

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

func GetPlayers() []string {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name FROM players")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		// if no rows, return empty slice
		if err == sql.ErrNoRows {
			return []string{}
		}
		log.Fatal(err)
	}

	players := make([]string, 0)
	for rows.Next() {
		var player string
		err = rows.Scan(&player)
		if err != nil {
			log.Fatal(err)
		}
		players = append(players, player)
	}
	return players
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

func AddItem(player, item string, count int) {
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

	_, err = stmt.Exec(player, item, count)
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

	stmt, err := db.Prepare("SELECT count FROM inventory WHERE player = ? AND item = 'coin'")
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

func RemovePlayer(name string) {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM players WHERE name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		log.Fatal(err)
	}
}

func PurchaseItem(player, item string, count int) {
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

	_, err = stmt.Exec(player, item, count)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckInventory(player, item string) int {
	db, err := sql.Open("sqlite3", "./game_data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT count FROM inventory WHERE player = ? AND item = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(player, item).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}