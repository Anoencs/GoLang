-- SELECT  b.*,a.name as parent1_name
-- FROM okr_orgs a,
--     (
--         SELECT u.name, u.role, u.department, o.name as org, o.org_id as parent_org_id
--         FROM okr_users u
--         left join okr_orgs o on u.org_id  = o.id
--     ) b
-- WHERE a.id = b.parent_org_id;

SELECT c.*,d.name as parent2_name
FROM okr_orgs d, 
    (
        SELECT  b.*,a.org_id as parent2_id,a.name as parent1_name
        FROM okr_orgs a,
        (
        SELECT u.name, u.role, u.department, o.name as org, o.org_id as parent_org_id
        FROM okr_users u
        left join okr_orgs o on u.org_id  = o.id
        ) b
        WHERE a.id = b.parent_org_id
    ) c
WHERE c.parent2_id = d.id