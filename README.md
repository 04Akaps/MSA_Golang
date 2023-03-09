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
