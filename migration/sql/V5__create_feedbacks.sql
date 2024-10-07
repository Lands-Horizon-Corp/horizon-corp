CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    feedback_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO feedback (email, description, feedback_type)
VALUES
('john.doe@example.com', 'This is a test feedback message.', 'Positive'),
('jane.smith@example.com', 'This is a suggestion for improvement.', 'Suggestion');
