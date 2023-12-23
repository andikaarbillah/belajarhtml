package setupQuery

const(
	InsertUser = `insert into users(name, email, password, role) values ($1,$2,$3, $4);`
)