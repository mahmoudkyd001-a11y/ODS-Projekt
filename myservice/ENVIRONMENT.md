MYSERVICE_DEBUG   enable debug level for logging
MYSERVICE_TRACING   enable tracing level for logging
MYSERVICE_NAME   set the name of the instance of the service
MYSERVICE_TITLE   set the title in the web page
MYSERVICE_PORT_NB   the local port of the web service (default=8080)
MYSERVICE_API_KEYS   space separated list of valid API keys
MYSERVICE_SESSION_KEY
MYSERVICE_POLICY   OPA policy for access control
MYSERVICE_OPASVC   OPA service port to get the OPA policy for access control
MYSERVICE_REALM   Basic authentication realm
MYSERVICE_STAFF_USER   username of the administrator
MYSERVICE_STAFF_PASSWORD   password of the administrator
MYSERVICE_PARTICIPANT_USER   username of the user
MYSERVICE_PARTICIPANT_PASSWORD   password of the user
MYSERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
MYSERVICE_KEY_PEM   key for TLS (HTTPS) communication
MYSERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
MYSERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
MYSERVICE_LOKI_USER   user name as defined for the data source as basic authentication
MYSERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
MYSERVICE_LOKI_KEY   key/token as defined for the data source
MYSERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
MYSERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
MYSERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
MYSERVICE_LANGUAGE
MYSERVICE_LANGUAGES
MYSERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
MYSERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
MYSERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
MYSERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
