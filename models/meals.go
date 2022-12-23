package models

import (
	"log"
	"time"
)

type Meal struct {
	ID            int
	MealDate      string
	MealKind      string
	MealMenu      string
	MealPlace     string
	MealShopName  string
	MealShopPrice int
	CreatedAt     time.Time
}

func CreateMeal(meal_date string, meal_kind string, meal_menu string, meal_place string, meal_shop_name string, meal_shop_price int) (err error) {
	cmd := `insert into meals(meal_date, meal_kind, meal_menu, meal_place, meal_shop_name, meal_shop_price, created_at) values(?, ?, ?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd, meal_date, meal_kind, meal_menu, meal_place, meal_shop_name, meal_shop_price, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetMeal(id int) (meal Meal, err error) {
	cmd := `select * from meals where id = ?`

	meal = Meal{}

	err = Db.QueryRow(cmd, id).Scan(
		&meal.ID,
		&meal.MealDate,
		&meal.MealKind,
		&meal.MealMenu,
		&meal.MealPlace,
		&meal.MealShopName,
		&meal.MealShopPrice,
		&meal.CreatedAt)

	return meal, err
}

func GetMeals() (meals []Meal, err error) {
	cmd := `select * from meals`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var meal Meal
		err = rows.Scan(
			&meal.ID,
			&meal.MealDate,
			&meal.MealKind,
			&meal.MealMenu,
			&meal.MealPlace,
			&meal.MealShopName,
			&meal.MealShopPrice,
			&meal.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		meals = append(meals, meal)
	}
	rows.Close()

	return meals, err
}

func (m *Meal) UpdateMeal() error {
	cmd := `update meals set meal_date = ?, meal_kind = ?, meal_menu = ?, meal_place = ?, meal_shop_name = ?, meal_shop_price = ? where id = ?`
	_, err = Db.Exec(cmd, m.MealDate, m.MealKind, m.MealMenu, m.MealPlace, m.MealShopName, m.MealShopPrice, m.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (m *Meal) DeleteMeal() error {
	cmd := `delete from meals where id = ?`
	_, err = Db.Exec(cmd, m.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
