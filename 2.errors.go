


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

	err := getUserInfo()
	if err != nil {
		fmt.Printf("FAIL: %+v\n", err)
	}
	//...
}


