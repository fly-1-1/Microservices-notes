# 一 dockerfile 初识

## 1 dockerfile入门案例

​	使用Dockerfile来基于ubuntu创建一个定制化的镜像：nginx。 

```
#创建Dockerfile专用目录
:~$ mkdir ./docker/images/nginx -p
:~$ cd docker/images/nginx/
#创建Dockerfile文件
:~/docker/images/nginx$ vim Dockerfile
```

​	dockerfile内容

```
# 构建一个基于ubuntu的docker定制镜像
# 基础镜像
FROM ubuntu
# 镜像作者
MAINTAINER panda kstwoak47@163.com
# 执行命令
RUN mkdir hello
RUN mkdir world
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update
RUN apt-get install nginx -y
# 对外端口
EXPOSE 80
```

```
docker build -t ubuntu-nginx:v1.0 .
```

# 二 dockerfile 基础

## 1 dockerfile基础指令

### 1 FROM

```
FROM
#格式：
FROM <image>
FROM <image>:<tag>
#解释：
#FROM 是 Dockerfile 里的第一条而且只能是除了首行注释之外的第一条指令
#可以有多个FROM语句，来创建多个image
#FROM 后面是有效的镜像名称，如果该镜像没有在你的本地仓库，那么就会从远程仓库Pull取，如果远程也
没有，就报错失败
#下面所有的 系统可执行指令 在 FROM 的镜像中执行。
```

### 2 MAINTAINER

```
MAINTAINER
#格式：
MAINTAINER <name>
#解释：
#指定该dockerfile文件的维护者信息。类似我们在docker commit 时候使用-a参数指定的信息
```

### 3 RUN

```
RUN
#格式：
RUN <command> (shell模式)
RUN["executable", "param1", "param2"] (exec 模式)
#解释：
#表示当前镜像构建时候运行的命令，如果有确认输入的话，一定要在命令中添加 -y
#如果命令较长，那么可以在命令结尾使用 \ 来换行
#生产中，推荐使用上面数组的格式
#注释：
#shell模式：类似于 /bin/bash -c command
#举例： RUN echo hello
#exec模式：类似于 RUN["/bin/bash", "-c", "command"]
#举例： RUN["echo", "hello"]
```

### 4 EXPOSE

```
EXPOSE
#格式：
EXPOSE <port> [<port>...]
#解释：
设置Docker容器对外暴露的端口号，Docker为了安全，不会自动对外打开端口，如果需要外部提供访问，
还需要启动容器时增加-p或者-P参数对容器的端口进行分配。
```

## 2 docker 运行时指令

### 1 CMD

```
CMD
#格式：
CMD ["executable","param1","param2"] (exec 模式)推荐
CMD command param1 param2 (shell模式)
CMD ["param1","param2"] 提供给ENTRYPOINT的默认参数；
#解释：
#CMD指定容器启动时默认执行的命令
#每个Dockerfile只能有一条CMD命令，如果指定了多条，只有最后一条会被执行
#如果你在启动容器的时候使用docker run 指定的运行命令，那么会覆盖CMD命令。
#举例： CMD ["/usr/sbin/nginx","-g","daemon off；"]
"/usr/sbin/nginx" nginx命令
"-g" 设置配置文件外的全局指令
"daemon off；" 后台守护程序开启方式 关闭


#CMD指令实践:
#修改Dockerfile文件内容：
#在上一个Dockerfile文件内容基础上，末尾增加下面一句话：
CMD ["/usr/sbin/nginx","-g","daemon off;"]
#构建镜像
:~/docker/images/nginx$ docker build -t ubuntu-nginx:v0.3 .
#根据镜像创建容器,创建时候，不添加执行命令
:~/docker/images/nginx$ docker run --name nginx-1 -itd ubuntu-nginx:v0.3
#根据镜像创建容器,创建时候，添加执行命令/bin/bash
:~/docker/images/nginx$ docker run --name nginx-2 -itd ubuntu-nginx:v0.3
/bin/bash
docker ps
#发现两个容器的命令行是不一样的
docker ps -a
CONTAINER ID IMAGE COMMAND CREATED NAMES
921d00c3689f ubuntu-nginx:v0.3 "/bin/bash" 5 seconds ago nginx-
2
e6c39be8e696 ubuntu-nginx:v0.3 "/usr/sbin/nginx -g …" 14 seconds ago nginx-
1
```

### 2 ENTRYPOINT

```
ENTRYPOINT
#格式：
ENTRYPOINT ["executable", "param1","param2"] (exec 模式)
ENTRYPOINT command param1 param2 (shell 模式)
#解释：
#和CMD 类似都是配置容器启动后执行的命令，并且不会被docker run 提供的参数覆盖。
#每个Dockerfile 中只能有一个ENTRYPOINT，当指定多个时，只有最后一个起效。
#生产中我们可以同时使用ENTRYPOINT 和CMD，
#想要在docker run 时被覆盖，可以使用"docker run --entrypoint"
#ENTRYPOINT指令实践：
#修改Dockerfile文件内容：
#在上一个Dockerfile 文件内容基础上，修改末尾的CMD 为ENTRYPOINT：
ENTRYPOINT ["/usr/sbin/nginx","-g","daemon off;"]
#构建镜像
:~/docker/images/nginx$ docker build -t ubuntu-nginx:v0.4 .
#根据镜像创建容器,创建时候，不添加执行命令
:~/docker/images/nginx$ docker run --name nginx-3 -itd ubuntu-nginx:v0.4
#根据镜像创建容器,创建时候，添加执行命令/bin/bash
:~/docker/images/nginx$ docker run --name nginx-4 -itd ubuntu-nginx:v0.4
/bin/bash
#查看ENTRYPOINT是否被覆盖
CONTAINER ID IMAGE COMMAND CREATED
NAMES
e7a2f0d0924e ubuntu-nginx:v0.4 "/usr/sbin/nginx -g …" 59 seconds ago
nginx-4
c92b2505e28e ubuntu-nginx:v0.4 "/usr/sbin/nginx -g …" About a minute ago
nginx-3
#根据镜像创建容器,创建时候，使用--entrypoint参数，添加执行命令/bin/bash
docker run --entrypoint "/bin/bash" --name nginx-5 -itd ubuntu-nginx:v0.4
#查看ENTRYPOINT是否被覆盖
:~/docker/images/nginx$ docker ps
CONTAINER ID IMAGE COMMAND CREATED
NAMES
6c54726b2d96 ubuntu-nginx:v0.4 "/bin/bash" 3 seconds ago
nginx-5
```

## 3 docker 文件编辑指令

### 1 add

​	在Docker中，`ADD`指令用于将本地文件、目录或远程文件 URL 添加到容器的文件系统中。它有一些用法和选项，让我们来看一下：

```
ADD <src> <dest>
```

- `<src>` 是构建上下文中的文件或目录的路径，可以是相对路径或绝对路径。
- `<dest>` 是容器中的目标路径，表示将文件或目录复制到容器中的位置。

#### 复制本地文件到容器中：

```
ADD myapp.jar /app/
```

​	这将把当前构建上下文中的 `myapp.jar` 复制到容器内的 `/app/` 目录下。

#### 复制本地目录到容器中：

```
ADD myapp /app/
```

​	这将把当前构建上下文中的 myapp 目录复制到容器内的 /app/ 目录下。

#### 使用URL复制文件到容器中：

```
ADD http://example.com/myfile.tar.gz /data/
```

​	这将从指定的 URL 下载文件并将其复制到容器内的 `/data/` 目录下。

注意：

1. 如果 `<src>` 是一个 URL，Docker 将会尝试自动解压缩压缩文件，例如 `.zip` 或 `.tar.gz`。
2. 如果 `<src>` 是一个目录，它的内容将被复制到 `<dest>` 中，而不是整个目录。

尽管`ADD`是一个功能强大的指令，但在许多情况下，建议使用更为灵活的`COPY`指令来完成文件复制的任务，因为`COPY`通常更直观且具有更明确的行为。

### 2 copy

`		COPY`指令用于将文件或目录从构建上下文复制到容器的文件系统中。与`ADD`相比，`COPY`更简单，更直接，通常是更好的选择，特别是在不需要`ADD`的高级功能时。

```
COPY <src> <dest>
```

- `<src>` 是构建上下文中的文件或目录的路径，可以是相对路径或绝对路径。
- `<dest>` 是容器中的目标路径，表示将文件或目录复制到容器中的位置。

#### 复制本地文件到容器中：

```
COPY myapp.jar /app/
```

​	这将把当前构建上下文中的 myapp.jar 复制到容器内的 /app/ 目录下。

#### 复制本地目录到容器中：

```
COPY myapp /app/
```

这将把当前构建上下文中的 myapp 目录复制到容器内的 /app/ 目录下。

注意:

1. 如果 `<src>` 是一个目录，它的内容将被复制到 `<dest>` 中，而不是整个目录。
2. 与`ADD`不同，`COPY`不具备`ADD`的自动解压缩功能。如果要复制压缩文件并解压缩，需要在`COPY`之后使用其他命令进行解压。

通常情况下，如果只是简单地将文件从构建上下文复制到容器中，建议使用`COPY`。 `COPY`的行为更为明确，不会涉及到`ADD`的一些高级功能，因此更容易理解和维护。

### 3 volums

​	在Dockerfile中，`VOLUME`指令用于声明容器中某个目录是持久化存储数据的目录，该目录可以在容器之间共享数据，并且这个数据不会随着容器的销毁而被删除。这允许容器之间或容器和宿主机之间共享数据，是持久存储的一种常见方式。

```
VOLUME ["<path>", ...]
<path> 是容器内的路径，用于指定一个或多个需要持久化的目录。
```

```
VOLUME ["/data"]
```

这将在容器内创建一个卷（volume），使 `/data` 目录成为一个持久化的存储目录。当容器启动时，Docker会自动创建这个卷，并且可以在其他容器中使用相同的卷来共享数据。

注意:

1. 使用`VOLUME`指令声明的卷是在容器启动时创建的，如果在`Dockerfile`中指定了初始数据，它将被复制到卷中。
2. `VOLUME`指令不会在本地主机上创建一个实际的目录，而是创建一个 Docker 管理的卷。
3. 当容器销毁时，由`VOLUME`创建的卷中的数据将保留，这使得容器可以在启动时访问之前存储的数据。
4. 可以使用`docker run`命令的`-v`或`--volume`选项将本地主机上的目录挂载到容器内的卷上。

```
docker run -v /path/on/host:/data myimage
```

这样做的效果是将本地主机上的 `/path/on/host` 目录挂载到容器内的 `/data` 目录上，实现数据共享。

`VOLUME`指令是用于声明在容器中创建卷，从而实现数据持久化和容器之间数据共享的目的。

### 4 env

​	在Dockerfile中，`ENV`指令用于设置环境变量。这些环境变量可以在后续的指令中使用，也可以在容器运行时通过`docker run`命令行选项进行覆盖。

```
ENV <key> <value>
```

- `<key>` 是环境变量的名称。

- `<value>` 是环境变量的值。


```
ENV DATABASE_URL="mysql://user:password@localhost:3306/mydatabase"
ENV APP_HOME="/usr/src/app"

```

在这个例子中，定义了两个环境变量，一个用于存储数据库连接字符串，另一个用于存储应用程序的根目录路径。

#### 多个键值对的示例：

```
ENV \
    DATABASE_URL="mysql://user:password@localhost:3306/mydatabase" \
    APP_HOME="/usr/src/app"
```

这是一种更为清晰的写法，适用于设置多个环境变量。

注意：`ENV`指令在构建时设置环境变量，而在运行时，可以通过`docker run`的`-e`选项来覆盖这些值。

```
docker run -e DATABASE_URL="new_value" myimage
```

​	ENV`指令设置的环境变量对于构建中的后续指令都是可见的，包括在`CMD`、`RUN`等指令中使用。

​	为了提高可读性，可以使用反斜杠 `\` 来换行，以便在Dockerfile中更清晰地排列多个`ENV`指令。

环境变量在容器中是一种重要的配置机制，它允许将配置信息与镜像的构建和运行过程分离，提高了灵活性和可维护性。

### 5 workdir

​	在Dockerfile中，`WORKDIR`指令用于设置容器中的工作目录。工作目录是容器中执行命令时的默认目录，可以简化命令的书写，并使 Dockerfile 更易读和维护。

```
WORKDIR /path/to/directory
/path/to/directory 是容器中的工作目录的路径。
```

#### 示例：

```
WORKDIR /usr/src/app
```

这将设置容器中的工作目录为 `/usr/src/app`。在这之后的指令中，如果使用相对路径，它们将相对于这个工作目录。

#### 多个`WORKDIR`指令：

​	可以在Dockerfile中多次使用`WORKDIR`指令，每次设置的值都会影响接下来的指令，直到遇到下一个`WORKDIR`指令。

```
WORKDIR /usr/src/app
RUN npm install

WORKDIR /usr/src/app/subdirectory
RUN npm install
```

​	在这个例子中，第一个`WORKDIR`将工作目录设置为 `/usr/src/app`，第二个`WORKDIR`将工作目录更改为 `/usr/src/app/subdirectory`。接下来的`RUN`指令将在这个新的工作目录中执行。

注意：

1. 如果提供的路径不存在，`WORKDIR`指令将会创建相应的目录。
2. 使用相对路径时，它们将相对于前一个`WORKDIR`指令设置的目录。
3. 可以使用绝对路径或相对路径来设置工作目录。

`WORKDIR`指令对于使Dockerfile更易读，更易维护，以及在容器中执行命令时更方便，是一个有用的工具。

### 6 onbuild

`ONBUILD`是Dockerfile中的一个特殊指令，它用于定义构建过程中的触发操作。当一个镜像被用作另一个镜像的基础镜像时，`ONBUILD`指令定义的操作将会在下游镜像的构建过程中被触发。

#### 基本语法：

```
ONBUILD <INSTRUCTION>
```

其中 `<INSTRUCTION>` 可以是任何合法的Dockerfile指令，例如`RUN`、`COPY`、`CMD`等。

#### 示例：

​	考虑一个基础镜像，其Dockerfile包含了`ONBUILD`指令：

```
# Base image Dockerfile
ONBUILD COPY . /app/src
ONBUILD RUN make /app/src
ONBUILD CMD ["python", "/app/src/app.py"]
```

这个基础镜像的`ONBUILD`指令表示，在该镜像被用作其他镜像的基础镜像时，执行以下步骤：

1. 将构建上下文中的所有文件复制到 `/app/src` 目录下。
2. 在 `/app/src` 目录中执行 `make` 命令。
3. 设置默认的启动命令为运行 `/app/src/app.py`。

现在，如果有一个新的Dockerfile使用上述基础镜像：

```
# New Dockerfile using the ONBUILD instructions
FROM baseimage:latest
```

​	在构建这个新的镜像时，上述`ONBUILD`指令中定义的步骤将会在构建过程中被触发。这样，你可以将一些通用的构建逻辑和步骤放在基础镜像中，然后在派生镜像中定义特定的应用逻辑。

需要注意的是，`ONBUILD`指令在Docker 17.05及更高版本中被废弃，因为它引入了一些复杂性和不确定性，导致构建过程的难以理解。因此，**建议尽量避免使用**`ONBUILD`指令，而是考虑使用更明确和可控的方式组织镜像构建过程。

# 4 docker compose

​	Docker Compose 是一个用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，你可以使用一个简单的 YAML 文件来配置应用的服务、网络和卷等，然后使用一个命令就可以启动整个应用。它简化了多容器应用的管理和部署。

Docker Compose 使用一个名为 `docker-compose.yml` 的文件来定义整个应用的配置。这个文件包含了服务、网络、卷等的定义。一个简单的 `docker-compose.yml` 文件的示例：

```
version: '3'
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
  database:
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: example
```

```
在 docker-compose.yml 文件中，你可以定义多个服务，每个服务对应一个容器。每个服务可以使用一个特定的镜像，配置容器的参数、环境变量、端口映射等。
```

### 服务定义

在 `docker-compose.yml` 文件中，你可以定义多个服务，每个服务对应一个容器。每个服务可以使用一个特定的镜像，配置容器的参数、环境变量、端口映射等。

### 启动应用

要启动应用，只需在包含 `docker-compose.yml` 文件的目录中运行以下命令：

```
docker-compose up
```

这将按照配置启动所有服务。你可以通过添加 `-d` 选项来在后台运行。

### 停止应用

要停止应用，可以运行：

```
docker-compose down
```

这将停止和删除所有服务和相关的网络、卷等。

### 多环境支持

`docker-compose.yml` 文件支持变量和环境覆盖，使得可以在不同的环境中使用相同的配置文件，并通过环境变量进行微调。

### 扩展和依赖

通过 Docker Compose，你可以轻松地定义多容器应用中的服务之间的依赖关系，并控制容器的启动顺序。这对于具有多个服务、例如 Web 服务器和数据库的应用程序非常有用。

Docker Compose 提供了许多其他命令和选项，例如构建镜像、查看服务日志、查看运行中的容器等。通过 `docker-compose --help` 命令可以查看所有可用的选项。