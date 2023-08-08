# sdd-learn-course

システム開発部(System Development Department -SDD)で使用する技術等の、学習ロードマップを示します。
順番を守らなくても良いです。例えば、goの勉強しながら、飽きたらReactをやるとかでもいいです。全部こなせばそれでいいんじゃないかな✨。

あと、実際に手を動かして欲しいです。

## 学習手順

### 0. 番外編(作成途中)
1. Web APIの基礎を学ぼう！
   - [【HTTP入門】Webページがブラウザに表示されるまでの流れを解説](https://youtu.be/b_apIgHNqtk)
   - [【図解】RESTful APIとは何なのか【2023年版】](https://ramble.impl.co.jp/2886/)
2. フロントエンドとバックエンドの分離についてイメージを掴もう！
   - [【CORS入門】もうCORSエラーに苦しむことはありません。Webエンジニア必見です。](https://youtu.be/8fE2TmbPqlU)
      CORSというものの解説動画ですが、フロントエンドとバックエンドの分離についてイメージを掴むのには最適です。
3. Githubの使い方とGithub Flowについて

### 1. サーバーサイド

 1. [Goの環境構築をしよう！](./go-env/README.md)

    Goの環境構築をしてみましょう。方法は[Goの環境構築をしよう！](./go-env/README.md)のREADMEにまとめました。hello worldの出力と、簡単なユニットテストまでを実施します。

 2. [A Tour of Go](https://go-tour-jp.appspot.com/) 

    Go言語公式が出しているチュートリアルです。最後までやろう⭐️
    敷居が高いと感じたら以下のYoutube動画を見ることを勧めます。(二倍速でいいんじゃないかな)
    [【たった1時間で学べる】Go言語のプログラミング初心者向けの超入門講座](https://youtu.be/kPXfMFJ0oIE)
    - 課題1.
        - [Go言語の基礎](./go)
        - Go言語の基礎文法を確かめるための問題を用意しました。テストコードを用意したので、それを使って答え合わせをしてください。具体的な取り組み方は[Go言語の基礎](./go-basic)にあるREADMEを参照してください。

 3. [ginを最速でマスターしよう -Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)

    ginはGo（Golang）で書かれたWebフレームワークです。
    - 課題2. 
        - [写経 $+ \alpha$](./gin)
        - この記事を写経しながら実装してください。ただし、この記事ではデータベース操作にxormというライブラリを使っていますが、我々はgormを使います。gormに書き換えながら実装してください。具体的な取り組み方については[写経 $+ \alpha$](./gin)にあるREADMEを参照してください。
    - [公式ドキュメント](https://gin-gonic.com/ja/)
        ginの公式ドキュメントですが、あまり豊富ではないです。[APIの使い方のサンプル](https://gin-gonic.com/ja/docs/examples/)はある程度参考になると思います。
   
   4. Clean Architecture
   
   まずClean Architectureについて自身で調べてみて下さい。
   具体的な実装は以下の記事がわかりやすいです。疎結合サイコー！！！
      - [ChatGPTと対話しながら、クリーンアーキテクチャに基づくTodo APIを実装してみた](https://qiita.com/Sakaguchi_0725/items/bf613c805ba89afa39bf)
### 2. フロントエンド

2と3はドキュメントを読む作業なので、少し退屈かと思います。2と3を並行して取り組めば、多少マシになると思います。

1. [Next.jsで学ぶReact講座 -YouTube](https://www.youtube.com/watch?v=15WLMqnkPsE&list=PLwM1-TnN_NN6fUhOoZyU4iZiwhLyISopO)

   そうです、これはReactのみを学ぶ講座ではありません。Next.jsと一緒に学ぶことで余計な**つまずき**を減らすことができるというのです。一度Reactを学んだことがある人は倍速で見ることをお勧めします。動画の後半ではgithubにソースを反映させるのですが、コミットの粒度という観点から参考になるので、最後まで見ましょう！

   また、以下について説明できるようになると良いです。
   - useCallback,useEffect, useMemoの違いとユースケース

2. [【React+TypeScript】Netflixのクローンを作るチュートリアル -Zenn](https://zenn.dev/gunners6518/books/4c4672f32dd100)

   React+TypeScriptを使ったNetflixの映画一覧を表示するアプリケーションのチュートリアルです。

3. Shin codeさんのYoutubeで実践を学ぼう！

   - [【初心者OK】Next.jsで自作ブログを作ってみよう【MicroCMSを利用】](https://youtu.be/dNpONz4Yi04)
   - [Todoアプリを作りながらNext.js13の新機能を理解してみよう 〜Next.js13入門〜](https://youtu.be/VcMW2C9VNtI)
   - [React × TailwindCSSでポートフォリオサイトを構築してみよう](https://youtu.be/82cN8zwDhbY)
   - Udemy 【Discordクローン開発】React/Redux/Typescript/Firebaseで作るアプリ開発実践講座
4. [【Recoil入門】１からReactの状態管理ライブラリのRecoilを学んでみよう](https://youtu.be/S93hsNFmIcM)

   グローバルなstateを管理するライブラリです。例えば子コンポーネントのさらに子コンポーネントにstateを渡す時とか、propsだとめんどいよねってのを解消してくれます。

5. [React公式ドキュメント](https://ja.react.dev/)

   ただひらすら読みましょう。チュートリアルは手を動かしましょう。
   
   - 課題3. 
        - [Reactの基礎](./react-base)
           Hooksの説明やコードのリファクタリングがメインです。具体的な取り組み方については[Reactの基礎](./react-base)にあるREADMEを参照してください。
6. [サバイバルTypeScript](https://typescriptbook.jp/)

   大きくは[作って学ぶTypeScript](https://typescriptbook.jp/tutorials)と[読んで学ぶTypeScript](https://typescriptbook.jp/reference)に分かれています。前者は手を動かして欲しいですが、後者は読むだけで構いません。

### 番外編 ハンズオン

1. Next.jsとGoでシンプルなToDoアプリを作ってみよう！
### 3. AWS

### 4. Docker

### 5. セキュリティ
