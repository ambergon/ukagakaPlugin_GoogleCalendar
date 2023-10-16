# ukagakaPlugin_GoogleCalendar
このプラグインはグーグルカレンダーの予定をSSPで確認できるように作成されました。<br>
以前作成した、SAORI-BASICでは、ゴーストから予定を教えてもらう事ができるという楽しさがありますが、通信を同期処理で行うため、SSP全体の処理が止まってしまうという問題がありました。<br>
また、関数へのアクセスをすべてゴーストから行う為、引数の受け渡しが非常に面倒で、設定の幅が狭かった問題がありました。<br>
<br>
より、一般的な(機能的でない)ゴーストを楽しむ為に、これらの問題を解消し、プラグイン化しました。<br>
<br>
前回作成したもの<br>
[GitHub - ambergon/ukagakaExe_GoogleCalendar](https://github.com/ambergon/ukagakaExe_GoogleCalendar)<br>
<br>

## 事前準備
#### 認証ファイルの作成
必要な認証ファイルの作成は同じものを使用します。<br>
こちらの作成のめんどくささは改善していません。<br>
[【伺か/SSP】ゴーストとグーグルカレンダーを連携させる設定備忘録 -- 異風堂々](https://ambergonslibrary.com/ukagaka/8900/)<br>


## 設定
Config.jsonを編集します。<br>
```
    "FilePath"        : "C:/Users/YourUserName/.../認証ファイル.json" ,
    "GmailAddress"    : "Calendarを共有したGmailAddress@gmail.com" ,
    "TimeZone"        : -9 , 
    "StartMin"        : 30 ,
    "From"            : 0 ,
    "Util"            : 7 ,
    "Sep"             : "\\n" 
```

- TimeZone<br>
    指定した時間だけタイムゾーンを変更します。

- StartMin<br>
    分単位で指定します。<br>
    指定した時間の経過後&カレンダーの取得が終わったタイミングで予定を表示します。<br>
    0-180まで指定可能です。0を指定した場合は、カレンダーからの情報を取得しません。<br>

- From    <br>
    カレンダーの予定を取得を開始する日。<br>

- Util    <br>
    カレンダーの予定を取得を終了する日。<br>
    どちらも0にした場合は起動時の当日の予定を取得します。<br>

- Sep     <br>
    取得した複数の予定をどのように区切ってゴーストに渡すか指定します。

デフォルトだと起動時から一週間の予定を改行で区切って表示します。<br>
予定が存在しなかった場合は特に何もしません。<br>

## 問題点
Golangを使用したdllなのでfreelibraryやSSPから有効->無効に切り替えなどをするとSSPが落ちます。


## Author
ambergon
