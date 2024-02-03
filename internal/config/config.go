package config

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/Falokut/admin_movies_persons_service/internal/repository"
	"github.com/Falokut/admin_movies_persons_service/pkg/jaeger"
	"github.com/Falokut/admin_movies_persons_service/pkg/metrics"
	logging "github.com/Falokut/online_cinema_ticket_office.loggerwrapper"
	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type DialMethod = string

const (
	Insecure                 DialMethod = "INSECURE"
	NilTlsConfig             DialMethod = "NIL_TLS_CONFIG"
	ClientWithSystemCertPool DialMethod = "CLIENT_WITH_SYSTEM_CERT_POOL"
	Server                   DialMethod = "SERVER"
)

type ConnectionSecureConfig struct {
	Method DialMethod `yaml:"dial_method"`
	// Only for client connection with system pool
	ServerName string `yaml:"server_name"`
	CertName   string `yaml:"cert_name"`
	KeyName    string `yaml:"key_name"`
}

type Config struct {
	LogLevel        string `yaml:"log_level" env:"LOG_LEVEL"`
	HealthcheckPort string `yaml:"healthcheck_port" env:"HEALTHCHECK_PORT"`

	Listen struct {
		Host string `yaml:"host" env:"HOST"`
		Port string `yaml:"port" env:"PORT"`
		Mode string `yaml:"server_mode" env:"SERVER_MODE"` // support GRPC, REST, BOTH
	} `yaml:"listen"`

	PrometheusConfig struct {
		Name         string                      `yaml:"service_name" ENV:"PROMETHEUS_SERVICE_NAME"`
		ServerConfig metrics.MetricsServerConfig `yaml:"server_config"`
	} `yaml:"prometheus"`

	ImageStorageService struct {
		StorageAddr      string                 `yaml:"storage_addr" env:"IMAGE_STORAGE_ADDRESS"`
		ConnectionConfig ConnectionSecureConfig `yaml:"connection_config"`
		BasePhotoUrl     string                 `yaml:"base_photo_url" env:"BASE_PHOTO_URL"`
		PhotoCategory    string                 `yaml:"photo_category" env:"PHOTO_CATEGORY"`
	} `yaml:"image_storage_service"`

	ImageProcessingService struct {
		Addr                 string                 `yaml:"addr" env:"IMAGE_PROCESSING_ADDRESS"`
		ConnectionConfig     ConnectionSecureConfig `yaml:"connection_config"`
		ImageResizeMethod    string                 `yaml:"resize_type" env:"RESIZE_TYPE"`
		ProfilePictureHeight int32                  `yaml:"photo_height" env:"PHOTO_HEIGHT"`
		ProfilePictureWidth  int32                  `yaml:"photo_width" env:"PHOTO_WIDTH"`
		AllowedTypes         []string               `yaml:"allowed_types"`
		MaxImageWidth        int32                  `yaml:"max_image_width" env:"MAX_IMAGE_WIDTH"`
		MaxImageHeight       int32                  `yaml:"max_image_height" env:"MAX_IMAGE_HEIGHT"`
		MinImageWidth        int32                  `yaml:"min_image_width" env:"MIN_IMAGE_WIDTH"`
		MinImageHeight       int32                  `yaml:"min_image_height" env:"MIN_IMAGE_HEIGHT"`
	} `yaml:"image_processing_service"`

	DBConfig     repository.DBConfig `yaml:"db_config"`
	JaegerConfig jaeger.Config       `yaml:"jaeger"`
	KafkaConfig  struct {
		Brokers []string `yaml:"brokers"`
	} `yaml:"kafka"`
}

var instance *Config
var once sync.Once

const configsPath = "configs/"

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		instance = &Config{}

		if err := cleanenv.ReadConfig(configsPath+"config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Fatal(help, " ", err)
		}
	})

	return instance
}

func (c ConnectionSecureConfig) GetGrpcTransportCredentials() (grpc.DialOption, error) {
	if c.Method == Insecure {
		return grpc.WithTransportCredentials(insecure.NewCredentials()), nil
	}

	if c.Method == NilTlsConfig {
		return grpc.WithTransportCredentials(credentials.NewTLS(nil)), nil
	}

	if c.Method == ClientWithSystemCertPool {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return grpc.EmptyDialOption{}, err
		}
		return grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, c.ServerName)), nil
	}

	cert, err := tls.LoadX509KeyPair(c.CertName, c.KeyName)
	if err != nil {
		return grpc.EmptyDialOption{}, err
	}
	return grpc.WithTransportCredentials(credentials.NewServerTLSFromCert(&cert)), nil
}
