# chain_learn

## ブロックチェーンの構造解説

ブロックチェーンは``ブロック``と呼ばれるデータ集合を``チェーン``のように連なって構成されている．

各ブロックには複数の取引記録や前のブロックへの参照(ハッシュ値)が入っている．

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

equenceDiagram
    autonumber
    actor User as ユーザー (Tier 2)
    participant P2P as P2Pゴシップ通信
    participant Pool as 未承認プール (Mempool)
    participant Node as 検証ノード群 (Tier 1)
    participant Ledger as ブロックチェーン (台帳)

    User->>P2P: 1. トランザクション発行 (ZKP/Vouch)
    P2P->>Pool: 2. バケツリレー伝播
    Pool->>Node: 3. Txを集約
    Node->>Node: 4. BFT合意形成 (2/3以上の署名)
    Node->>Ledger: 5. 新ブロック接続
    Ledger-->>User: 6. 状態更新完了

