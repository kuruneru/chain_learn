# chain_learn

## ブロックチェーンの構造解説

ブロックチェーンは``ブロック``と呼ばれるデータ集合を``チェーン``のように連なって構成されている．

各ブロックには複数の取引記録や前のブロックへの参照(ハッシュ値)が入っている．

flowchart LR
    subgraph B0 ["ジェネシスブロック (#0)"]
        direction TB
        H0["自分のHash: 0000a1b2..."]
        PH0["前Hash: 00000000..."]
        TX0["データ: Genesis"]
    end

    subgraph B1 ["ブロック #1"]
        direction TB
        H1["自分のHash: 0000c3d4..."]
        PH1["前Hash: 0000a1b2..."]
        TX1["データ: ZKP登録申請"]
    end

    subgraph B2 ["ブロック #2"]
        direction TB
        H2["自分のHash: 0000e5f6..."]
        PH2["前Hash: 0000c3d4..."]
        TX2["データ: Vouchスコア譲渡"]
    end

    H0 -- "暗号学的な鎖 (PrevHash)" --> PH1
    H1 -- "暗号学的な鎖 (PrevHash)" --> PH2

equenceDiagram
    autonumber
    actor User as ユーザー / 中小企業 (Tier 2)
    participant P2P as P2Pゴシップ通信
    participant Pool as 未承認プール (Mempool)
    participant Node as 検証ノード群 (Tier 1)
    participant Ledger as ブロックチェーン (台帳)

    User->>P2P: ① トランザクション発行 (ZKP / Vouch署名)
    P2P->>Pool: ② 各ノードのメモリ領域へバケツリレー伝播
    Pool->>Node: ③ ブロック生成ノードが未承認Txを集約
    Node->>Node: ④ BFT合意形成 (全体の2/3以上の署名検証)
    Node->>Ledger: ⑤ 新ブロックを末尾へ接続 (Hashチェーン更新)
    Ledger-->>User: ⑥ 状態更新完了 (ライトクライアントで一瞬で検証)

