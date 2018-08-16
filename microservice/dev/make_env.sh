#!/bin/bash

set -e
set -u
# set -x

DIR="$(cd $(dirname ${BASH_SOURCE[0]})/../ && pwd)"
sampleapp=$DIR/src/crossent/micro/broker/config/assets/sampleapp
springapp="$(cd $DIR/../../MsXpert-spring-cloud && pwd)"
configapp=""
registryapp=""
gatewayapp=""

createCheck(){
    cf marketplace
    count=$(cf marketplace | grep server | wc -l)
    if [ $count != 0 ]; then
        echo "delete first"
        exit 1
    fi
}
createServiceBroker(){
    bosh -e vbox -d msxpert run-errand broker-registrar
}
createService(){
    cf create-service micro-config-server standard config-server
    cf create-service micro-registry-server standard registry-server
#    cf create-service micro-gateway-server standard gateway-server
}
startApp(){
    configapp=$(cf apps | grep ^configapp | awk '{print $1}')
    cf start $configapp
    sleep 1

    registryapp=$(cf apps | grep ^registryapp | awk '{print $1}')
    cf start $registryapp
    sleep 1

    gatewayapp=$(cf apps | grep ^gatewayapp | awk '{print $1}')
    cf start $gatewayapp
    sleep 1

    cf start front
    sleep 1

    cf start back    
    sleep 1
}
pushAppNostart(){
    echo $sampleapp
    pushd $sampleapp/front
        cf push --no-start
    popd

    pushd $sampleapp/back
        cf push --no-route --no-start
    popd
}
createNetworkPolicy(){
    cf add-network-policy front --destination-app $gatewayapp --protocol tcp --port 8080
    cf add-network-policy $gatewayapp --destination-app back --protocol tcp --port 8080
}
resultLog(){
    cf marketplace
    cf services
    cf apps
    cf list-access
}
deleteApp(){
    apps=$(cf apps | grep com$ | awk '{print $1}')
    for a in ${apps}
    do 
        cf delete $a -f
    done

    cf delete back -f
}
deleteServiceBroker(){
    bosh -e vbox -d msxpert run-errand broker-deregistrar
}
springBuild(){
    echo $springapp
    pushd $springapp/config
        mvn -DskipTests=true package
        cp target/config-0.0.1-SNAPSHOT.jar $DIR/src/crossent/micro/broker/config/assets/configapp
    popd

    pushd $springapp/registry
        mvn -DskipTests=true package
        cp target/registry-0.0.1-SNAPSHOT.jar $DIR/src/crossent/micro/broker/config/assets/registryapp
    popd   

    pushd $springapp/gateway
        mvn -DskipTests=true package
        cp target/gateway-0.0.1-SNAPSHOT.jar $DIR/src/crossent/micro/broker/config/assets/gatewayapp
    popd      

    pushd $springapp/front
        mvn -DskipTests=true package
        cp target/front-0.0.1-SNAPSHOT.jar $DIR/src/crossent/micro/broker/config/assets/sampleapp/front
    popd 

    pushd $springapp/back
        mvn -DskipTests=true package
        cp target/back-0.0.1-SNAPSHOT.jar $DIR/src/crossent/micro/broker/config/assets/sampleapp/back
    popd         
}
msxpertDeploy(){
    pushd ../
        bosh -e vbox update-cloud-config cloud-config-lite.yml -n
        bosh -e vbox create-release --name msxpert --force -n && bosh -e vbox upload-release --name msxpert && bosh -e vbox -d msxpert deploy microservice-msxpert.yml -n    
    popd
    bosh -e vbox -d msxpert vms
}

main(){
    if [ "$1" == "-i" ]; then
        echo "======[msxpert bosh deploy]======"
#        springBuild
        msxpertDeploy
    elif [ "$1" == "-d" ]; then
        echo "======[delete env]======"
        deleteApp
        deleteServiceBroker
        resultLog
    else
        echo $ "======[create env]======"
        createCheck
        createServiceBroker
        createService
        pushAppNostart
        startApp
        createNetworkPolicy
        resultLog
    fi

}

main "$*"