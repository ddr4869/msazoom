# MSACHAT

### TODO
- 기능별 MSA 분리
    - Pub/Sub 적용 (Kafka?)
    - gRPC로 서비스간 통신
- 1:N Streaming

## 주요 기능 
- **사용자 인증**: 로그인 및 로그아웃 기능.
    - OAuth 적용 (Todo)
- **화상 채팅**: WebRTC를 이용한 실시간 화상 통신.
    - 공개/비공개 방 구분
    - 공개방 랜덤 채팅 기능
    - 채팅 참여자 친구 추가 기능
    - 단체 화상 채팅 기능 (TODO)
    - 음소거/캠 off/Mute 기능 적용 (TODO)
- **채팅방**: 채팅방 생성 및 참여 기능.
    - File/VOD 등 전송 기능 (TODO)
- **친구 목록**: 친구 추가 및 관리 기능.
    - 1:1 채팅
    - 친구 삭제
    - 안 읽은 메세지가 몇개 왔는지 확인 (TODO)
    - 팔로우 요청 보내기 및 수락하기 기능
- **팔로우 요청**: 

## 사용 기술

- **WebRTC**: 실시간 화상 통신을 위해 사용.
- **WebSocket**: 실시간 텍스트 메시징을 위해 사용(Friend).
- **Go(Gin, Ent)**
- **Typescript(NextJS)**

## SetUp

1. Server
    ```bash
    go mod init
    go mod tidy
    ```
    ```bash
    go run main
    ```

2. Client
    ```bash
    git clone git@github.com:ddr4869/msazoom.git
    cd msazoom
    ```
    ```bash
    npm install
    ```
    ```bash
    npm run build
    ```
    http://localhost:3000

3. Deployment(daemon)
    ```bash
    docker-compose up -d
    ``` 


## 사용 방법

1. **로그인** 또는 **회원가입**하여 채팅 보드에 접근.
2. **채팅방 생성 또는 참여**를 채팅 목록에서 선택.
3. **친구 이름을 클릭**하여 화상 채팅 시작.
4. **채팅방에서 메시지 주고받기**.

## Picture

### 채팅 보드
![채팅 보드](./docs/picture/dashboard.png)

### 화상 채팅
![화상 채팅](./docs/picture/chat.png)

### 텍스트 채팅
![텍스트 채팅](./docs/picture/message.png)