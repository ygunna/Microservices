#!/bin/bash

exec 2>&1

export PATH=$PATH:/var/vcap/packages/jq/bin
export PATH=$PATH:/var/vcap/packages/cf-cli/bin

set -eu

<% cf = nil; if_link("cf-admin-user") { |link| cf = link } -%>
CF_API_URL='<%= cf ? cf.p("api_url") : p("cf.api_url") %>'
CF_ADMIN_USERNAME='<%= cf ? cf.p("admin_username") : p("cf.username") %>'
CF_ADMIN_PASSWORD='<%= cf ? cf.p("admin_password") : p("cf.password") %>'
<% if cf -%>
mkdir -p /var/vcap/sys/run
cat > /var/vcap/sys/run/cf.crt <<EOF
<%= cf.p("ca_cert") %>
EOF
export SSL_CERT_FILE=/var/vcap/sys/run/cf.crt
<% end -%>
CF_SKIP_SSL_VALIDATION=<%= p("cf.skip_ssl_validation") %>

<%
  broker_url = p("servicebroker.url", nil)
  broker_name = p("servicebroker.name", nil)
  broker_username = p("servicebroker.username", nil)
  broker_password = p("servicebroker.password", nil)
  unless broker_url
    broker = link("servicebroker")
    external_host = broker.p("external_host", "#{broker.instances.first.address}:#{broker.p("broker.port")}")
    protocol      = broker.p("protocol", broker.p("ssl_enabled", false) ? "https" : "http")
    broker_url  ||= "#{protocol}://#{external_host}"
    broker_name ||= broker.p("name")
    broker_username ||= broker.p("username")
    broker_password ||= broker.p("password")
  end
%>
BROKER_NAME='<%= broker_name %>'
BROKER_URL='<%= broker_url %>'
BROKER_USERNAME='<%= broker_username %>'
BROKER_PASSWORD='<%= broker_password %>'

function createOrUpdateServiceBroker() {
  if [[ "$(cf curl /v2/service_brokers\?q=name:${BROKER_NAME} | jq -r .total_results)" == "0" ]]; then
    echo "Service broker '${BROKER_NAME}' does not exist - creating broker"
    cf create-service-broker ${BROKER_NAME} ${BROKER_USERNAME} ${BROKER_PASSWORD} ${BROKER_URL}
  else
    echo "Service broker '${BROKER_NAME}' already exists - updating broker"
    cf update-service-broker ${BROKER_NAME} ${BROKER_USERNAME} ${BROKER_PASSWORD} ${BROKER_URL}
  fi
}

echo "CF_API_URL=${CF_API_URL}"
echo "CF_SKIP_SSL_VALIDATION=${CF_SKIP_SSL_VALIDATION}"
echo "CF_ADMIN_USERNAME=${CF_ADMIN_USERNAME}"
echo "BROKER_NAME=${BROKER_NAME}"
echo "BROKER_URL=${BROKER_URL}"
echo "BROKER_USERNAME=${BROKER_USERNAME}"

if [[ ${CF_SKIP_SSL_VALIDATION} == "true" ]]; then
  cf api ${CF_API_URL} --skip-ssl-validation
else
  cf api ${CF_API_URL}
fi

cf auth \
  ${CF_ADMIN_USERNAME} \
  ${CF_ADMIN_PASSWORD}

createOrUpdateServiceBroker

cf service-access

service_names=($(curl -s -H "X-Broker-Api-Version: 2.10" -u ${BROKER_USERNAME}:${BROKER_PASSWORD} ${BROKER_URL}/v2/catalog | jq -r ".services[].name"))
for service_name in "${service_names[@]}"; do
  cf enable-service-access $service_name
done

cf service-access

curl -X PUT -H "X-Broker-Api-Version: 2.10" -u ${BROKER_USERNAME}:${BROKER_PASSWORD} ${BROKER_URL}/v2/service_instances/-micro -d '{"service_id": "64aca71f-f2e9-4f3d-8e0e-9a3e1e5e3bb8", "plan_id": "334744a3-f12f-4004-a94e-d7132a0d0708", "organization_guid": "", "space_guid": ""}'
#curl -X PUT -H "X-Broker-Api-Version: 2.10" -u ${BROKER_USERNAME}:${BROKER_PASSWORD} ${BROKER_URL}/preapp/v1/create_service_app -d '{"service_id": ["64aca71f-f2e9-4f3d-8e0e-9a3e1e5e3bb6", "64aca71f-f2e9-4f3d-8e0e-9a3e1e5e3bb7", "64aca71f-f2e9-4f3d-8e0e-9a3e1e5e3bb8"]}'

# prometheus
cd /var/vcap/packages/prometheus-binary
cf api ${CF_API_URL} --skip-ssl-validation
cf auth \
  ${CF_ADMIN_USERNAME} \
  ${CF_ADMIN_PASSWORD}
cf target -o org-micro -s space-micro
cf push prometheus-micro -b binary_buildpack -c './prometheus_start.sh' -m 128m --no-start