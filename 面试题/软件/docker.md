# Docker

1. 什么是 Docker

Docker 使用容器技术来把应用和依赖包打包

2. Docker Image

Docker 镜像是由一系列只读层构建的。DockerFile中的每一条指令或者我们在Docker容器中的每一步操作都会形成一层新的镜像

3. Docker 容器状态
    a. 运行
    b. 暂停
    c. 重新启动
    d. 退出

4. DockerFile
    a. From 制定镜像文件
    b. Label 指定标签
    c. RUN 运行什么语句
    d. CMD 容器启动时的运行语句

5. 什么是虚拟化
    a. 虚拟化就是在同一块硬件上运行多个操作系统。操作系统之间互相隔离

6. Docker vs 虚拟机
    a. 虚拟机使用 hypervisor，是操作系统内核级别上的隔离，安全性高。但是创建时间慢，而且资源消耗大。Docker 可以快速创建，类似进程
    b. Docker 使用主机调度资源，类似进程

7. 容器内存机制
    a. 每个容器都有属于自己的命名空间，容器看不到其他命名空间的资源，也就无法使用。但毕竟命名空间的实现依靠内核的调度，存在安全问题

8. 怎么监控 Docker
    a. 

9. Docker ADD vs COPY
    a. Docker Copy 可以把本地文件复制到 Docker 中
    b. Docker Add 可以把URL文件下载到 Docker 中