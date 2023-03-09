# MSA_Golang

Golang을 활용하여 MSA 구조 서버 만들기

# NewEvent Request Sample

```
{
"name" : "뭐냐",
"duration" : "300s",
"start_date" : 3,
"end_date" : 10,
"location" : {
"name" : "seoul",
"address" : "어디로 찍냐",
"country" : "seoul",
"open_time" : 3,
"close_time" : 3,
"halls" : [
{"name" : "hojin","location" : "house", "capaciry" : 5},
{"name" : "123","location" : "123", "capaciry" : 10000}
]
}
}

```

# openssl을 통해서 개인 인증서 생성

openssl req -x509 -newkey rsa:2048 -nodes keyout key.pem -out cert.pem -days 365

- 간단하게 바로 SSL/TLS인증서를 만드는 명령어 입니다.

1. -x509 : 셀프 서명 인증서를 생성하도록 지정합니다.
2. -newkey res:2048 : RSA알고리즘을 사용하여 2048비트 개인키를 생성합니다.
3. -nodes : 개인키를 암호화하지 않도록 지정합니다. 만약 해당 옵션을 제외하면, 키 생성 후 암호를 지정해야 합니다.
4. -keyout key.pem : 개인키를 `key.pem`으로 생성합니다.
5. -out cert.pem : 셀프 서명 인증서(CSR)을 cert.pem에 지정 합니다.
6. -days 365 : 유효기간을 1년으로 설정합니다.

# docker

1. docker build -t name/repository:tag path

docker를 빌드하는 기본적인 명령어 입니다.

이미지의 이름 + 사용될 레파지토리 + 태그 + 포함할 파일을 입력합니다.

- 예시는 다음과 같습니다.
- docker build -t hojin/msa:1.0 ./
  // 해당 명령어는 로컬에 빌드하는 명령어 입니다.

2. docker push name/repository:tag

해당 명령어는 docker hub에 image를 업로드 하는 명령어 입니다.
기존에 구워낸 -> build해서 이미지가 생성된 이미지를 이제 hub에 업로드 하게 됩니다.

# Jenkins

Jenkins에서 github나 EC에 접근하기 위해서 SSH를 사용할 것이고 다음과 같은 명령어로 생성 하면 됩니다.

- ssh-keygen -t ed25519 -a 100 -f jenkins-ssh-key

해당 명령어로 key를 두가지 생성해 줍니다.

- \*.pub => public Key
- 나머지 => private Key

그러면 이제 Jenkins에는 private Key를 등록해 두고, Github에는 Public Key를 등록하게 된다면 서로 암호화 통신이 가능합니다.

- Github의 경우에는 settings 에서 Deply Keys에 등록

Jenkins의 경우에는 Credentials페이지로 접속하여 Private Key를 등록하면 됩니다.

<img src="./img/github PubKey.png">
<img src="./img/Jenkins PrivateKey.png">
