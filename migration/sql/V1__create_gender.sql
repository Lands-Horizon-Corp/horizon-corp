CREATE TABLE IF NOT EXISTS genders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);


INSERT INTO genders (name, description) VALUES
('Male', 'A person who identifies as male, typically associated with masculinity.'),
('Female', 'A person who identifies as female, typically associated with femininity.'),
('Non-binary', 'A person who does not exclusively identify as male or female.'),
('Genderqueer', 'A person who identifies outside of the traditional gender binary.'),
('Genderfluid', 'A person whose gender identity varies over time.'),
('Agender', 'A person who identifies as having no gender or being gender-neutral.'),
('Bigender', 'A person who identifies as two genders, either simultaneously or at different times.'),
('Two-Spirit', 'A term used by some Indigenous peoples to describe a person with both masculine and feminine spirits.'),
('Other', 'A gender identity that does not fall within traditional categories.');