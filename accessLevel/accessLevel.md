# interfaceの役割1
### アクセスレベルを調整する
### Mockや依存性逆転でも活躍

一つ目のclientだと関係ないパスワードにアクセスできてしまう

client2はget postしかアクセスできない
globalなNewしかアクセスできず返り値がinterfaceなので
interfaceの中のget post にしかアクセスさせたくない意図がわかる