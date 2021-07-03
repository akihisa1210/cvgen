# cvgen

職務経歴書ジェネレーター

## Usage

### Example input file

`career.yml`

```yaml
companies:
  - name: A
    summary: 〇〇をやった
    projects:
      - period: 2021年05月〜2021年06月
        role: リーダー
        technology: Go, TypeScript
        activities:
          - 実装をした
          - テストをした
      - period: 2021年04月〜2021年05月
        role: リーダー
        technology: C++
        activities:
          - 実装をした
          - テストをした
  - name: B
    summary: xxをやった
    projects:
      - period: 2021年03月〜2021年04月
        role: メンバー
        technology: PHP, MySQL
        activities:
          - マネジメントをした
```

### Generate HTML

```sh
$ cvgen career.yml
<!DOCTYPE html>
        <html>
        <head>
                <style>
                        table {
                                border-collapse: collapse;
                        }

                        th, td {
                                border: solid 1px;
                                text-align: left;
                                vertical-align: top;
                        }
                </style>
        </head>
        <body>
                <table>
                        <thead>
                        </thead>
                        <tbody>

                                <tr>
                                        <td colspan="4">社名: A</td>
                                </tr>
                                <tr>
                                        <td colspan="4">〇〇をやった</td>
                                </tr>

                                <tr>
                                        <td rowspan="3" >2021年05月〜2021年06月</td>
                                        <td>役割: リーダー</td>
                                </tr>
                                <tr>
                                        <td>使用技術: Go, TypeScript</td>
                                </tr>
                                <tr>
                                        <td>
                                                <ul>

                                                        <li>実装をした</li>

                                                        <li>テストをした</li>

                                                </ul>
                                        </td>
                                </tr>

                                <tr>
                                        <td rowspan="3" >2021年04月〜2021年05月</td>
                                        <td>役割: リーダー</td>
                                </tr>
                                <tr>
                                        <td>使用技術: C++</td>
                                </tr>
                                <tr>
                                        <td>
                                                <ul>

                                                        <li>実装をした</li>

                                                        <li>テストをした</li>

                                                </ul>
                                        </td>
                                </tr>


                                <tr>
                                        <td colspan="4">社名: B</td>
                                </tr>
                                <tr>
                                        <td colspan="4">xxをやった</td>
                                </tr>

                                <tr>
                                        <td rowspan="3" >2021年03月〜2021年04月</td>
                                        <td>役割: メンバー</td>
                                </tr>
                                <tr>
                                        <td>使用技術: PHP, MySQL</td>
                                </tr>
                                <tr>
                                        <td>
                                                <ul>

                                                        <li>マネジメントをした</li>

                                                </ul>
                                        </td>
                                </tr>


                        </tbody>
                </table>
        </body>
        </html>
```

### Generate markdown

```sh
$ cvgen -m career.yml

# 職務経歴書

## A

〇〇をやった

### 2021年05月〜2021年06月

役割: リーダー

使用技術: Go, TypeScript

- 実装をした
- テストをした

### 2021年04月〜2021年05月

役割: リーダー

使用技術: C++

- 実装をした
- テストをした

## B

xxをやった

### 2021年03月〜2021年04月

役割: メンバー

使用技術: PHP, MySQL

- マネジメントをした
```
