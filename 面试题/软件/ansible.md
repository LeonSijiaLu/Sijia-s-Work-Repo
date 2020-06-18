# Ansible

1. Ansible 架构

Ansible 采用无代理架构，工作原理是使用SSH直接连接node, 并将Ansible模块和命令推送到目标执行的。

好处就是无代理，有效减少网络和运维开销。

2. 如果想在 Ansible 中保存秘密数据
    a. 使用 vault 来加密 ansibe playbook 中的一些变量
    b. 在编写 playbook 的时候，我们有的时候需要从别的 yml 文件引入变量
    c. 可以使用 vault 加密变量
    d. 提供密码才能使用那些变量

3. 怎么为 playbook 设置其他环境变量

4. 如何处理不同用户账户登陆的计算机

在 Ansible hosts 配置文件中进行配置。根据不同的主机要求可以进行分组，指定登录用户名和密码

5. Roles

是 ansible 一组文件的抽象集合。一个 role 一般都有自己的 handlers, variables. 加载role就加载了哪些文件。

6. Handler

当我们在执行 task 的时候。我们使用 notify 去使用 handler. 但是 handler 只有在所有 tasks 都执行结束之后才会执行。
