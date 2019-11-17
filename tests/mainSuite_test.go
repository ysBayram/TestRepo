package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/ysbayram/TestRepo/server"
)

type TestSuite struct {
	suite.Suite
	db             *gorm.DB
	router         *gin.Engine
	performRequest func(r http.Handler, method, path, bodyStr string) *httptest.ResponseRecorder
}

func (suite *TestSuite) SetupSuite() {
	suite.db = server.CreateDBCon()
	suite.router = server.SetupRouter(suite.db)
	suite.performRequest = func(r http.Handler, method, path, bodyStr string) *httptest.ResponseRecorder {
		var body io.Reader
		if len(bodyStr) > 0 {
			body = bytes.NewBuffer([]byte(bodyStr))
		}
		req, _ := http.NewRequest(method, path, body)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w
	}
}

func (suite *TestSuite) TearDownSuite() {
	defer func() {
		suite.db.Close()
		os.Remove("./testRepo.db")
	}()
}

func TestRunAllSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
