# 使用Docker部署go语言教程
使用git克隆项目：git clone https://github.com/qbanxiaoli/resource-pool
## 方式一
1. 使用Dockerfile文件构建出镜像  
 ```
 docker build -t resource-pool .
 ```    
2. 使用命令启动服务
 ```
 docker run -d -p 80:80 -e TZ=Asia/Shanghai --name=resource-pool --restart=always resource-pool
 ```  
## 方式二
1. 使用docker-compose编排容器
 ```
curl -L https://github.com/docker/compose/releases/download/1.26.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
 ```    
2. 构建镜像并启动服务
 ```
 docker-compose up -d
 ```  