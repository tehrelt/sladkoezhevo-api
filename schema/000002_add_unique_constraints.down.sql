ALTER TABLE city                DROP CONSTRAINT IF EXISTS city_name_unique;
ALTER TABLE confectionary_type  DROP CONSTRAINT IF EXISTS confectionary_type_name_unique;
ALTER TABLE property_type       DROP CONSTRAINT IF EXISTS property_type_name_unique;
ALTER TABLE units               DROP CONSTRAINT IF EXISTS units_name_unique;