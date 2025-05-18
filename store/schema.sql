-- Create engine table
CREATE TABLE IF NOT EXISTS engine (
    id UUID PRIMARY KEY,
    displacement INT NOT NULL,
    no_of_cylinders INT NOT NULL,
    car_range INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS car (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    engine_id UUID NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Add foreign key constraint on engine_id in car table
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'fk_engine_id'
    ) THEN
        ALTER TABLE car
        ADD CONSTRAINT fk_engine_id
        FOREIGN KEY (engine_id)
        REFERENCES engine(id)
        ON DELETE CASCADE;
    END IF;
END$$;