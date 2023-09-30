package query

const QueryInsertLinkMedia = `
	insert into link_medias (link_id, media_path)
		values ($1, $2)
		returning id, link_id, media_path, created_at, updated_at;
`
