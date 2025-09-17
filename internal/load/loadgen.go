package load

import (
	
	"fmt"
	

	"github.com/Black-tag/kafka-sampler/internal/kafka/publisher"
	
)


func Generate(producer *publisher.Producer) {
    for i := 0; i < 5; i++ {
        msg := fmt.Sprintf("message-%d", i)
		if err := producer.SendMessage("key", msg); err != nil {
			fmt.Println("error in producing message", err)
		} else {
			fmt.Println("produced:", msg)
		}
        
    }
}

