### Linux

```shell
curl -L https://github.com/jenkins-x/jx-ui/releases/download/v{{.Version}}/jx-ui-linux-amd64.tar.gz | tar xzv
sudo mv jx-ui /usr/local/bin
```

### macOS

```shell
curl -L  https://github.com/jenkins-x/jx-ui/releases/download/v{{.Version}}/jx-ui-darwin-amd64.tar.gz | tar xzv
sudo mv jx-ui /usr/local/bin
```
