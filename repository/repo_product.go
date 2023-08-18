package repository

import (
	"errors"
	"fmt"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/models"
)

func RepoAddProduct(productIn models.Product) (productRes models.Product, err error) {
	fmt.Println("RepoAddProduct")
	tx := config.DB.Model(&models.Product{}).Create(&productIn)
	if err = tx.Error; err != nil {
		return
	}
	return productIn, err

}

func RepoShowAllProduct() (productRes []models.Product, err error) {
	fmt.Println("RepoShowAllProduct")
	tx := config.DB.Model(&models.Product{}).Find(&productRes)
	if err = tx.Error; err != nil {
		return
	}
	return
}
func RepoFindById(productId uint64) (productRes models.Product, err error) {
	fmt.Println("RepoFindById")
	tx := config.DB.Model(&models.Product{}).Where("id=?", productId).Find(&productRes)
	if err = tx.Error; err != nil {
		return
	}
	if productRes.Id <= 0 {
		err = errors.New("NOT FOUND")
	}
	return
}

func RepoUpdateById(productIn models.Product) (productRes models.Product, err error) {
	fmt.Println("RepoUpdateById")
	tx := config.DB.Model(&models.Product{}).Where("id=?", productIn.Id).Updates(&productIn)
	if err = tx.Error; err != nil {
		return
	}
	return productIn, err
}

func RepoSoftDeleteByid(productId uint64) (err error) {
	fmt.Println("RepoSoftDeleteByid")
	tx := config.DB.Model(&models.Product{}).Where("id=?", productId).Delete(&models.Product{})
	if err = tx.Error; err != nil {
		return
	}
	return
}
