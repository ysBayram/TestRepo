package tests

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (suite *TestSuite) Test_Product_1_CreateProduct() {
	expectedBody := gin.H{"Title": "TestProduct", "Price": 15000.0, "CategoryID": 1}

	requestBody :=
		`{"Title":"TestProduct","Price": 15000.0,"CategoryID":1}`

	w := suite.performRequest(suite.router, "POST", "/product", requestBody)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	suite.Equal(nil, err)

	value, exists := response["Title"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)

	value, exists = response["Price"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Price"], value)

	value, exists = response["CategoryID"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["CategoryID"], int(value.(float64)))
}

func (suite *TestSuite) Test_Product_2_UpdateProduct() {
	expectedBody := gin.H{"Title": "UpdatedTestProduct", "Price": 25000.0, "CategoryID": 1}

	requestBody := `{"ID": 1,"Title": "UpdatedTestProduct", "Price": 25000.0, "CategoryID": 1}`

	w := suite.performRequest(suite.router, "PUT", "/product/1", requestBody)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)

	value, exists = response["Price"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Price"], value)

	value, exists = response["CategoryID"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["CategoryID"], int(value.(float64)))
}

func (suite *TestSuite) Test_Product_3_GetProductByID() {
	expectedBody := gin.H{"Title": "UpdatedTestProduct", "Price": 25000.0, "CategoryID": 1}

	w := suite.performRequest(suite.router, "GET", "/product/1", "")

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)

	value, exists = response["Price"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Price"], value)

	value, exists = response["CategoryID"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["CategoryID"], int(value.(float64)))
}

func (suite *TestSuite) Test_Product_4_GetAllProduct() {
	expectedBody := gin.H{"Title": "UpdatedTestProduct", "Price": 25000.0, "CategoryID": 1}

	w := suite.performRequest(suite.router, "GET", "/product", "")

	suite.Equal(http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response[0]["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)

	value, exists = response[0]["Price"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Price"], value)

	value, exists = response[0]["CategoryID"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["CategoryID"], int(value.(float64)))
}

func (suite *TestSuite) Test_Product_5_DeleteProduct() {
	expectedBody := gin.H{"Title": "UpdatedTestProduct", "Price": 25000.0, "CategoryID": 1}

	w := suite.performRequest(suite.router, "DELETE", "/product/1", "")

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)

	value, exists = response["Price"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Price"], value)

	value, exists = response["CategoryID"]
	suite.Equal(true, exists)
	suite.Equal(expectedBody["CategoryID"], int(value.(float64)))
}
