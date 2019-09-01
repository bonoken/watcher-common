# watcher-common

watcher 공통 모듈

## project source download
```go
$ git clone https://github.com/bonoken/watcher-common.git
```

## GO Module 적용 방법
```go
// GO Module을 첫 설정
$ go mod init github.com/bonoken/watcher-common
$ go build

// 옵션으로 vendor 명령어를 사용하면 vendor 라는 디렉토리가 생성되며, 의존성 패키지를 vendor디렉토리로 자동으로 
$ go mod vendor 
$ go run -mod vendor main.go
```
* 주의 사항 :
go modules 사용하게 되면 상대경로를 사용 못 함. 프로젝트 내부 상대경로를 변경 "../global" -> "모듈이름/global"

## 모듈 설명
### config
외부 config.yml파일 리딩 모듈, example 적용 참조

### echozap
경량 웹 프레임워크인 echo의 logging을 uber의 zap에 기록을 위한 모듈, example 적용 참조

### gorm
go ORM 연동 모듈, example 적용 참조

### pagination
gorm을 사용한 pagination 기능 모듈, example 적용 참조

### strcase
string case 전환 util, strcase 내부 README 참조

### strutil
string util 모음

### zap
uber의 zap logging 설정 , example 적용 참조


