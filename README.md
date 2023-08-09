# sdd-learn-course

システム開発部(System Development Department -SDD)で使用する技術等の、学習ロードマップを示します。
順番を守らなくても良いです。例えば、goの勉強しながら、飽きたらReactをやるとかでもいいです。全部こなせばそれでいいんじゃないかな✨。

あと、実際に手を動かして欲しいです。  

Udemyは見るタイミングで連絡ください。

## 学習手順

### 0. 番外編(作成途中)
1. Web APIの基礎を学ぼう！  
   - [【HTTP入門】Webページがブラウザに表示されるまでの流れを解説](https://youtu.be/b_apIgHNqtk)
   - [【図解】RESTful APIとは何なのか【2023年版】](https://ramble.impl.co.jp/2886/)
2. フロントエンドとバックエンドの分離についてイメージを掴もう！  
   - [【CORS入門】もうCORSエラーに苦しむことはありません。Webエンジニア必見です。](https://youtu.be/8fE2TmbPqlU)
      CORSというものの解説動画ですが、フロントエンドとバックエンドの分離についてイメージを掴むのには最適です。
3. Githubの使い方とGithub Flowについて  
   - [Git研修【MIXI 23新卒技術研修】](https://speakerdeck.com/mixi_engineers/2023-git-training)　　
      - [動画](https://www.youtube.com/watch?v=lWkO8bQ9pSo)もあります。
   - [GitHub Flowとは](https://qiita.com/tatane616/items/aec00cdc1b659761cf88)

### 1. サーバーサイド

 1. [Goの環境構築をしよう！](./go-env/README.md)  

    Goの環境構築をしてみましょう。方法は[Goの環境構築をしよう！](./go-env/README.md)のREADMEにまとめました。hello worldの出力と、簡単なユニットテストまでを実施します。

 2. [A Tour of Go](https://go-tour-jp.appspot.com/)    

    Go言語公式が出しているチュートリアルです。最後までやろう⭐️
    敷居が高いと感じたら以下のYoutube動画を見ることを勧めます。(二倍速でいいんじゃないかな)  

    [Gopher道場](https://www.youtube.com/watch?v=YrGUL6atiJ4&list=PL9MOSAifWs3whvWOsObk3uCBXtVAhD2A7)
    - 課題1.
        - [Go言語の基礎](./go)
        - Go言語の基礎文法を確かめるための問題を用意しました。テストコードを用意したので、それを使って答え合わせをしてください。具体的な取り組み方は[Go言語の基礎](./go-basic)にあるREADMEを参照してください。

 3. [ginを最速でマスターしよう -Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)  

    ginはGo（Golang）で書かれたWebフレームワークです。
    - 課題2. 
        - [写経 + alpha](./gin)
        - この記事を写経しながら実装してください。ただし、この記事ではデータベース操作にxormというライブラリを使っていますが、我々はgormを使います。gormに書き換えながら実装してください。
    - [公式ドキュメント](https://gin-gonic.com/ja/)
        ginの公式ドキュメントですが、あまり豊富ではないです。[APIの使い方のサンプル](https://gin-gonic.com/ja/docs/examples/)はある程度参考になると思います。
   
1. Clean Architecture  
   
   まずClean Architectureについて自身で調べてみて下さい。
   具体的な実装は以下の記事がわかりやすいです。疎結合サイコー！！！
      - [ChatGPTと対話しながら、クリーンアーキテクチャに基づくTodo APIを実装してみた](https://qiita.com/Sakaguchi_0725/items/bf613c805ba89afa39bf)

2. フロントエンドを勉強してみよう！
   1. [Next.jsで学ぶReact講座 -YouTube](https://www.youtube.com/watch?v=15WLMqnkPsE&list=PLwM1-TnN_NN6fUhOoZyU4iZiwhLyISopO) 
   2. [Todoアプリを作りながらNext.js13の新機能を理解してみよう 〜Next.js13入門〜](https://youtu.be/VcMW2C9VNtI)
      
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
  
4. Udemy Nextjs + Tailwind CSS +  Django REST Framework で学ぶモダンReact開発

4. [【Recoil入門】１からReactの状態管理ライブラリのRecoilを学んでみよう](https://youtu.be/S93hsNFmIcM)  

   グローバルなstateを管理するライブラリです。例えば子コンポーネントのさらに子コンポーネントにstateを渡す時とか、propsだとめんどいよねってのを解消してくれます。

5. [React公式ドキュメント](https://ja.react.dev/)  

   ただひらすら読みましょう。チュートリアルは手を動かしましょう。
   
6. [サバイバルTypeScript](https://typescriptbook.jp/)  

   大きくは[作って学ぶTypeScript](https://typescriptbook.jp/tutorials)と[読んで学ぶTypeScript](https://typescriptbook.jp/reference)に分かれています。前者は手を動かして欲しいですが、後者は読むだけで構いません。

7. Udemy AWSで作るWEBアプリケーション 実践講座  

   見よう。

### 番外編 ハンズオン

#### フロントエンド担当  

Next.jsでmarkdownブログを構築してみよう！

#### バックエンド担当  

Next.jsとGoでシンプルなToDoアプリを作ってみよう！
### 3. AWS

1. Udemy AWSで作るWEBアプリケーション 実践講座
2. Udemy AWSで学ぶ！マイクロサービス入門
3. Udemy 【ベストセラー完全日本語化】AWS 認定ソリューションアーキテクト アソシエイト SAA-C03 対応 2022 最新版  
これは流し見でいい。

### 4. Docker

1. [初心者が絵で理解する Docker](https://zenn.dev/suzuki_hoge/books/2021-04-docker-picture-60fbe950136be9c7ad85)

2. [実践 Docker - ソフトウェアエンジニアの「Docker よくわからない」を終わりにする本](https://zenn.dev/suzuki_hoge/books/2022-03-docker-practice-8ae36c33424b59)

### 5. Terraform

Udemy(まだ買ってない)