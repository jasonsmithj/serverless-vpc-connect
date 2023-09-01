## local環境
```bash
go run main.go
```

### GAE Standard Environment
#### 環境設定
export GCP_PROJECT=<your gcp project>

#### config作成
```bash
cp app_sample.yaml app.yaml
```

gcloud app deploy --project=${GCP_PROJECT} --version=$(git show --format='%h' --no-patch ) --quiet

