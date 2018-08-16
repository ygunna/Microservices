# MsXpert

## org, space 생성
```
cf create-org org-micro
cf create-space space-micro -o org-micro
```

## UAA Client 등록

UAA에 Microservice Client를 등록
```
uaac client add micro --name micro -s micro-secret \
   --authorities "oauth.login,scim.write,clients.read,scim.userids,password.write,clients.secret,clients.write,uaa.admin,scim.read,doppler.firehose" \
   --authorized_grant_types "authorization_code,client_credentials,password,refresh_token" \
   --scope "cloud_controller.read,cloud_controller.write,openid,cloud_controller.admin,scim.read,scim.write,doppler.firehose,uaa.user,routing.router_groups.read,uaa.admin,password.write" \
   --redirect_uri "https://uaa.bosh-lite.com/login"
```

## config spring app 복사
```
$ cp config-0.0.1-SNAPSHOT.jar microservice/src/crossent/micro/broker/config/assets/configapp
```

## registry spring app 복사
```
cp registry-0.0.1-SNAPSHOT.jar microservice/src/crossent/micro/broker/config/assets/registryapp
```

## gateway spring app 복사
```
cp gateway-0.0.1-SNAPSHOT.jar microservice/src/crossent/micro/broker/config/assets/gatewayapp
```

## microservice bosh release 설치 (bosh2 기준)
```
bosh -e vbox upload-release https://bosh.io/d/github.com/cloudfoundry/postgres-release?v=23

bosh -e vbox update-cloud-config cloud-config-azure.yml
bosh -e vbox create-release --name msxpert --force
bosh -e vbox upload-release --name msxpert
bosh -e vbox -d msxpert deploy microservice-msxpert-azure.yml

bosh -e vbox -d msxpert run-errand broker-registrar
```