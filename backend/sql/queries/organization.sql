INSERT INTO organizations(id, org_name, reg_number, contact_email, website, is_verified, is_active)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;