package controllers

import (
	"errors"
	"fmt"
	"github.com/digininja/go_practice/gorm/models"
	"log"
)

func CreateURL(theUrl string) {
	if theUrl == "" {
		log.Printf("Blank URL passed")
		return
	}
	url := models.URL{URL: theUrl, Status: "processing"}
	models.DB.Create(&url)
}

func FindURLByURL(theUrl string) (models.URL, error) {
	var url models.URL

	if theUrl == "" {
		log.Printf("Blank URL provided")

		return models.URL{}, errors.New("Blank URL provided")
	}
	if err := models.DB.Where("url = ?", theUrl).First(&url).Error; err != nil {
		log.Printf("Error, URL not found")

		return models.URL{}, errors.New("URL with URL provided not found")
	}

	log.Printf("Found URL: %d\n", url.ID)
	return url, nil
}

func FindURLByID(id uint) (models.URL, error) {
	var url models.URL

	if err := models.DB.Where("id = ?", id).First(&url).Error; err != nil {
		log.Printf("Error, URL not found")

		return models.URL{}, errors.New("URL with ID provided not found")
	}

	log.Printf("Found URL: %d\n", url.ID)
	return url, nil
}

func GetURLs() []models.URL {
	var urls []models.URL
	models.DB.Find(&urls)

	for _, url := range urls {
		log.Printf("URL: %s\n", url.URL)
	}

	return urls
}

type UpdateURLStatus struct {
	Status string `json:"status"`
}

func UpdateURL(id uint, status string) error {
	// Get model if exist
	var url models.URL
	url, err := FindURLByID(id)
	if err != nil {
		log.Printf("Error, could not find the specified ID")
		return errors.New("URL with ID provided not found")
	}

	// Validate input
	if status != "complete" && status != "error" {
		log.Printf("Error, invalid new status given")
		return errors.New("Invalid new status given")
	}

	input := UpdateURLStatus{Status: status}

	models.DB.Model(&url).Updates(input)

	fmt.Printf("Update went OK\n")
	return nil
}

func DeleteURL(id uint) error {
	// Get model if exist
	// Get model if exist
	var url models.URL
	url, err := FindURLByID(id)
	if err != nil {
		log.Printf("Error, could not find the specified ID")
		return errors.New("URL with ID provided not found")
	}

	models.DB.Delete(&url)

	log.Printf("All good, deleted\n")
	return nil
}
