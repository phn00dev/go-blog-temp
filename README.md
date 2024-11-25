### Golang blog project with template


#### Router Gin framework
    go get -u github.com/gin-gonic/gin

#### Configuration Loader with Viper
    go get github.com/spf13/viper

#### Code : 
    func configSet() config.Config {
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.AddConfigPath("config")
        viper.AddConfigPath(".")
        if err := viper.ReadInConfig(); err != nil {
            fmt.Println("read config yaml file error")
        }

        var configs config.Config
        err := viper.Unmarshal(&configs)
        if err != nil {
            fmt.Printf("unable to decode into struct, %v", err)
        }
        return configs
    }