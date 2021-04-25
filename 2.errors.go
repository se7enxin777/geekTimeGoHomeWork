//...

var id int
var name string
var age string

func getUserInfo() error {
	err := db.QueryRow("select id,name,age from User where xxx").Scan(&id, &name, &age)
	if err != nil {
		return error.Wrap(err, "failed to get user info")
	}
	return nil
}

func main() {

	//...

	if err := getUserInfo(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("FAIL: %+v\n", err)
		} else {
			fmt.Printf(err)
		}
	}

	//...
}


