package query

const QueryCreateUser = `
	insert into users (username, password)		
		values ($1, $2)
		returning id, username, password, created_at, updated_at;
`
const QueryGetOneUserByID = `
	select id, username, password, created_at, updated_at		
		from users
		where id = $1;
`
const QueryGetOneUserByUsername = `
	select id, username, password, created_at, updated_at		
		from users
		where username = $1;
`
const QueryUpdateOneUserByID = `
	update users
		set
			username = $2,
			password = $3,
			updated_at = current_timestamp
		where id = $1
		returning id, username, password, created_at, updated_at;
`
const QueryDeleteOneUserByID = `
	delete from users
		where id = $1
		returning id, username, password, created_at, updated_at;
`
