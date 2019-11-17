package tests

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (suite *TestSuite) Test_Category_1_CreateCategory() {
	expectedBody := gin.H{
		"Title":            "TestCategoryTitle",
		"ParentCategoryID": 0,
	}

	requestBody :=
		`{"Title":"TestCategoryTitle","ParentCategoryID": 0}`

	w := suite.performRequest(suite.router, "POST", "/category", requestBody)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)
}

func (suite *TestSuite) Test_Category_2_UpdateCategory() {
	expectedBody := gin.H{
		"Title":            "UpdatedTestCategoryTitle",
		"ParentCategoryID": 0,
	}

	requestBody := `{"ID": 1,"Title":"UpdatedTestCategoryTitle","ParentCategoryID": 0}`

	w := suite.performRequest(suite.router, "PUT", "/category/1", requestBody)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)
}

func (suite *TestSuite) Test_Category_3_GetCategoryByID() {
	expectedBody := gin.H{
		"Title":            "UpdatedTestCategoryTitle",
		"ParentCategoryID": 0,
	}

	w := suite.performRequest(suite.router, "GET", "/category/1", "")

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)
}

func (suite *TestSuite) Test_Category_4_GetAllCategory() {
	expectedBody := gin.H{
		"Title":            "UpdatedTestCategoryTitle",
		"ParentCategoryID": 0,
	}

	w := suite.performRequest(suite.router, "GET", "/category", "")

	suite.Equal(http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response[0]["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)
}

func (suite *TestSuite) Test_Category_5_DeleteCategory() {
	expectedBody := gin.H{
		"Title":            "UpdatedTestCategoryTitle",
		"ParentCategoryID": 0,
	}

	requestBody := `{"ID": 1,"Title":"UpdatedTestCategoryTitle","ParentCategoryID": 0}`

	w := suite.performRequest(suite.router, "DELETE", "/category/1", requestBody)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["Title"]

	suite.Equal(nil, err)
	suite.Equal(true, exists)
	suite.Equal(expectedBody["Title"], value)
}
