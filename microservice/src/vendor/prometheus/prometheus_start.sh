#!/bin/bash
set -eu

#####################################################
##   TARGETS, BASICSECRET, GATEWAYSERVER, SPARAMS
#####################################################

targets=$(env | grep TARGETS | awk -F"=" '{print $2}' | sed 's/,/","/g')
cat > targets.json << EOF
[
   {
     "targets": [ "${targets}" ],
     "labels": {
       "env": "prod",
       "job": "microservice"
     }
   }
]
EOF

basicauth=$(env | grep BASICSECRET | awk -F"=" '{print $2}')
cat > password_file << EOF
${basicauth}
EOF

gatewayserver=$(env | grep GATEWAYSERVER | awk -F"=" '{print $2}')

cat > targets-backend.json << EOF
[
   {
     "targets": [ "${gatewayserver}" ],
     "labels": {
       "env": "prod",
       "job": "microservice-backend"
     }
   }
]
EOF

params=$(env | grep SPARAMS | awk -F"=" '{print $2}' | sed 's/,/","/g')
echo ${params}
sed "s/\${BACKEND_SERVICES}/${params}/g"  ./prometheus_template.yml > ./prometheus.yml

./prometheus --web.listen-address=:8080

#tmp_targets=$(env | grep TARGETS)
#configserver=""
#STR_ARRAY=(`echo $tmp_targets | tr "," "\n"`)
#for x in "${STR_ARRAY[@]}"
#do
#  if [[ $x == configapp* ]]; then
#    configserver=$x
#  fi
#done