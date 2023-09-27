package query

const QueryCreateUserWithDefaultFolder = `
	with 
		new_user as (
			insert into users (username, password)		
				values ($1, $2)
				returning id, username, password, created_at, updated_at
		),
		user_default_folder as (
			insert into folders (unique_name, user_id)
				select 'default', id from new_user limit 1
				returning id, unique_name, user_id, created_at, updated_at
		)
		select id, username, password, created_at, updated_at
			from new_user;
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
