# NOTAM Viewer

Golang 기반 backend와 React Native(KakaoMap) 기반 frontend로 구현하는 항공 NOTAM(Notice To Air Missions) 조회 애플리케이션.

## 목표

- 국내/국제 공항 NOTAM을 지도로 직관적으로 조회
- 특정 공항, 항로, FIR 기준 필터링 및 검색
- 캐싱(Redis)과 공간쿼리(PostGIS)를 활용한 빠른 응답
- 모바일 환경에서의 실시간 운항 지원

## 전체 아키텍처(초안)

```mermaid
flowchart LR
  RN[React Native App(KakaoMap)] -->|REST API| BE[Go Backend]
  BE -->|SQL/Spatial| DB[(PostgreSQL + PostGIS)]
  BE -->|Cache| R[(Redis)]
  BE -->|외부 API| NOTAM[NOTAM Provider]
```

## 기술 스택

- Backend: Go (Golang), Echo 또는 Gin, gorm 등 (추후 결정)
- DB: PostgreSQL + PostGIS
- Cache: Redis
- Frontend: React Native, Kakao Map SDK
- Infra: Docker Compose(개발용), (추후 Kubernetes 또는 클라우드 배포 검토)

## 디렉터리 구조(계획)

```text
.
├── backend/          # Go 기반 REST API 서버
├── mobile/           # React Native 앱 (KakaoMap)
├── infra/            # docker-compose, DB, Redis 설정
├── docs/             # 설계 문서, API 명세
├── progress.md       # 개발 진행 상황 로그
└── README.md
```

## 기능 요구사항(초안)

1. 공항 NOTAM 조회
   - ICAO/IATA 코드 기반 검색
   - 유효 시간대(TIME FROM/UNTIL) 필터링
2. 지도 기반 조회
   - KakaoMap 상에 NOTAM 위치(포인트/폴리곤) 표시
   - 줌 레벨에 따른 클러스터링 또는 집계
3. 즐겨찾기 & 최근 조회
   - 자주 사용하는 공항 리스트 관리
   - 최근 조회한 공항/NOTAM 이력
4. 오프라인/약전 구간 대응(선택)
   - 최근 조회 결과를 로컬에 캐시

## 개발 단계 계획(요약)

1. 요구사항/도메인 정리 및 데이터 모델 설계
2. PostgreSQL + PostGIS, Redis 로컬 개발 환경 구성
3. Go 기반 REST API 스켈레톤 프로젝트 생성
4. 기본 NOTAM 조회 API 구현 및 테스트
5. React Native 앱 초기 화면 + KakaoMap 연동
6. 지도 상 NOTAM 마커/도형 렌더링
7. 인증/권한(필요 시) 및 배포 파이프라인 구성

## 개발 환경 준비(개요)

- Go 1.25.5
- Node.js / yarn 또는 npm
- React Native 개발 환경 (iOS/Android SDK)
- Docker & Docker Compose

세부적인 설치 및 실행 방법은 [`progress.md`](progress.md)에 단계별로 정리할 예정입니다.