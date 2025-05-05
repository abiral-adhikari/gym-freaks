package queries

const CreateUserTableSQL = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	phone BIGINT UNIQUE NOT NULL,
	dob TIMESTAMP,
	role TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	token TEXT
);
CREATE TABLE IF NOT EXISTS foods (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	calories INT NOT NULL,
	unit TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS meals (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	food_id INT NOT NULL REFERENCES foods(id) ON DELETE CASCADE,
	quantity INT NOT NULL,
	time TIMESTAMPTZ NOT NULL,
	meal_type TEXT,
	notes TEXT
);

CREATE TABLE IF NOT EXISTS exercises (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	type TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS workouts (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
	sets INT NOT NULL,
	reps INT NOT NULL,
	weight INT,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
`

// This file contains SQL queries for the Gym Freaks application.
const CreateUserQuery = `
	INSERT INTO users (username, password, email, phone, dob, role)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;
	`

const GetUserByIDQuery = `
	SELECT * FROM users WHERE id = $1;
	`
const UpdateUserTokenQuery = `
	UPDATE users SET token = $1 WHERE id = $2;
	`

const LoginQuery = `
	SELECT * FROM users WHERE phone = $1	
`

const LogoutQuery = `
	UPDATE users SET token = '' WHERE id = $1;
	`

const CreateFoodQuery = `
	INSERT INTO food (name,calories,unit)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

const UpdateFoodQuery = `
	UPDATE food 
	SET name=$2, calories=$3,unit=$4
	where id=$1
	returning id;
`

const DeleteFoodQuery = `
	DELETE FROM food
	where id=$1
	`

const GetFoodQuery = `
	SELECT id, name, calories, unit
	FROM food
	WHERE 1=1
`
