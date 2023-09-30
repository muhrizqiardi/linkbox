package query

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
			where
				user_id = $1 and 
				unique_name = 'default'
	)
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where folder_id = (select id from default_folder)
		order by created_at asc
		limit $2 offset $3;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortDESC = `
	with default_folder as (
		select id from folders
			where
				user_id = $1 and 
				unique_name = 'default'
	)
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where folder_id = (select id from default_folder)
		order by created_at desc
		limit $2 offset $3;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortASC = `
	with default_folder as (
		select id from folders
			where
				user_id = $1 and 
				unique_name = 'default'
	)
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where folder_id = (select id from default_folder)
		order by updated_at asc
		limit $2 offset $3;
`
const QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC = `
	with default_folder as (
		select id from folders
			where
				user_id = $1 and 
				unique_name = 'default'
	)
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where folder_id = (select id from default_folder)
		order by updated_at desc
		limit $2 offset $3;
`
const QueryGetManyLinksInsideFolder_OrderByCreatedAtSortASC = `
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where 
			user_id = $1 and
			folder_id = $2
		order by created_at asc
		limit $3 offset $4;
`
const QueryGetManyLinksInsideFolder_OrderByCreatedAtSortDESC = `
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where 
			user_id = $1 and
			folder_id = $2
		order by created_at desc
		limit $3 offset $4;
`
const QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortASC = `
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where 
			user_id = $1 and
			folder_id = $2
		order by updated_at asc
		limit $3 offset $4;
`
const QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortDESC = `
	select l.id, url, title, description, user_id, folder_id, lm.media_path as media_path, l.created_at, l.updated_at
		from links as l
		right join link_medias as lm on lm.link_id = l.id
		where 
			user_id = $1 and
			folder_id = $2
		order by updated_at desc
		limit $3 offset $4;
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
