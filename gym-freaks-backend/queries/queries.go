package queries

const CreateUserTableSQL = `
--DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	phone BIGINT UNIQUE NOT NULL,
	dob TIMESTAMP,
	role TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	token TEXT,
	goal INT ,
	weight INT NOT NULL
);

CREATE TABLE IF NOT EXISTS foods(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	calories INT NOT NULL,
	unit TEXT NOT NULL,
	createdby INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
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
	type TEXT NOT NULL,
	createdby INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
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
	INSERT INTO users (username, password, email, phone, dob, role,goal,weight)
	VALUES ($1, $2, $3, $4, $5, $6,$7,$8)
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

// Food And Meals
const CreateFoodQuery = `
	INSERT INTO foods (name,calories,unit)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

const UpdateFoodQuery = `
	UPDATE foods 
	SET name=$2, calories=$3,unit=$4
	where id=$1
	returning id;
`

const DeleteFoodQuery = `
	DELETE FROM foods
	where id=$1
	`

const GetFoodByIDQuery = `
	SELECT * FROM foods WHERE id = $1;
`

const GetFoodByNameQuery = `
	SELECT 
		f.id,f.name,f.calories,f.unit,
		u.id,u.username,u.role
		FROM foods f
	JOIN users u ON f.createdyby = u.id
	WHERE f.id =$1 `

const SearchFoodQuery = `
	SELECT f.id,f.name,f.calories,f.unit,u.id,u.username
	FROM foods f
	JOIN users u ON f.createdyby = u.id
	WHERE 1=1
`
const SearchMealQuery = `
	SELECT m.id, u.id, u.username, f.id, f.name, m.quantity, m.time, m.meal_type, m.notes
	FROM meals m
	JOIN users u ON m.user_id = u.id
	JOIN foods f ON m.food_id = f.id
	WHERE 1=1
`
