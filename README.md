# go-s3-updown-server

#### Demo for how to upload and download files to s3 via a a golang web server.

### Tasks:

- [x] create basic webpage for webserver to upload and download files locally
  - [x] upload
    - [x] create form
    - [x] create handling request
	- [x] save file
  - [x] download
    - [x] create url that will download via link
    - [x] create routing to download
- [x] upload file to s3 bucket
  - [x] create bucket with authentication [great article](https://github.com/keithweaver/python-aws-s3)
  - [x] upload file in local filesystem authenticated in script [aws-sdk-go-docs](https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/s3_upload_object.go)
  - [x] successfully upload file authenticated without saving to local filesystem
- [x] show list of files in s3 bucket [aws-sdk-go-docs](https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/s3_list_objects.go)
- [x] download file from s3 bucket without saving to local file system.
- [ ] refactor for DRY
- [ ] document how download works
- [ ] PR to keep abstraction layer in gin gonic with download endpoint
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
