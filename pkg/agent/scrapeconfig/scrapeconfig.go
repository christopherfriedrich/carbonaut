/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package scrapeconfig

type Config struct {
	AwsTargetConfig *AWSTargetConfig
}

type AWSTargetConfig struct {
	Regions []string
}
