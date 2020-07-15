# brave-gecko

Discord bot

# 運用ルール

## git

git-flow をベースとする。

開発時は最新の develop から作業ブランチを切り、プルリクエストを出す。
master ブランチは常に production の状態となる。
**master, develop で直にコミットを行わない**。

1. `$ git switch develop`
1. `$ git pull`
1. `$ git switch -c feat/new-featuer`
1. `$ git push -u origin feat/new-featuer`

### ブランチ名

| 区分                               | 命名                     |
| ---------------------------------- | ------------------------ |
| 新機能、機能変更                   | feat/wakari-yasui-namae  |
| リリース済みのバグ修正             | fix/wakari-yasui-namae   |
| 機能追加を伴わないドキュメント追加 | doc/wakari-yasui-namae   |
| 機能追加を伴わないテスト追加       | test/wakari-yasui-namae  |
| その他                             | chore/wakari-yasui-namae |

### コミットメッセージ

| 区分                               | 命名                        |
| ---------------------------------- | --------------------------- |
| 新機能、機能変更                   | feat: wakari yasui message  |
| リリース済みのバグ修正             | fix: wakari yasui message   |
| 機能追加を伴わないドキュメント追加 | doc: wakari yasui message   |
| 機能追加を伴わないテスト追加       | test: wakari yasui message  |
| その他                             | chore: wakari yasui message |

## デプロイ

master ブランチへ push があると自動で heroku にデプロイされる。
`--ff-only` で develop ブランチを master にマージし、push する。

1. `$ git switch develop`
1. `$ git pull`
1. `$ git switch master`
1. `$ git pull`
1. `$ git merge --ff-only develop`
1. `$ git push`
1. `$ git switch develop`
