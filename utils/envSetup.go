package utils

import "os"

func SetEnv() {
	os.Setenv("TOKEN_HOUR_LIFESPAN", "8")
	os.Setenv("API_SERCET", "88fb65e144526e6f96451e768d94da3a17f86e31105ea1454b75461c4a452e39")
}
