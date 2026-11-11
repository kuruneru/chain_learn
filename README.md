# ブロックチェーン

## ブロックチェーンの構造解説

ブロックチェーンは `ブロック` と呼ばれるデータ集合を `チェーン` のように連なって構成されている．

各ブロックには複数の取引記録や前のブロックへの参照(ハッシュ値)が入っている．

```mermaid
graph LR
    subgraph Block0 ["ジェネシスブロック No.0"]
        H0["自分のHash: 0000a1b2"]
        PH0["前Hash: 00000000"]
        TX0["データ: Genesis"]
    end

    subgraph Block1 ["ブロック No.1"]
        H1["自分のHash: 0000c3d4"]
        PH1["前Hash: 0000a1b2"]
        TX1["データ: ZKP登録申請"]
    end

    subgraph Block2 ["ブロック No.2"]
        H2["自分のHash: 0000e5f6"]
        PH2["前Hash: 0000c3d4"]
        TX2["データ: Vouchスコア譲渡"]
    end

    H0 --> PH1
    H1 --> PH2
