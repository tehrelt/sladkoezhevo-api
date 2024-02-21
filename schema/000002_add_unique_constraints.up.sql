ALTER TABLE city                ADD CONSTRAINT city_name_unique                 UNIQUE (name);
ALTER TABLE confectionary_type  ADD CONSTRAINT confectionary_type_name_unique   UNIQUE (name);
ALTER TABLE property_type       ADD CONSTRAINT property_type_name_unique        UNIQUE (name);
ALTER TABLE units               ADD CONSTRAINT units_name_unique                UNIQUE (name);