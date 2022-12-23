package controllers

import (
	"log"
	"net/http"
	"strconv"
	"todo/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		meals, _ := models.GetMeals()
		user.Meals = meals

		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "create")
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		meal_date := r.PostFormValue("meal_date")
		meal_kind := r.PostFormValue("meal_kind")
		meal_menu := r.PostFormValue("meal_menu")
		meal_place := r.PostFormValue("meal_place")
		meal_shop_name := r.PostFormValue("meal_shop_name")

		meal_shop_price := r.PostFormValue("meal_shop_price")
		msp, _ := strconv.Atoi(meal_shop_price)

		if err := models.CreateMeal(meal_date, meal_kind, meal_menu, meal_place, meal_shop_name, msp); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/meals", 302)
	}
}

func edit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetMeal(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "edit")
	}
}

func update(w http.ResponseWriter, r *http.Request, id int) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		meal_date := r.PostFormValue("meal_date")
		meal_kind := r.PostFormValue("meal_kind")
		meal_menu := r.PostFormValue("meal_menu")
		meal_place := r.PostFormValue("meal_place")
		meal_shop_name := r.PostFormValue("meal_shop_name")

		meal_shop_price := r.PostFormValue("meal_shop_price")
		msp, _ := strconv.Atoi(meal_shop_price)

		m := &models.Meal{
			ID:            id,
			MealDate:      meal_date,
			MealKind:      meal_kind,
			MealMenu:      meal_menu,
			MealPlace:     meal_place,
			MealShopName:  meal_shop_name,
			MealShopPrice: msp}

		if err := m.UpdateMeal(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/meals", 302)
	}
}

func delete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetMeal(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteMeal(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/meals", 302)
	}
}
