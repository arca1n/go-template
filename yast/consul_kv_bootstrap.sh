curl \
    --request PUT \
    --data '{"db_host":"postgres","db_user":"postgres","db_password":"example","db_name":"auth","db_usessl":"false","db_port":5432}' \
    http://127.0.0.1:8500/v1/kv/DB_CONFIG
curl \
    --request PUT \
    --data '{"redis_host":"redis:6379","user":"root","password":""}' \
    http://127.0.0.1:8500/v1/kv/REDIS_CONFIG
curl \
    --request PUT \
    --data '{"nats_host":"nats"}' \
    http://127.0.0.1:8500/v1/kv/NATS_CONFIG
curl \
    --request PUT \
    --data '{"type":"dev","log_level":"debug"}' \
    http://127.0.0.1:8500/v1/kv/ENVIRONMENT
