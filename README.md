# go-s3-updown-server

#### Demo for how to upload and download files to s3 via a a golang web server.

### TODO:
- [ ] refactor and clean up
  - [ ] create s3 package to abstract, make s3 struct.
  - [ ] share session between s3 calls
  - [ ] document how download works
  - [ ] PR to keep abstraction layer in gin gonic with download endpoint
- [ ] delete files
- [ ] write how to deploy the server... fargate would be cool
- [ ] provision bucket
   - [ ] create terrform for bucket and keep it in repo
- [ ] Add name and uploaded date
- [ ] add drag and drop
- [ ] Write blog entry
- [ ] research the following fields
  - [ ] ContentLength:        aws.Int64(size)
  - [ ] ContentType:          aws.String(http.DetectContentType(buffer))
  - [ ] ContentDisposition:   aws.String("attachment")
  - [ ] ServerSideEncryption: aws.String("AES256")
- [ ] [Filename SHOULD NOT be trusted. See Content-Disposition on MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives)
- [ ] add directories
