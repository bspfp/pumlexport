# pumlexport

PlantUML 파일을 PlantUML Server를 사용해서 이미지 파일로 생성

- [사용법](#사용법)
- [Docker로 PlantUML Server 실행](#docker로-plantuml-server-실행)
  - [WSL2](#wsl2)
  - [Docker Desktop](#docker-desktop)
  - [PlantUML Server](#plantuml-server)

## 사용법

```
pumlexport [-t type] [-o /path/output/file] [-s serverurl] /path/input/puml/file.puml
```

| 옵션         | 기본값                | 설명                                   |
| ------------ | --------------------- | -------------------------------------- |
| -t type      | svg                   | 출력 파일 형식<br />png, svg, txt, pdf |
| -o path      | 입력 파일과 동일 경로 | 출력 파일의 경로                       |
| -s serverurl | http://127.0.0.1:8080 | PlantUML 서버 주소                     |

## Docker로 PlantUML Server 실행

### WSL2

Windows 11 기준

```powershell
wsl --install
wsl --set-default-version 2
```

### Docker Desktop

```powershell
winget install --id Docker.DockerDesktop -e
```

`-e`는 정확한 ID 매칭을 위한 옵션

### PlantUML Server

```powershell
docker run -d -p 8080:8080 plantuml/plantuml-server:latest
```
