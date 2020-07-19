package config

const Default = `
db:
  host: 127.0.0.1
  port: "5431"
  user: postgres
  dbname: monitor
  password: postgres
  sslmode: disable
redis:
  host: 127.0.0.1
  port: "6379"
  threshold: 4
jwt:
  secret: jdnfksdmfks
  exp: 15
`
