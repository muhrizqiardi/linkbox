package link

const QueryCreateLink = `
	insert into links (url, title, description, user_id, folder_id)
		values ($1, $2, $3, $4, $5)
		returning id, url, title, description, user_id, folder_id, created_at, updated_at;
`
const QueryGetOneLinkByID = `
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where id = $1
		limit 1;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortASC = `
	with default_folder as (
		select id from folders
			where unique_name = 'default'
	)
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by created_at asc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortDESC = `
	with default_folder as (
		select id from folders
			where unique_name = 'default'
	)
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by created_at desc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortASC = `
	with default_folder as (
		select id from folders
			where unique_name = 'default'
	)
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by updated_at asc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC = `
	with default_folder as (
		select id from folders
			where unique_name = 'default'
	)
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by updated_at desc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideFolder_OrderByCreatedAtSortASC = `
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by created_at asc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideFolder_OrderByCreatedAtSortDESC = `
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by created_at desc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortASC = `
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by updated_at asc
		limit $1 offset $2;
`
const QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortDESC = `
	select id, url, title, description, user_id, folder_id, created_at, updated_at
		from links
		where folder_id = (select id from default_folder)
		order by updated_at desc
		limit $1 offset $2;
`
const QueryUpdateOneLinkByID = `
	update links
		set
			url = $2,
			title = $3,
			description = $4,
			user_id = $5,
			folder_id = $6,
			created_at = current_timestamp
		where id = $1
		returning id, url, title, description, user_id, folder_id, created_at, updated_at;
`
const QueryDeleteOneLinkByID = `
	delete from links
		where id = $1
		returning id, url, title, description, user_id, folder_id, created_at, updated_at;
`
