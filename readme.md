# Go Fiber을 이용한 REST API Boilerpalte

go fiber를 이용한 RESTful한 API 작성해보려고 합니다.
​
제가 node.js기반 백엔드 개발에 익숙해서 npm , yarn등의 패키지 매니저가 익숙한데요 js기반 개발을 많이 하신 분들은 이런 중앙 package 저장소에서 패키지를 관리하는데 go는 이런 중앙 패키지 저장소가 없습니다. package이름이 URL형태로 제공됩니다. github주소가 각 저장소 URL로 import해서 사용할 수 있습니다. npm , yarn을 설치 후 npm install <패키지 이름> , yarn install <패키지 이름> 이렇게 사용하려면 npm , yarn을 설치해야하는데 그럴 필요가 없습니다:)
​
예시)
​

```
module github.com/corgi93/go-fiber-rest-api
​
go 1.19
​
require (
    github.com/arsmn/fiber-swagger/v2 v2.31.1
    github.com/go-playground/validator/v10 v10.10.1
    github.com/gofiber/fiber/v2 v2.43.0
    github.com/gofiber/jwt/v2 v2.2.7
    github.com/golang-jwt/jwt v3.2.2+incompatible
    github.com/google/uuid v1.3.0
    github.com/jackc/pgx/v4 v4.15.0
    github.com/jmoiron/sqlx v1.3.5
    github.com/joho/godotenv v1.4.0
    github.com/stretchr/testify v1.7.0
    github.com/swaggo/swag v1.8.1
)
```

​
require로 github주소에 코드와 버전을 명시해서 종속성 관리 (package.json같은)를 합니다.
​

## go mod init

​
go.mod는 root에 하나 생성되어야 합니다. 그러기 위해선 go mod 명령어로 초기화해줍니다. (npm init같은 것으로 봐주시면 될 것 같습니다.)
​

### go modules를 쓰는 이점

​

- 패키지의 버전관리가 가능해진다
- npm의 package.json처럼 go에서도 사용할 패키지 리스트를 하나의 파일로 명시할 수 있어 코드 deploy시 참조한 패키지를 이 파일에서 go가 설치합니다
  ​

```
$ go mod init github.com/corgi93/go-fiber-api
```

​
go module init시 GO111MODULE 관련  아래 에러가 난다면?
​
```
go: modules disabled by GO111MODULE=off; see 'go help modules'
```
​
해당 명령어로 기존에 off로 되있으면 go 빌드중 $GOPATH에 있는 패키지를 사용한다는 것입니다. on으로 바꿔 빌드 중 $GOPATH 대신 모듈에 있는 패키지를 사용한다고 바꿔줘야합니다.

```
$ go env -w GO111MODULE=on
```


## API 메서드 
- public 
    * GET: /api/v1/books  : 모든 책 정보
    * GET: /api/v1/book/{id} : id에 대한 책 정보 
    * GET: /api/v1/token   : new access token 생성
- private (JWT로 개인화)
    * POST: /api/v1/book : 새 책 생성
    * PATCH : /api/v1/book : 기존 책 업데이트 


## 디렉토리 구조
```
app  : 실제 프로덕트의 business 로직 
    |- controllers      : 기능 컨트롤러 (라우터) 
    |- models           : 비니지스 모델 
    |- queries          : 모델에 대한 쿼리 

pkg  : 프로젝트 util, middleware, configs같은 use case 패키지들
    |- configs          : 구성 기능
    |- middlewares      : 미들웨어
    |- routes           : 프로젝트 route
    |- utils            : utility성 코드들(server starter, generators등)

platform : DB또는 Cache 서버 인스턴스 설정 및 마이그레이션 저장 등의 플랫폼 수준의 논리 
    |- database         : database 설정 기능
    |- migrations       : migration 파일 있는 폴더 (golang-migrate/migrate 툴을 이용)

```

## Makefile
Makefile은 프로그램을 컴파일하고 링크하는 방법을 알려주는 파일입니다
Makefile을 통해서 프로젝트 관리를 용이하게 해주고 프로젝트의 일부 부분이 수정되어 재컴파일이 필요할 경우 이를 쉽게 해줍니다

### Makefile필요성
- 프로그램 개발 시 라인이 길어지면 모듈로 나눠 개발하는데 이 때 입력 파일이 바뀌게 되면 바뀐 파일과 관계가 있는 파일 또한 다시 컴파일 해야하는 불편함이 발생. 이때 Makefile을 만들어 빌드(프로그램 실행 파일 만들기)시 불편함을 줄일 수 있습니다


### migrate
go의 https://github.com/golang-migrate/migrate 해당 툴로 cli로 명령어 한줄로 쉽게 database를 up & down (마이그레이트) 할 수 있습니다. 
```
migrate \
    -path $(PWD)/platform/migrations \
    -database "postgres://postgres:password@localhost/postgres?sslmode=disable" \
    up
```


## docker

### docker network
docker 컨테이너(container)는 격리된 환경에서 돌아가 기본적으로 다른 컨테이너와 통신이 불가능하지만 여러 개의 컨테이너를 하나의 docker 네트워크(network)에 연결 시키면 컨테이너간 네트워킹이 가능해집니다!

```
$ docker network ls
NETWORK ID     NAME          DRIVER    SCOPE
fc1986164fda   bridge        bridge    local
107c70277bdf   host          host      local
fcb2a74932f7   none          null      local
```

* 네트워크 종류
``bridge``, ``host``,``none`` 은 docker 데몬이 실행되면서 디폴트로 생성되는 네트워크로 이런 디폴트 네트워크를 사용하기 보다 사용자가 직접 네트워크를 생서해서 사용하도록 권장합니다.
    - ``bridge`` 네트워크는 하나의 호스트 컴퓨터 내에서 여러 컨테이너들이 서로 소통할 수 있게 해줍니다.
    - ``host`` 네트워크는 컨테이너를 호스트 컴퓨터와 동일한 네트워크에서 컨테이너를 돌리기 위해서 사용합니다.
    - ``overlay`` 네트워크는 여러 호스트에 분산되 돌아가는 컨테이너들 간 네트워킹을 위해서 사용됩니다.

