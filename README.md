# go-pet-api
https://petstore.swagger.io

## 要件
Swaggerドキュメントには記載なし、よくありそうな要件を想像
- 管理者はpetを任意に登録/取得/更新/削除できる
- categoryは事前にまたはpet作成時に作成され、管理者がpetに紐付ける
- tagは管理者が任意に0つ以上設定し、他のpetに設定されたタグを再利用できる
- tagは必ず1つ以上のペットに紐づき、Update/Deleteで紐づくpetが無くなった場合は同時に削除される

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
