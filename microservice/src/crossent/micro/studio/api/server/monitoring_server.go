package server

import (
	"net/http"

	"encoding/json"
	"strconv"
	"fmt"
	"crossent/micro/studio/domain"
	"strings"
	"time"
	"crypto/tls"
	"bytes"
	"errors"
	"code.cloudfoundry.org/lager"
	"io/ioutil"
	"text/template"
	"net"
	"encoding/base64"
	"io"
)

const JVM_MICROMETER_JSON = `
{
  "dashboard": {
  "__inputs": [],
  "overwrite": false,
  "__requires": [
    {
      "type": "grafana",
      "id": "grafana",
      "name": "Grafana",
      "version": "4.6.1"
    },
    {
      "type": "panel",
      "id": "graph",
      "name": "Graph",
      "version": ""
    },
    {
      "type": "datasource",
      "id": "prometheus",
      "name": "Prometheus",
      "version": "1.0.0"
    },
    {
      "type": "panel",
      "id": "singlestat",
      "name": "Singlestat",
      "version": ""
    }
  ],
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "limit": 100,
        "name": "Annotations & Alerts",
        "showIn": 0,
        "type": "dashboard"
      },
      {
        "datasource": "{{.DsPrometheus}}",
        "enable": true,
        "expr": "resets(process_uptime_seconds{application=\"$application\", instance=\"$instance\"}[1m]) > 0",
        "iconColor": "rgba(255, 96, 96, 1)",
        "name": "Restart Detection",
        "showIn": 0,
        "step": "1m",
        "tagKeys": "restart-tag",
        "textFormat": "uptime reset",
        "titleFormat": "Restart"
      }
    ]
  },
  "description": "Dashboard for Micrometer instrumented applications (Java, Spring Boot, Micronaut)",
  "editable": false,
  "gnetId": 4701,
  "graphTooltip": 1,
  "hideControls": false,
  "id": null,
  "links": [],
  "refresh": "1m",
  "rows": [
    {
      "collapse": false,
      "height": "100px",
      "panels": [
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(245, 54, 54, 0.9)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(50, 172, 45, 0.97)"
          ],
          "datasource": "{{.DsPrometheus}}",
          "decimals": 1,
          "editable": false,
          "error": false,
          "format": "s",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "",
          "id": 63,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "70%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 3,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "process_uptime_seconds{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "",
              "metric": "",
              "refId": "A",
              "step": 14400
            }
          ],
          "thresholds": "",
          "title": "Uptime",
          "transparent": false,
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(245, 54, 54, 0.9)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(50, 172, 45, 0.97)"
          ],
          "datasource": "{{.DsPrometheus}}",
          "decimals": null,
          "editable": false,
          "error": false,
          "format": "dateTimeAsIso",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "",
          "id": 92,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "70%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 3,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "process_start_time_seconds{application=\"$application\", instance=\"$instance\"}*1000",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "",
              "metric": "",
              "refId": "A",
              "step": 14400
            }
          ],
          "thresholds": "",
          "title": "Start time",
          "transparent": false,
          "type": "singlestat",
          "valueFontSize": "70%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "{{.DsPrometheus}}",
          "decimals": 2,
          "editable": false,
          "error": false,
          "format": "percent",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "id": 65,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "70%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 3,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "sum(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"heap\"})*100/sum(jvm_memory_max_bytes{application=\"$application\",instance=\"$instance\", area=\"heap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "",
              "refId": "A",
              "step": 14400
            }
          ],
          "thresholds": "70,90",
          "title": "Heap used",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "{{.DsPrometheus}}",
          "decimals": 2,
          "editable": false,
          "error": false,
          "format": "percent",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "id": 75,
          "interval": null,
          "links": [],
          "mappingType": 2,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "70%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            },
            {
              "from": "-99999999999999999999999999999999",
              "text": "N/A",
              "to": "0"
            }
          ],
          "span": 3,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "sum(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"nonheap\"})*100/sum(jvm_memory_max_bytes{application=\"$application\",instance=\"$instance\", area=\"nonheap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "",
              "refId": "A",
              "step": 14400
            }
          ],
          "thresholds": "70,90",
          "title": "Non-Heap used",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            },
            {
              "op": "=",
              "text": "x",
              "value": ""
            }
          ],
          "valueName": "current"
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "Quick Facts",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 111,
          "legend": {
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(http_server_requests_seconds_count{application=\"$application\", instance=\"$instance\"}[1m]))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "HTTP",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Rate",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "decimals": null,
              "format": "ops",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {
            "HTTP": "#890f02",
            "HTTP - 5xx": "#bf1b00"
          },
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 112,
          "legend": {
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(http_server_requests_seconds_count{application=\"$application\", instance=\"$instance\", status=~\"5..\"}[1m]))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "HTTP - 5xx",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Errors",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "decimals": null,
              "format": "ops",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 113,
          "legend": {
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(http_server_requests_seconds_sum{application=\"$application\", instance=\"$instance\", status!~\"5..\"}[1m]))/sum(rate(http_server_requests_seconds_count{application=\"$application\", instance=\"$instance\", status!~\"5..\"}[1m]))",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "HTTP - AVG",
              "refId": "A"
            },
            {
              "expr": "max(http_server_requests_seconds_max{application=\"$application\", instance=\"$instance\", status!~\"5..\"})",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "HTTP - MAX",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Duration",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "s",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "description": "",
          "fill": 1,
          "id": 119,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "tomcat_threads_busy{application=\"$application\", instance=\"$instance\"} or tomcat_threads_busy_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "TOMCAT - BSY",
              "refId": "A"
            },
            {
              "expr": "tomcat_threads_current{application=\"$application\", instance=\"$instance\"} or tomcat_threads_current_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "TOMCAT - CUR",
              "refId": "B"
            },
            {
              "expr": "tomcat_threads_config_max{application=\"$application\", instance=\"$instance\"} or tomcat_threads_config_max_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "TOMCAT - MAX",
              "refId": "C"
            },
            {
              "expr": "jetty_threads_busy{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "JETTY - BSY",
              "refId": "D"
            },
            {
              "expr": "jetty_threads_current{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "JETTY - CUR",
              "refId": "E"
            },
            {
              "expr": "jetty_threads_config_max{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "JETTY - MAX",
              "refId": "F"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Utilisation",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "I/O Overview",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 24,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"heap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\", area=\"heap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "committed",
              "refId": "B",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_max_bytes{application=\"$application\", instance=\"$instance\", area=\"heap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "max",
              "refId": "C",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "JVM Heap",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 25,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"nonheap\"})",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\", area=\"nonheap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "committed",
              "refId": "B",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_max_bytes{application=\"$application\", instance=\"$instance\", area=\"nonheap\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "max",
              "refId": "C",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "JVM Non-Heap",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 26,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "committed",
              "refId": "B",
              "step": 2400
            },
            {
              "expr": "sum(jvm_memory_max_bytes{application=\"$application\", instance=\"$instance\"})",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "max",
              "refId": "C",
              "step": 2400
            },
            {
              "expr": "process_memory_vss_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": true,
              "intervalFactor": 2,
              "legendFormat": "vss",
              "metric": "",
              "refId": "D",
              "step": 2400
            },
            {
              "expr": "process_memory_rss_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "rss",
              "refId": "E",
              "step": 2400
            },
            {
              "expr": "process_memory_pss_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "pss",
              "refId": "F",
              "step": 2400
            },
            {
              "expr": "process_memory_swap_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "swap",
              "refId": "G",
              "step": 2400
            },
            {
              "expr": "process_memory_swappss_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "swappss",
              "refId": "H",
              "step": 2400
            },
            {
              "expr": "process_memory_pss_bytes{application=\"$application\", instance=\"$instance\"} + process_memory_swap_bytes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "phys (pss+swap)",
              "refId": "I",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "JVM Total",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 86,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "(process_memory_pss_bytes{application=\"$application\", instance=\"$instance\"} + process_memory_swap_bytes{application=\"$application\", instance=\"$instance\"}  - on(application,instance) sum(jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\"})  by(application,instance)) >= 0",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "native",
              "metric": "",
              "refId": "A",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "JVM Native Memory",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "JVM Memory",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 106,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "system_cpu_usage{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "system",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "process_cpu_usage{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "process",
              "refId": "B"
            },
            {
              "expr": "avg_over_time(process_cpu_usage{application=\"$application\", instance=\"$instance\"}[1h])",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "process-1h",
              "refId": "C"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "CPU",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 1,
              "format": "percentunit",
              "label": "",
              "logBase": 1,
              "max": "1",
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 93,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "system_load_average_1m{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "system-1m",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "",
              "format": "time_series",
              "intervalFactor": 2,
              "refId": "B"
            },
            {
              "expr": "system_cpu_count{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "cpu",
              "refId": "C"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Load",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 1,
              "format": "short",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 32,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_threads_live{application=\"$application\", instance=\"$instance\"} or jvm_threads_live_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "live",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "jvm_threads_daemon{application=\"$application\", instance=\"$instance\"} or jvm_threads_daemon_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "daemon",
              "metric": "",
              "refId": "B",
              "step": 2400
            },
            {
              "expr": "jvm_threads_peak{application=\"$application\", instance=\"$instance\"} or jvm_threads_peak_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "peak",
              "refId": "C",
              "step": 2400
            },
            {
              "expr": "process_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "process",
              "refId": "D",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Threads",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 0,
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {
            "blocked": "#bf1b00",
            "new": "#fce2de",
            "runnable": "#7eb26d",
            "terminated": "#511749",
            "timed-waiting": "#c15c17",
            "waiting": "#eab839"
          },
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 124,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_threads_states_threads{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "{{.State}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Thread States",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {
            "debug": "#1F78C1",
            "error": "#BF1B00",
            "info": "#508642",
            "trace": "#6ED0E0",
            "warn": "#EAB839"
          },
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "height": "",
          "id": 91,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "hideEmpty": false,
            "hideZero": false,
            "max": true,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": true,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [
            {
              "alias": "error",
              "yaxis": 1
            },
            {
              "alias": "warn",
              "yaxis": 1
            }
          ],
          "spaceLength": 10,
          "span": 9,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "increase(logback_events_total{application=\"$application\", instance=\"$instance\"}[1m])",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "{{.Level}}",
              "metric": "",
              "refId": "A",
              "step": 1200
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Log Events (1m)",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "transparent": false,
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 0,
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 61,
          "legend": {
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "process_open_fds{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "open",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "process_max_fds{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "max",
              "metric": "",
              "refId": "B",
              "step": 2400
            },
            {
              "expr": "process_files_open{application=\"$application\", instance=\"$instance\"} or process_files_open_files{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "open",
              "refId": "C"
            },
            {
              "expr": "process_files_max{application=\"$application\", instance=\"$instance\"} or process_files_max_files{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "max",
              "refId": "D"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "File Descriptors",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 0,
              "format": "short",
              "label": null,
              "logBase": 10,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "JVM Misc",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 3,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "minSpan": 4,
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "repeat": "jvm_memory_pool_heap",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_heap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 1800
            },
            {
              "expr": "jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_heap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "commited",
              "metric": "",
              "refId": "B",
              "step": 1800
            },
            {
              "expr": "jvm_memory_max_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_heap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "max",
              "metric": "",
              "refId": "C",
              "step": 1800
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$jvm_memory_pool_heap",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": "persistence_counts",
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "JVM Memory Pools (Heap)",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 78,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": true,
            "max": true,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "minSpan": 4,
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "repeat": "jvm_memory_pool_nonheap",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_nonheap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 1800
            },
            {
              "expr": "jvm_memory_committed_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_nonheap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "commited",
              "metric": "",
              "refId": "B",
              "step": 1800
            },
            {
              "expr": "jvm_memory_max_bytes{application=\"$application\", instance=\"$instance\", id=\"$jvm_memory_pool_nonheap\"}",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "max",
              "metric": "",
              "refId": "C",
              "step": 1800
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$jvm_memory_pool_nonheap",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "mbytes",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "JVM Memory Pools (Non-Heap)",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 98,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "rate(jvm_gc_pause_seconds_count{application=\"$application\", instance=\"$instance\"}[1m])",
              "format": "time_series",
              "hide": false,
              "intervalFactor": 2,
              "legendFormat": "{{.Action}} ({{.Cause}})",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Collections",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "ops",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 101,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "rate(jvm_gc_pause_seconds_sum{application=\"$application\", instance=\"$instance\"}[1m])/rate(jvm_gc_pause_seconds_count{application=\"$application\", instance=\"$instance\"}[1m])",
              "format": "time_series",
              "hide": false,
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "avg {{.Action}} ({{.Cause}})",
              "refId": "A"
            },
            {
              "expr": "jvm_gc_pause_seconds_max{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "hide": false,
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "max {{.Action}} ({{.Cause}})",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Pause Durations",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "s",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "fill": 1,
          "id": 99,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "rate(jvm_gc_memory_allocated_bytes_total{application=\"$application\", instance=\"$instance\"}[1m])",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "allocated",
              "refId": "A"
            },
            {
              "expr": "rate(jvm_gc_memory_promoted_bytes_total{application=\"$application\", instance=\"$instance\"}[1m])",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "promoted",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Allocated/Promoted",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "Garbage Collection",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 37,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_classes_loaded{application=\"$application\", instance=\"$instance\"} or jvm_classes_loaded_classes{application=\"$application\", instance=\"$instance\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "loaded",
              "metric": "",
              "refId": "A",
              "step": 1200
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Classes loaded",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 38,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "delta(jvm_classes_loaded{application=\"$application\",instance=\"$instance\"}[5m]) or delta(jvm_classes_loaded_classes{application=\"$application\",instance=\"$instance\"}[5m])",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 2,
              "legendFormat": "delta",
              "metric": "",
              "refId": "A",
              "step": 1200
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Class delta (5m)",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "ops",
            "short"
          ],
          "yaxes": [
            {
              "decimals": null,
              "format": "short",
              "label": "",
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "Classloading",
      "titleSize": "h6"
    },
    {
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 33,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_buffer_memory_used_bytes{application=\"$application\", instance=\"$instance\", id=\"direct\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "jvm_buffer_total_capacity_bytes{application=\"$application\", instance=\"$instance\", id=\"direct\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "capacity",
              "metric": "",
              "refId": "B",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Direct Buffers",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 83,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_buffer_count{application=\"$application\", instance=\"$instance\", id=\"direct\"} or jvm_buffer_count_buffers{application=\"$application\", instance=\"$instance\", id=\"direct\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "count",
              "metric": "",
              "refId": "A",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Direct Buffers",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 0,
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 85,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_buffer_memory_used_bytes{application=\"$application\", instance=\"$instance\", id=\"mapped\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "used",
              "metric": "",
              "refId": "A",
              "step": 2400
            },
            {
              "expr": "jvm_buffer_total_capacity_bytes{application=\"$application\", instance=\"$instance\", id=\"mapped\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "capacity",
              "metric": "",
              "refId": "B",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Mapped Buffers",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "{{.DsPrometheus}}",
          "editable": false,
          "error": false,
          "fill": 1,
          "grid": {
            "leftLogBase": 1,
            "leftMax": null,
            "leftMin": null,
            "rightLogBase": 1,
            "rightMax": null,
            "rightMin": null
          },
          "id": 84,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "span": 3,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "jvm_buffer_count{application=\"$application\", instance=\"$instance\", id=\"mapped\"} or jvm_buffer_count_buffers{application=\"$application\", instance=\"$instance\", id=\"mapped\"}",
              "format": "time_series",
              "intervalFactor": 2,
              "legendFormat": "count",
              "metric": "",
              "refId": "A",
              "step": 2400
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Mapped Buffers",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "type": "graph",
          "x-axis": true,
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "y-axis": true,
          "y_formats": [
            "short",
            "short"
          ],
          "yaxes": [
            {
              "decimals": 0,
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": 0,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": true,
      "title": "Buffer Pools",
      "titleSize": "h6"
    }
  ],
  "schemaVersion": 14,
  "style": "light",
  "tags": [],
  "templating": {
    "list": [
      {
        "allValue": null,
        "current": {},
        "datasource": "{{.DsPrometheus}}",
        "hide": 0,
        "includeAll": false,
        "label": "Application",
        "multi": false,
        "name": "application",
        "options": [],
        "query": "label_values(application)",
        "refresh": 2,
        "regex": "",
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allFormat": "glob",
        "allValue": null,
        "current": {},
        "datasource": "{{.DsPrometheus}}",
        "hide": 0,
        "includeAll": false,
        "label": "Instance",
        "multi": false,
        "multiFormat": "glob",
        "name": "instance",
        "options": [],
        "query": "label_values(jvm_memory_used_bytes{application=\"$application\"}, instance)",
        "refresh": 2,
        "regex": "",
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allFormat": "glob",
        "allValue": null,
        "current": {},
        "datasource": "{{.DsPrometheus}}",
        "hide": 0,
        "includeAll": true,
        "label": "JVM Memory Pools Heap",
        "multi": false,
        "multiFormat": "glob",
        "name": "jvm_memory_pool_heap",
        "options": [],
        "query": "label_values(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"heap\"},id)",
        "refresh": 1,
        "regex": "",
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allFormat": "glob",
        "allValue": null,
        "current": {},
        "datasource": "{{.DsPrometheus}}",
        "hide": 0,
        "includeAll": true,
        "label": "JVM Memory Pools Non-Heap",
        "multi": false,
        "multiFormat": "glob",
        "name": "jvm_memory_pool_nonheap",
        "options": [],
        "query": "label_values(jvm_memory_used_bytes{application=\"$application\", instance=\"$instance\", area=\"nonheap\"},id)",
        "refresh": 1,
        "regex": "",
        "sort": 2,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allValue": null,
        "current": {},
        "datasource": "{{.DsPrometheus}}",
        "definition": "",
        "hide": 0,
        "includeAll": false,
        "label": "Job",
        "multi": false,
        "name": "job",
        "options": [],
        "query": "label_values(job)",
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }
    ]
  },
  "time": {
    "from": "now-24h",
    "to": "now"
  },
  "timepicker": {
    "now": true,
    "refresh_intervals": [
      "30s",
      "1m",
      "5m"
    ],
    "time_options": [
      "5m",
      "15m"
    ]
  },
  "timezone": "browser",
  "title": "{{.Title}}",
  "uid": "{{.Uid}}",
  "tags": [ "auto" ],
  "version": 0
},
"folderId": 0
}
`


func (s *Server) GetMicroserviceMonitoring(w http.ResponseWriter, r *http.Request) {
	logger := s.logger.Session("monitoring")
	logger.Debug("GetMicroserviceMonitoring")

	refresh := r.URL.Query().Get("refresh")

	localip := "";
	extUrl := strings.Split(s.uaa.ExternalURL, ":")
	if len(extUrl) >= 2 {
		localip = fmt.Sprintf("%s:%s", extUrl[0], extUrl[1])
	}

	id := r.FormValue(":id")

	idint, _ := strconv.Atoi(id)
	view, err := s.repositoryFactory.View().GetMicroservice(idint)

	if b := s.access(r, view.SpaceGuid); !b {
		logger.Error("no auth", fmt.Errorf("no auth"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no auth"))
		return
	}


	token, err := s.uaa.GetAuthToken()
	if err != nil {
		logger.Error("failed cf get auth token", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	viewApps, err := s.repositoryFactory.View().ListMicroserviceAppApp(view.ID)

	urls := []string{}
	backends := []string{}
	gwurl := ""
	basicsecret := ""
	isMonitoring := false
	monitoringAppGuid := ""
	for _, viewApp := range viewApps {
		summary := s.GetAppSummary(viewApp.AppGuid, token)

		for _, route := range summary.Routes {
			//urls = append(urls, fmt.Sprintf("%s:%s.%s", summary.Name, route.Host, route.Domain.Name))
			if viewApp.Essential != string(domain.MsBackEnd) {
				urls = append(urls, fmt.Sprintf("%s.%s", route.Host, route.Domain.Name))
			}

			if viewApp.Essential == string(domain.MsGateway) {
				gwurl = fmt.Sprintf("%s.%s", route.Host, route.Domain.Name)
			}
		}

		if viewApp.Essential == string(domain.MsBackEnd) {
			backends = append(backends, summary.Name)
		}
		if viewApp.Essential == domain.CONFIG_NAME {
			basicsecret = summary.Environment["monitoring-basic-password"]
		}
		if viewApp.Essential == string(domain.MsMonitoring) {
			isMonitoring = true
			monitoringAppGuid = viewApp.AppGuid
		}
	}

	urlstr := strings.Join(urls, ",")
	backendstr := strings.Join(backends, ",")
	appEnv := make(map[string]interface{})
	appEnv["TARGETS"] = urlstr
	if basicsecret == "" {
		basicsecret = "secret"
	}
	appEnv["BASICSECRET"] = fmt.Sprintf("%s", basicsecret)
	appEnv["GATEWAYSERVER"] = gwurl
	appEnv["SPARAMS"] = backendstr

	//if checkapp, err := s.GetAppByName(fmt.Sprintf("%s-monitoring", view.Name)); err == nil {
	if isMonitoring {
		//if checkapp.Count > 0 {
		if refresh == "refresh" {
			go func(appGuid string, setEnv map[string]interface{}) {
				time.Sleep(10 * time.Second)
				//body := map[string]string{ "environment_json" : domain.APP_STATE_STARTED }
				//data := domain.App {
				//	Environment: setEnv,
				//}
				//fmt.Println(setEnv)
				environment := map[string]interface{}{ "environment_json" : setEnv }
				_, err = s.UpdateApp(appGuid, environment, token)
				if err != nil {
					logger.Error("UpdateApp start err", err)
				}
				time.Sleep(10 * time.Second)
				_, err = s.RestageApp(appGuid, token)
				if err != nil {
					logger.Error("RestageApp start err", err)
				}
			}(monitoringAppGuid, appEnv)
			//}(checkapp.Resources[0].Meta.Guid)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"url":"`+ fmt.Sprintf("%s:%d/d/micrometer-%s/jvm-micrometer-%s", localip, s.uaa.GrafanaPort, view.Name, view.Name) +`"}`))

		return
		//}
	}

	fmt.Println(">>>>>>0")

	spaceName, err := s.userSpaceName(view.SpaceGuid)
	if err != nil {
		logger.Error("failed userSpaceName", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// prometheus app copy
	app, err := s.GetAppByName("prometheus-micro")
	if err != nil {
		logger.Error("prometheus GetAppByName err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sourceAppGuid := app.Resources[0].Meta.Guid
	summary := s.GetAppSummary(sourceAppGuid, token)
	data := domain.App {
		Name: fmt.Sprintf("%s-monitoring-%s", view.Name, spaceName),
		Instances: summary.Instances,
		Memory: summary.Memory,
		DiskQuota: summary.DiskQuota,
		State: domain.APP_STATE_STOPPED,
		SpaceGuid: view.SpaceGuid,
		Buildpack: summary.Buildpack,
		Command: "./prometheus_start.sh",
		//Command: "./prometheus_start.sh --web.listen-address=:8080",
	}

	data.Environment = appEnv
	createdApp, err := s.CreateApp(data, token)
	if err != nil {
		logger.Error("CreateApp err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createAppGuid := createdApp.Meta.Guid

	m2 := map[string]string{
		"source_app_guid" : sourceAppGuid,
	}
	_, err = s.CopyAppBits(createAppGuid, m2)
	if err != nil {
		logger.Error("CopyAppBits err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sharedDomains, err := s.ListSpaceDomains(view.SpaceGuid)
	if err != nil {
		logger.Error("ListSpaceDomains err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sharedDomain := sharedDomains.Resources[0]
	//CF Creating a Route
	routeData := domain.Route{
		Host: createdApp.Entity.Name,
		DomainGuid: sharedDomain.Meta.Guid,
		SpaceGuid: view.SpaceGuid,
	}
	route, err := s.CreateRoute(routeData)
	if err != nil {
		logger.Error("CreateRoute err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// CF Associate Route with the App
	_, err = s.AssociateRoute(createAppGuid, route.Meta.Guid)
	if err != nil {
		logger.Error("AssociateRoute err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// app start
	go func(appGuid string) {
		time.Sleep(10 * time.Second)
		body := map[string]string{ "state" : domain.APP_STATE_STARTED }
		_, err = s.UpdateApp(appGuid, body, token)
		if err != nil {
			logger.Error("UpdateApp start err", err)
		}
	}(createAppGuid)

	// db insert
	appRequest := domain.MicroserviceApp{MicroID: view.ID, AppGuid: createAppGuid, SourceGuid: sourceAppGuid, Essential: string(domain.MsMonitoring)}
	_, err = s.repositoryFactory.Compose().CreateMicroserviceApp(appRequest)
	if err != nil {
		logger.Error("CreateMicroserviceApp err", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
	}


	fmt.Println(">>>>>>1")

	// grafana datasource create

	jsonStr := []byte(`{
	  "name":"`+view.Name+`_datasource",
	  "version": 1,
	  "type":"prometheus",
	  "orgId": "1",
	  "url":"`+fmt.Sprintf("http://%s.%s", createdApp.Entity.Name, sharedDomain.Entity.Name)+`",
	  "access":"proxy",
	  "basicAuth":false,
	  "editable": false
	}`)

	if err = requestGrafanaServer("POST", s.uaa.GrafanaUrl,s.uaa.GrafanaAdminPassword,  "api/datasources", jsonStr, nil, logger); err != nil {
		logger.Error("http post grafana datasource request", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	fmt.Println(">>>>>>2")

	// grafana dashboard create
	/*
	jsonStr = []byte(`{
	  "dashboard": {
	    "id": null,
	    "uid": "micrometer-`+view.Name+`",
	    "title": "JVM (Micrometer) `+view.Name+`",
	    "tags": [ "templated" ],
	    "timezone": "browser",
	    "version": 0
	  },
	  "folderId": 0
	}`)
	if err = requestGrafanaServer("api/dashboards/db", jsonStr, logger); err != nil {
		logger.Error("http post grafana dashboard request", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/


	// grafana dashboard create
	type param struct{
		DsPrometheus string
		Title        string
		Uid          string
		State        string
		Action       string
		Level        string
		Cause        string
	}

	tmpl := param{
		DsPrometheus: view.Name+"_datasource",
		Title: "JVM (Micrometer) "+view.Name,
		Uid: "micrometer-"+view.Name,
		State: "{{state}}",
		Action: "{{action}}",
		Level: "{{level}}",
		Cause: "{{cause}}",
	}
	t, err := template.New("Grafana import").Parse(JVM_MICROMETER_JSON)
	if err != nil {
		logger.Error("template parse", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buff := bytes.NewBufferString("")
	err = t.Execute(buff, tmpl)
	if err != nil {
		logger.Error("template execute", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = requestGrafanaServer("POST", s.uaa.GrafanaUrl, s.uaa.GrafanaAdminPassword,  "api/dashboards/db", buff.Bytes(), nil, logger); err != nil {
		logger.Error("http post grafana dashboard import", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(">>>>>>3")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"url":"`+ fmt.Sprintf("%s:%d/d/micrometer-%s/jvm-micrometer-%s", localip, s.uaa.GrafanaPort, view.Name, view.Name) +`"}`))


}

func requestGrafanaServer(method, url, password, path string, data []byte, respData interface{}, logger lager.Logger) error {
	//localip, err := localip.LocalIP()
	//if err != nil {
	//	return err
	//}

	//localip := "10.244.228.5"

	var reader io.Reader
	if(data != nil) {
		reader = bytes.NewReader(data)
	}

	grafanaUrl := fmt.Sprintf("%s/%s", url, path)
	//req, err := http.NewRequest("POST", grafanaUrl, bytes.NewBuffer(data))
	req, err := http.NewRequest(method, grafanaUrl, reader)
	if err != nil {
		//logger.Error("http post grafana request", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	basicAuthToken := base64.StdEncoding.EncodeToString([]byte("admin" + ":" + password))
	req.Header.Set("Authorization", "Basic "+basicAuthToken)
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)
	if err != nil {
		//logger.Error("http post grafana response", err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer response.Body.Close()


	bytes, err := ioutil.ReadAll(response.Body)
	if err == nil {
		fmt.Println(string(bytes))
	}

	if response.StatusCode > 299 {
		logger.Error("http grafana response", errors.New("http post grafana response"), lager.Data{"status": response.StatusCode})
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return errors.New("http grafana response")
	}

	if respData != nil {
		err = json.Unmarshal(bytes, respData)
		if err != nil {
			return fmt.Errorf("json unmarshal: %s", err)
		}
	}

	return nil
}


func localIP() (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "1.2.3.4:1")
	if err != nil {
		return "", err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	host, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}

	return host, nil
}


