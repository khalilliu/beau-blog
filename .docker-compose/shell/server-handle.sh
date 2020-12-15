#! /bin/bash

rm -f ./config.yaml

touch ./config.yaml
filename="./config.yaml"
cat>"${filename}"<<EOF
# Beau-Blog Global Configuration

# jwt configuration
jwt:
  singing-key: 'bbPlus'

# zap logger configuration
zap:
  level: 'info'
  format: 'console'
  prefix: '[BEA_BLOG]'
  director: 'log'
  link-name: 'latest_log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  stacktrace-key: 'stacktrace'
  log-in-console: true

# redis configuration
redis:
  db: 0
  addr: '127.0.0.1:6379'
  password: ''

# email configuration
email:
  to: 'xxx@qq.com'
  port: 465
  from: 'xxx@163.com'
  host: 'smtp.163.com'
  is-ssl: true
  secret: 'xxx'
  nickname: 'test'

# system configuration
system:
  env: 'public'  # Change to "develop" to skip authentication for development mode
  addr: 8888
  db-type: 'mysql'
  oss-type: 'local'
  use-multipoint: false

# captcha configuration
captcha:
  key-long: 6
  img-width: 240
  img-height: 80


# mysql connect configuration
mysql:
  path: '127.0.0.1:3306'
  config: 'charset=utf8mb4&parseTime=True&loc=Local'
  db-name: 'qmPlus'
  username: 'root'
  password: 'Aa@6447985'
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: false
  log-zap: false


# local configuration
local:
  path: 'uploads/file'


# qiniu configuration (请自行七牛申请对应的 公钥 私钥 bucket域名地址)
qiniu:
  zone: 'ZoneHuadong'
  bucket: 'qm-plus-img'
  img-path: 'http://qmplusimg.henrongyi.top'
  use-https: false
  access-key: '25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ'
  secret-key: 'pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY'
  use-cdn-domains: false
EOF