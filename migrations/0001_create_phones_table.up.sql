CREATE TABLE phones (
                        id SERIAL PRIMARY KEY,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW(),
                        deleted_at TIMESTAMP,

                        brand TEXT NOT NULL,
                        model_name TEXT NOT NULL,
                        price NUMERIC(10, 2) NOT NULL,
                        release_date TEXT NOT NULL
);
