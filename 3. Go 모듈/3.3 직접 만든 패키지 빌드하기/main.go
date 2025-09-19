package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/~/example3" // 직접 만든 패키지 임포트
)

func main() {
	e := echo.New()                                // 새로운 HTTP 서버 인스턴스 생성
	e.GET("/:number", func(c echo.Context) error { // number라고 부를 특정 시퀀스 캐릭터가 보이면 핸들러 실행
		nstr := c.Param("number")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(example3.IsPrime(n)))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
