package main

import (
  "fmt"
  "flag"
  "log"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/autoscaling"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "path/filepath"
  "os/user"
)

var (
  group = flag.String("group", "", "The group you wish to set to 1")
  region = flag.String("region", "ap-southeast-2", "Region")
  awsConfig = flag.String("awscfg", filepath.Join(homeDir(), ".aws/credentials"), "Config file to load credentails from")
  awsProfile = flag.String("awsprofile", "default", "Profile to use")
)

func homeDir() string {
  u, err := user.Current()
  if err != nil {
    return ""
  }
  return u.HomeDir
}

func main() {
  if flag.Parse(); !flag.Parsed() {
    log.Panicf("Error with flags")
  }
  if group == nil {
    log.Panicf("Please specify name of autoscale group")
  }

  sess := session.Must(session.NewSession())

  if awsConfig != nil && awsProfile != nil {
    sess.Config.Credentials = credentials.NewSharedCredentials(*awsConfig, *awsProfile)
  }

  if region != nil {
    sess.Config.Region = region
  }

  svc := autoscaling.New(sess)

  params := &autoscaling.UpdateAutoScalingGroupInput{
    AutoScalingGroupName: aws.String(*group),
    DesiredCapacity:      aws.Int64(1),
    MaxSize: aws.Int64(1),
    MinSize: aws.Int64(1),
  }
  resp, err := svc.UpdateAutoScalingGroup(params)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(resp)

}