package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"wetees.com/domain"
)

func NewWelcomeRoute(conf *domain.Config, app fiber.Router) {
	fmt.Printf("	Powered by %s %s\n", conf.AppName, conf.Version)

	banner := `
	WWWWWWWW                           WWWWWWWW            TTTTTTTTTTTTTTTTTTTTTTT              EEEEEEEEEEEEEEEEEEEEEE                 
	W::::::W                           W::::::W            T:::::::::::::::::::::T              E::::::::::::::::::::E                 
	W::::::W                           W::::::W            T:::::::::::::::::::::T              E::::::::::::::::::::E                 
	W::::::W                           W::::::W            T:::::TT:::::::TT:::::T              EE::::::EEEEEEEEE::::E                 
	W:::::W           WWWWW           W:::::W eeeeeeeeeeeeTTTTTT  T:::::T  TTTTTTeeeeeeeeeeee    E:::::E       EEEEEE    ssssssssss   
	W:::::W         W:::::W         W:::::Wee::::::::::::ee      T:::::T      ee::::::::::::ee  E:::::E               ss::::::::::s  
	W:::::W       W:::::::W       W:::::We::::::eeeee:::::ee    T:::::T     e::::::eeeee:::::eeE::::::EEEEEEEEEE   ss:::::::::::::s 
	W:::::W     W:::::::::W     W:::::We::::::e     e:::::e    T:::::T    e::::::e     e:::::eE:::::::::::::::E   s::::::ssss:::::s
	W:::::W   W:::::W:::::W   W:::::W e:::::::eeeee::::::e    T:::::T    e:::::::eeeee::::::eE:::::::::::::::E    s:::::s  ssssss 
	W:::::W W:::::W W:::::W W:::::W  e:::::::::::::::::e     T:::::T    e:::::::::::::::::e E::::::EEEEEEEEEE      s::::::s      
	W:::::W:::::W   W:::::W:::::W   e::::::eeeeeeeeeee      T:::::T    e::::::eeeeeeeeeee  E:::::E                   s::::::s   
	W:::::::::W     W:::::::::W    e:::::::e               T:::::T    e:::::::e           E:::::E       EEEEEEssssss   s:::::s 
	W:::::::W       W:::::::W     e::::::::e            TT:::::::TT  e::::::::e        EE::::::EEEEEEEE:::::Es:::::ssss::::::s
	W:::::W         W:::::W       e::::::::eeeeeeee    T:::::::::T   e::::::::eeeeeeeeE::::::::::::::::::::Es::::::::::::::s 
	W:::W           W:::W         ee:::::::::::::e    T:::::::::T    ee:::::::::::::eE::::::::::::::::::::E s:::::::::::ss  
	WWW             WWW            eeeeeeeeeeeeee    TTTTTTTTTTT      eeeeeeeeeeeeeeEEEEEEEEEEEEEEEEEEEEEE  sssssssssss    
	`
	fmt.Println(banner)
}
