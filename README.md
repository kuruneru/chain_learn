id
graph TB
    subgraph Flow ["【処理フロー】トランザクションが記録されるまで"]
        User["ユーザー"] -->|"1. Tx発行"| P2P["P2P通信"]
        P2P -->|"2. 伝播"| Pool["メモリプール"]
        Pool -->|"3. BFT検証"| Node["検証ノード"]
        Node -->|"4. 新ブロック生成"| H2
    end

    subgraph Structure ["【データ構造】ハッシュによる鎖の繋がり"]
        subgraph Block0 ["No.0 ジェネシス"]
            H0["Hash: 0000a1b2"]
            TX0["Data: Genesis"]
        end

        subgraph Block1 ["No.1 ブロック"]
            H1["Hash: 0000c3d4"]
            PH1["前Hash: 0000a1b2"]
            TX1["Data: ZKP登録"]
        end

        subgraph Block2 ["No.2 ブロック"]
            H2["Hash: 0000e5f6"]
            PH2["前Hash: 0000c3d4"]
            TX2["Data: Vouch譲渡"]
        end

        H0 --> PH1
        H1 --> PH2
    end
