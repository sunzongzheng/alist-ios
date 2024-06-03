# AList-iOS
这是一个可顺利构建出iOS framework的AList分支仓库

# 构建步骤
1. 拉取web代码
   
   ```bash
   sh ./fetch-web.sh
   ```
2. 使用`gomobile`进行构建

   ```bash
   gomobile bind -target ios -bundleid {你的bundleid} -o {alist-expo仓库目录}/ios/alist/Alistlib.xcframework -ldflags "-s -w" github.com/alist-org/alist/v3/alistlib
   ```


# 仓库代码说明
除了从上游拉取最新代码以外，有以下变更：
1. 为了移动端使用方便，密码改为了明文存储
2. 从[jing332/AListFlutter](https://github.com/jing332/AListFlutter)拷贝`alistlib`代码用于启动AList服务、输出日志等，并做少量改造
3. 修改了若干不支持ios构建的子依赖代码，详见当前仓库的git submodules
