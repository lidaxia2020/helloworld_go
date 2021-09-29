# redis发布订阅
- 应用场景
  - 1、异步消息通知
  - 2、任务通知
  - 3、参数刷新加载
- 缺陷
  - **redis系统的稳定性有关。**对于旧版的redis来说，如果一个客户端订阅了某个或者某些频道，
但是它读取消息的速度不够快，那么不断的积压的消息就会使得redis输出缓冲区的体积越来越大，这可能
会导致redis的速度变慢，甚至直接崩溃。也可能会导致redis被操作系统强制杀死，甚至导致操作系统本
身不可用。新版的redis不会出现这种问题，因为它会自动断开不符
合client-output-buffer-limit pubsub配置选项要求的订阅客户端。
  - 是和数据传输的可靠性有关。任何网络系统在执行操作时都可能会遇到断网的情况。而断线产生的连接
错误通常会使得网络连接两端中的一端进行重新连接。如果客户端在执行订阅操作的过程中断线，那么客户
端将会丢失在断线期间的消息，这在很多业务场景下是不可忍受的。Redis 并不会对发布的消息持久化的