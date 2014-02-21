package controllers

import (
	"github.com/robfig/revel"

	"gorevel/app/models"
	"gorevel/app/routes"
)

type Admin struct {
	Application
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) ListUser() revel.Result {
	var users []models.User
	engine.Find(&users)

	return c.Render(users)
}

func (c Admin) DeleteUser(id int64) revel.Result {
	aff, _ := engine.Id(id).Delete(&models.User{})
	if aff > 0 {
		return c.RenderJson(map[string]bool{"status": true})
	}

	return c.RenderJson(map[string]bool{"status": false})
}

func (c Admin) ActivateUser(id int64) revel.Result {
	aff, _ := engine.Id(id).Cols("is_active").Update(&models.User{IsActive: true})
	if aff > 0 {
		return c.RenderJson(map[string]bool{"status": true})
	}

	return c.RenderJson(map[string]bool{"status": false})
}

func (c Admin) ListCategory() revel.Result {
	categories := getCategories()

	return c.Render(categories)
}

func (c Admin) DeleteCategory(id int64) revel.Result {
	aff, _ := engine.Id(id).Delete(&models.Category{})
	if aff > 0 {
		return c.RenderJson(map[string]bool{"status": true})
	}

	return c.RenderJson(map[string]bool{"status": false})
}

func (c Admin) NewCategory() revel.Result {
	title := "新建分类"
	return c.Render(title)
}

func (c Admin) NewCategoryPost(category models.Category) revel.Result {
	category.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	aff, _ := engine.Insert(&category)
	if aff == 0 {
		c.Flash.Error("添加分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}

func (c Admin) EditCategory(id int64) revel.Result {
	title := "编辑分类"

	var category models.Category
	has, _ := engine.Id(id).Get(&category)
	if !has {
		return c.NotFound("分类不存在")
	}

	c.Vars(Vars{
		"title":    title,
		"category": category,
	})

	return c.RenderTemplate("admin/NewCategory.html")
}

func (c Admin) EditCategoryPost(id int64, category models.Category) revel.Result {
	category.Id = id
	category.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	aff, _ := engine.Id(id).Update(&category)
	if aff == 0 {
		c.Flash.Error("编辑分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}

func getCategories() []models.Category {
	var categories []models.Category
	engine.Find(&categories)

	return categories
}
