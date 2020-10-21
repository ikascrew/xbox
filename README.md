このリポジトリは、USBのジョイコンを操作する為に作成したものです。
USB IDを利用して、Eventをチャンネルで受け取ります。

# 生成方法

```
ctrl,err := controller.New(0)
```

使用するコントローラーのボタンと、ジョイスティックに名称をつけます
コントローラーにより、ボタンの配置は違う為、任意になります

# ボタン名の付与

```
ctrl.ButtonNames("A","B","X","Y",,,,,)
ctrl.AxisNames("LEFT_JOY","RIGHT_JOY",,,)
```

# 使い方

ボタン名をつけるとイベントが実行できます

```
ch := c.Event()
for {
    select {
    case ev := <-ch:
        if ev.Error() != nil {
            fmt.Printf("%+v\n", ev.Error())
            c.Terminate()
        } else {
            fmt.Print(ev)
        }
    default:
    }

    if c.Closed() {
        break
    }
}
```

ボタンを押すとイベントが起こります。
イベント内には押しているボタンの情報とエラーがあります。

終了する場合はTerminate()を行います。



## オプション

現状あるオプションは以下です。

- controller.Duration(50)  -> 内部のループTickerのミリ秒
- controller.RapidFire(false) -> ボタンを押し続けた場合のイベント有無
- controller.AxisMargin(2000)  -> ジョイスティックの０値のマージン 
- controller.Logger(log.New()) -> 内部ログの設定

controller.New()に引数で渡せます
