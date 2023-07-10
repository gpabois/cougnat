package reporting_pg

// Query to get a report by its ID
//
// Selects : report.id, owner is set, report.owner__id, report.__type,
// report_type.id, report_type.name, report_type.label, report_type.nature,
// report.location,
const PG_REPORT_GET_BY_ID_QUERY = `
SELECT r.id, 
	CAST(CASE WHEN r.owner__id IS NULL AND r.owner__nature == "" THEN 0 ELSE 1 END AS owner_is_set), r.owner__id, r.owner__type, 
	rt.id, rt.name, rt.label, rt.nature, ST_AsGeoJSON(r.location) 
FROM reports AS r 
LEFT JOIN reports_types AS rt ON rt.id = r.type__id
WHERE r.id = $1
`

// Insert report sql query
//
// Columns : owner__id, owner__nature, type__id, rate, location
const PG_INSERT_REPORT_QUERY = `
INSERT INTO reports AS r 
	(r.owner__id, r.owner__nature, r.type__id, r.rate, r.location)
VALUES 
	($1, $2, $3, $4, ST_GeomFromGeoJSON($5))
RETURNING r.ID
`

const PG_DELETE_BY_ID_REPORT_QUERY = `
DELETE FROM reports AS r
WHERE r.id = $1
`
