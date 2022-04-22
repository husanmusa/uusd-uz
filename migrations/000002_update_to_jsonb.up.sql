alter table companies alter column name type jsonb USING (trim(name)::jsonb);
alter table companies alter column slogan type jsonb USING (trim(slogan)::jsonb);

alter table services alter column name type jsonb USING (trim(name)::jsonb);
alter table services alter column description type jsonb USING (trim(description)::jsonb);

alter table sets alter column name type jsonb USING (trim(name)::jsonb);
alter table sets alter column description type jsonb USING (trim(description)::jsonb);

alter table packages alter column name type jsonb USING (trim(name)::jsonb);
alter table packages alter column description type jsonb USING (trim(description)::jsonb);

-- CREATE INDEX idx_companies_name ON companies USING gin (name);
