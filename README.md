# go-pet-api
https://petstore.swagger.io

## 要件
Swaggerドキュメントには記載ないが、よくありそうな要件を想像
- 管理者はpetを任意に登録/取得/更新/削除できる
- categoryは事前にまたはpet作成時存在しない場合に作成され、petは必ずいずれか1つのcategoryに紐づく
- tagはpet作成時存在しない場合に作成され、petに0つ以上のtagが紐づく
- tagは必ず1つ以上のペットに紐づき、Update/Deleteで紐づくpetが無くなった場合は同時に削除される

リクエストボディ例
```json
{
  "id": 0,
  "category": {
    "id": 0,
    "name": "string"
  },
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ],
  "status": "available"
}
```

## ER図
![mermaid-diagram-2023-10-24-173849](https://github.com/rare0b/go-pet-api/assets/125894090/c13d09f1-1f01-4994-b8a9-637ce47dff3f)

## メモ
別件で触れる機会のある技術を取り入れたい
- [x] chi
- [x] sqlx
- [x] wire
- [ ] golangci-lint
- [ ] gomock
- [x] swagger
- [ ] coverage
