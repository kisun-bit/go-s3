package go_s3

import (
	"github.com/journeymidnight/aws-sdk-go/service/s3"
)

func (s3client *S3Client) SetBucketWebsite(bucketName string, rules []*s3.CORSRule) error {
	//param := s3.PutBucketWebsiteInput{
	//	Bucket: aws.String(bucketName),
	//	WebsiteConfiguration: &s3.WebsiteConfiguration{
	//		IndexDocument: &s3.IndexDocument{
	//			Suffix: aws.String("index.html"),
	//		},
	//		ErrorDocument: &s3.ErrorDocument{
	//			Key: aws.String("error.html"),
	//		},
	//	},
	//}
	return nil
}