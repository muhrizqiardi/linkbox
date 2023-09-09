package folder

const QueryCreateFolder = `
	insert into folders (unique_name, user_id)
		values ($1, $2)
		returning id, user_id, unique_name, created_at, updated_at;
`
const QueryGetOneFolderByID = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		where id = $1
		limit 1;
`
const QueryGetOneFolderByUniqueName = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		where unique_name = $1, user_id = $2
		limit 1;
`
const QueryGetManyFolders_ByUpdatedAtASC = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		order by updated_at asc, user_id = $3
		limit $1 offset $2;
`
const QueryGetManyFolders_ByUpdatedAtDESC = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		order by updated_at desc, user_id = $3
		limit $1 offset $2;
`
const QueryGetManyFolders_ByCreatedAtASC = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		order by created_at asc, user_id = $3
		limit $1 offset $2;
`
const QueryGetManyFolders_ByCreatedAtDESC = `
	select id, user_id, unique_name, created_at, updated_at
		from folders
		order by created_at desc, user_id = $3
		limit $1 offset $2;
`
const QueryUpdateOneFolderByID = `
	update folders
		set unique_name = $2
		where id = $1
		returning id, user_id, unique_name, created_at, updated_at;
`
const QueryDeleteOneFolderByID = `
	delete from folders
		where id = $1
		returning id, user_id, unique_name, created_at, updated_at;
`
