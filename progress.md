# NOTAM Viewer 개발 진행 로그

## 0. 프로젝트 개요
- 백엔드: Go 기반 REST API 서버
- 프론트엔드: React Native (KakaoMap)
- 데이터베이스: PostgreSQL + PostGIS
- 캐시: Redis
- 통신: HTTPS 기반 REST API

## 1. 초기 셋업

### 1-1. 리포지토리 구조 확정 (계획)
- [ ] `backend/` - Go API 서버
- [ ] `mobile/` - React Native 앱
- [ ] `infra/` - docker-compose, DB, Redis
- [ ] `docs/` - 설계 문서, API 스펙

### 1-2. 기본 환경 정리
- [ ] Go 1.25.5 설치 확인
- [ ] Node.js / yarn(or npm) 설치
- [ ] React Native CLI 및 Android/iOS 개발 환경 준비
- [ ] Docker & Docker Compose 설치

## 2. 인프라/데이터베이스

### 2-1. PostgreSQL + PostGIS
- [ ] `infra/postgres/docker-compose.yml` 작성
- [ ] NOTAM 관련 스키마 초안 설계
- [ ] 공간 인덱스(GIST) 적용

### 2-2. Redis
- [ ] `infra/redis/docker-compose.yml` 작성
- [ ] 캐싱 전략(키 설계, TTL 정책) 정의

## 3. 백엔드(Go)

### 3-1. 프로젝트 부트스트랩
- [x] `backend/` Go module 초기화 (`go mod init`)
- [ ] 기본 폴더 구조 설계 (handler, service, repository, config, etc.)
- [ ] Health Check API 구현 (`GET /health`)

### 3-2. NOTAM 도메인 설계
- [ ] NOTAM 엔티티/DTO 정의
- [ ] 외부 NOTAM 소스 조사 및 연동 방식 결정
- [ ] PostGIS 기반 조회 쿼리 초안 작성 (공항별, 영역별, 시간 범위별)

### 3-3. REST API 설계
- [ ] API 엔드포인트 목록 정의
- [ ] Swagger/OpenAPI 초안 작성
- [ ] 기본 에러 응답 포맷 결정

## 4. 프론트엔드(React Native + KakaoMap)

### 4-1. 프로젝트 생성
- [ ] `mobile/` React Native 프로젝트 초기화
- [ ] 기본 네비게이션 구조(Stack/Tab) 설정
- [ ] 환경변수/빌드 설정 분리(dev, prod)

### 4-2. KakaoMap 연동
- [ ] Kakao 개발자 센터 앱 키 발급
- [ ] iOS/Android KakaoMap SDK 설정
- [ ] 지도 초기 렌더링 및 현재 위치 표시

### 4-3. NOTAM 조회 UI
- [ ] 공항 검색 화면
- [ ] 지도 위 NOTAM 마커/폴리곤 표시
- [ ] 상세 NOTAM 정보 화면

## 5. 품질/운영

### 5-1. 테스트 & 로깅
- [ ] 백엔드 단위/통합 테스트 환경 구성
- [ ] 기본 로깅 포맷 및 레벨 정의
- [ ] 간단한 모니터링/헬스체크 구성

### 5-2. 배포 전략(초안)
- [ ] 개발용 Docker Compose 환경 정리
- [ ] 운영 배포 후보 환경 정리 (예: AWS/GCP, k8s 등)

## 변경 이력

- 2025-12-04
  - 초기 [`README.md`](README.md:1) 작성
  - 초기 [`progress.md`](progress.md:1) 생성 및 1차 TODO 정리
  - `backend/` [`go.mod`](backend/go.mod:1) 생성 (Go 1.25.5, module 초기화)