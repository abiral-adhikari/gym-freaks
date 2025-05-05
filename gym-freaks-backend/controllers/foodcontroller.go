package controllers

import (
	"context"
	"fmt"
	database "gym-freaks-backend/connections"
	"gym-freaks-backend/models"
	queries "gym-freaks-backend/queries"
)

type FoodController struct{}

func (FC *FoodController) Create(food models.Food) (int, error) {
	conn := database.DBConnect()

	err := conn.QueryRow(context.Background(), queries.CreateFoodQuery, food.Name, food.Calories, food.Unit).Scan(&food.FoodID)
	if err != nil {
		return 0, fmt.Errorf("error inserting food %v", err)
	}
	return food.FoodID, nil

}

func (FC *FoodController) Update(food models.Food) (int, error) {
	conn := database.DBConnect()

	err := conn.QueryRow(context.Background(), queries.UpdateFoodQuery, food.FoodID, food.Name, food.Calories, food.Unit).Scan(&food.FoodID)
	if err != nil {
		return 0, fmt.Errorf("error inserting food %v", err)
	}
	return food.FoodID, nil
}

func (FC *FoodController) Delete(food models.Food) (int, error) {
	conn := database.DBConnect()
	var err error

	_, err = conn.Exec(context.Background(), queries.DeleteFoodQuery, food.FoodID)
	if err != nil {
		return 0, fmt.Errorf("error inserting food %v", err)
	}
	return food.FoodID, nil

}

func (fc *FoodController) Get(name, unit string, minCalories, maxCalories int) ([]models.Food, error) {
	query := queries.GetFoodQuery
	args := []interface{}{}
	argIndex := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+name+"%")
		argIndex++
	}
	if unit != "" {
		query += fmt.Sprintf(" AND unit = $%d", argIndex)
		args = append(args, unit)
		argIndex++
	}
	if minCalories > 0 {
		query += fmt.Sprintf(" AND calories >= $%d", argIndex)
		args = append(args, minCalories)
		argIndex++
	}
	if maxCalories > 0 {
		query += fmt.Sprintf(" AND calories <= $%d", argIndex)
		args = append(args, maxCalories)
		argIndex++
	}

	conn := database.DBConnect()
	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []models.Food
	for rows.Next() {
		var f models.Food
		if err := rows.Scan(&f.FoodID, &f.Name, &f.Calories, &f.Unit); err != nil {
			return nil, err
		}
		foods = append(foods, f)
	}
	return foods, nil
}

var FoodControllers = &FoodController{}
