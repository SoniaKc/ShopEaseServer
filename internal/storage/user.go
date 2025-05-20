package storage


func AddUser(firstName string, lastName string){
	var count int
    err := DB.QueryRow("SELECT * FROM users WHERE firstName = $1 AND lastname = $2", firstName, lastName).Scan(&count)

    if err == sql.ErrNoRows {
        _, errInsert := DB.Exec("INSERT INTO users (firstName, lastName) VALUES ($1, $2)", firstName, lastName)
        return errInsert
    }

    if err != nil {
        return fmt.Errorf("error while checking users: %v", err)
    }
    if count > 0 {
        return errors.New("user already exist")
    }
    _, errInsert := DB.Exec("INSERT INTO users (firstName, lastName) VALUES ($1, $2)", firstName, lastName)
    return errInsert
}