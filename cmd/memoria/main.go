package memoria
import (
	"log"
	"github.com/11himanshusharma/memoria/internal/cli"
)

func main(){
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}