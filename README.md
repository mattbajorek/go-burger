# Eat-Da-Burger
Week of 14 Homework Node Express Handlebars Redone in GO

LIVE PREVIEW --> https://go-eat-da-burger.herokuapp.com/

## Screenshots

Main | Type
-------------|--------
![Main Image](/readme_images/main.png?raw=true"main.png") | ![Type Image](/readme_images/type.png?raw=true"type.png")

Submit | Devour
-------------|--------
![Submit Image](/readme_images/submit.png?raw=true"submit.png") | ![Devour Image](/readme_images/devour.png?raw=true"devour.png")

## User Story
* User can type and submit a type of burger
* User can devour burger
* User can view both burgers to eat and burgers devoured

## Technologies used
- HTML (golang package template)
- CSS (Bootstrap)
- JavaScript (jQuery for Bootstrap)
- Go (external packages go-sql-driver and httprouter)
- MySQL

## How to Use

1. Click in text box and type burger desired.
2. Press "Submit" to post burger.
3. Press "Devour It!" to devour burger.
4. Admire all the burgers eaten.

## Built With

* Atom
* Gimp
* Go compiler

## Deployed With

* Heroku (Go)

## Walk throughs of code

Most interesting Go code
(Select all with MySQL driver)
```
// Create row struct
type row struct {
	Id         int
	BurgerName string
	Devoured   bool
	Date       string
}

// Select all the burgers
func selectAll() []row {

	// Gets a pointer to database connection
	db := dbConnect()
	defer db.Close()

	// Query to select all
	rows, err := db.Query("SELECT * FROM burgers")
	checkErr(err)
	defer rows.Close()

	// Create slice of rows
	burgers := []row{}

  // Loop through rows
	for rows.Next() {
		var r row
    // Add data to struct
		err = rows.Scan(&r.Id, &r.BurgerName, &r.Devoured, &r.Date)
		checkErr(err)
    // Append row to slice
		burgers = append(burgers, r)
	}

	return burgers
}
```

## Author

* [Matthew Bajorek](https://www.linkedin.com/in/matthewbajorek)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* [Todd McLeod](https://github.com/GoesToEleven) for the amazing YouTube tutorials for GO
