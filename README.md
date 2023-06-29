# data-platform-api-production-order-reads-rmq-kube

data-platform-api-production-order-reads-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API で製造指図データを取得するマイクロサービスです。  
https://xxx.xxx.io/api/API_PRODUCTION_ORDER_SRV/reads/

## 動作環境

data-platform-api-production-order-reads-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-production-order-reads-rmq-kube が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_PRODUCTION_ORDER_SRV/reads/

## 本レポジトリ に 含まれる API名
data-platform-api-production-order-reads-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_Header（データ連携基盤 オーダー - ヘッダデータ）
* A_HeaderDoc（データ連携基盤 オーダー - ヘッダ文書データ）
* A_Partner（データ連携基盤 オーダー - 取引先データ）
* A_Address（データ連携基盤 オーダー - 取引先データ）
* A_Item（データ連携基盤 オーダー - 明細データ）
* A_ItemComponent（データ連携基盤 オーダー - 明細構成品目データ）
* A_ItemComponentDeliveryScheduleLine（データ連携基盤 オーダー - 明細構成品目納入日程行データ）
* A_ItemComponentCosting（データ連携基盤 オーダー - 明細構成品目原価計算データ）
* A_ItemOperation（データ連携基盤 オーダー - 明細作業データ）
* A_ItemOperationComponent（データ連携基盤 オーダー - 明細作業構成品目データ）
* A_ItemOperationCosting（データ連携基盤 オーダー - 明細作業原価計算データ）

## API への 値入力条件 の 初期値
data-platform-api-production-order-reads-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMProdctionReads",
	"accepter": ["Header"],
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMProdctionReads",
	"accepter": ["All"],
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncProductionOrderReads(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}
	// SQL処理
	response = c.readSqlProcess(nil, &mtx, input, output, accepter, &errs, log)

	return response, nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 製造指図 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ProductionOrder" ～ "PlusMinusFlag" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
