# Go言語の環境構築方法について

## Windows

PowerShellで以下のコマンドを打ちます。
INSTALLDIRは好きなように変更して下さい。

```sh
winget install GoLang.Go --override "INSTALLDIR=C:\\lang\\go"
```

## Mac

brewで一発っすね。

## 環境変数の設定

| 環境変数 | 設定値                        | 説明                                      |
| -------- | ----------------------------- | ----------------------------------------- |
| GOPATH   | C:Users\ユーザー名\好きなとこ | 作業ディレクトリを指定するパス            |
| PATH     | C:\lang\go\bin;  %GOPATH%\bin | C:\lang\go\binと%GOPATH%bin  をPathに追加 |

## 動作確認

CMDかPowerShellで次のコマンドを打って下さい。いい感じにバージョンが返って来たら成功。

```sh
go version
```

## Hello World

1. GitHubに適当なリポジトリを作成して下さい。

2. GOPATH以下に

```
src/github.com/ユーザー名/
```

というディレクトリを作って下さい。

3. 1.で作ったリポジトリを2.で用意したディレクトリにgit cloneし、クローンしたディレクトリにcdして下さい。例えば、go-start-projectという名前のリポジトリなら、パスは次のようになります。

```
src/github.com/syuya2036/go-start-project
```

4. 次のようなコマンドを打ってgoのモジュールを初期化します。
go.modというファイルができて、ここに依存パッケージが追記されていきます。

```
go mod init
```

5. main.goファイルを作成し、以下のようにコードを書きます。

```go
package main

import (
    "fmt
)

func main() {
    fmt.Println("Hello World!")
}
```

6. ビルドして実行します。

```
go run main.go
```

## Unit Test

1. greetingディレクトリを作成し、移動します。

```
mkdir greeting
cd greeting
```

2. greeting.goというファイルを作成し、以下のコードを書きます。

```go
package greeting

import (
    "fmt"
)

// nameに挨拶をするメッセージを返す。
func GetGreetMsg(name string) (string, error) {
    // nameが空ならエラー
    if name == "" {
        return "", fmt.Error("name is empty.")
    }

    msg := fmt.Sprintf("Hello, %s!", name)

    return msg, nil
}
```

3. greeting_test.goというファイルを作成し、以下のコードを書きます。

```go
package greeting

import (
	"testing"
)

func TestGetGreetMsg(t *testing.T) {
    testCases := []struct {
        name     string
        expected string
    }{
        {"shuya", "Hello, shuya!"},
        {"goto", "Hello, goto!"},
    }

    for _, tc := range testCases {
        result, err := GetGreetMsg(tc.name)
        if err != nil {
            t.Errorf("GetGreetMsg(%s) = %s; expected %s", tc.name, result, tc.expected)
        }
        
        if result != tc.expected {
            t.Errorf("GetGreetMsg(%s) = %s; expected %s", tc.name, result, tc.expected)
        }
    }

    //  空文字を入れるとエラーになることを確認
    result, err := GetGreetMsg("")

    if err == nil {
        t.Errorf("GetGreetMsg(\"\") = %s; expected error", result)
    }
}
```

4. テストを実行します。

```
go test
```