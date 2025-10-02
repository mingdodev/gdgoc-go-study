```bash
go mod primechecker # 새로운 모듈 정의
go get github.com/otiai10/primes # 서드파티 모듈을 포함하도록 go mod 수정됨
go build . # 소스코드 컴파일 (모듈 이름으로 파일 생성됨)
```