package providers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/fatih/color"
	"go.uber.org/fx"
)

// NewTerminalService initializes and displays terminal messages and checks database connectivity
func NewTerminalService(cfg *config.AppConfig, lc fx.Lifecycle, db *DatabaseService, cache *CacheService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			time.Sleep(3 * time.Second)
			fmt.Println("≿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━༺❀༻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━≾")
			fmt.Println()
			printASCIIArt()
			printAppDetails(cfg, db, cache)
			fmt.Println()
			fmt.Println("≿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━༺❀༻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━≾")
			return nil
		},
	})
}

// printASCIIArt displays the application ASCII art with colored output
func printASCIIArt() {
	asciiArt := `
	        ..............                            
            .,,,,,,,,,,,,,,,,,,,                             
        ,,,,,,,,,,,,,,,,,,,,,,,,,,                          
      ,,,,,,,,,,,,,,  .,,,,,,,,,,,,,                        
    ,,,,,,,,,,           ,,,,,,,,,,,                     
    ,,,,,,,          .,,,,,,,,,,,                          
  @@,,,,,,          ,,,,,,,,,,,,                             
@@@,,,,.@@      .,,,,,,,,,,,                                
@,,,,,,,@@    ,,,,,,,,,,,                                   
  ,,,,@@@       ,,,,,,                                      
    @@@@@@@                                          
    @@@@@@@@@@           @@@@@@@@                          
      @@@@@@@@@@@@@@  @@@@@@@@@@@@                          
        @@@@@@@@@@@@@@@@@@@@@@@@@@                          
            @@@@@@@@@@@@@@@@@@@@                             
                  @@@@@@@@
	`

	navyBlue := color.New(color.BgBlue).SprintFunc()
	lightGreen := color.New(color.BgGreen).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()

	lines := strings.Split(asciiArt, "\n")

	for _, line := range lines {
		coloredLine := ""
		for _, char := range line {
			switch char {
			case '@':
				coloredLine += navyBlue(" ")
			case ',', '.':
				coloredLine += lightGreen(" ")
			default:
				coloredLine += white(string(char))
			}
		}
		fmt.Println(coloredLine)
	}
}

// printAppDetails prints the application details and environment information
func printAppDetails(cfg *config.AppConfig, db *DatabaseService, cache *CacheService) {
	if cfg == nil {
		color.Red("❌ Error: Application config is missing")
		return
	}

	// Define color schemes
	appName := color.New(color.FgHiCyan, color.Bold).SprintFunc()
	info := color.New(color.FgHiGreen).SprintFunc()
	success := color.New(color.FgHiBlue).SprintFunc()
	errorColor := color.New(color.FgHiRed).SprintFunc()
	version := color.New(color.FgYellow).SprintFunc()

	// Print application messages
	fmt.Println(appName("Engine: ", cfg.AppName))
	fmt.Println()
	fmt.Println(info("➥ Environment: ") + success(cfg.AppEnv))
	fmt.Println()

	fmt.Println(info("➥ Localhost is running on: ") + success(fmt.Sprintf("https://localhost:%s", cfg.AppPort)))

	// Database connection status
	if err := db.Ping(); err != nil {
		fmt.Println(errorColor("➥ ❌  Database connection failed: ") + err.Error())
	} else {
		fmt.Println(success("➥ ✅  Database running successfully"))
	}

	if err := cache.Ping(); err != nil {
		fmt.Println(errorColor("➥ ❌  Database connection failed: ") + err.Error())
	} else {
		fmt.Println(success("➥ ✅  Cache server running successfully"))
	}

	fmt.Println()

	fmt.Println()
	fmt.Println(version("✒ Version ", cfg.AppVersion))
	fmt.Println()
}
