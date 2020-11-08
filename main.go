//PARA UTILIZAR ESTE PROGRAMA PRIMERO NECESITAMOS UTILIZAR EL CODIGO COMENTADO
//ESE CODIGO GENERARE UNA PRIVATE KEY Y UNA PUBLIC KEY
//LUEGO CON ESA PUBLIC KEY NOSOTROS PODREMOS ENCRIPTAR MENSAJES QUE NUESTRO
//PRIVATE KEY GENERADO VA A PODER DESCIFRAR
//LUEGO DE TENER EL MENSAJES DESCIFRADO TENEMOS QUE PASAR POR EL RECIEVER

package main

import (
	"fmt"
	"log"
	"rsapem/internal/reciever"
)

func main() {
	encryptedMessage := `-----BEGIN MESSAGE-----
Iv7MSqrqDeWW5H9MMdBPn9WzsqgUa0TYBn+zakB8To10AyNK/HNmvS4sfCHSZQ77
8INr/sojF2w1tj0CUbtg87Bl5qmFL7ldl7QfYFPPrN8FRSUy3tu093LQ2LewJEns
jF16mAriHG2iUswpDOVCLRW6IcnDlTfOygeUOMy0Da2cyGr7RpdKax2P+lCq7wCC
Ecz3VWMff9kVIUA8sWhC7Af6nPP6JXRxk3VGZT3J0p3J5tEnlQx33BbrM38q465O
+jaI/kp0GHzgmurOXlPAlpUgMoyFOJwCxHSIkzKMiTzE0v0YkfONqCOUOJqkhEya
sJuvq+cMI0BudmRsHeO+9w==
-----END MESSAGE-----
`
	decryptedMessage, err := reciever.Decrypt("./certs/private.key", encryptedMessage)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(decryptedMessage)
}

/*
package main

import (
	"fmt"
	"log"
	"rsapem/internal/sender"
)

func main() {
	plainText := "This is a very secret message :)"

	encryptedMessage, err := sender.Encrypt("./certs/public.key", plainText)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(encryptedMessage)
}
*/
/*
package main

import (
	"fmt"
	"log"
	"rsapem/internal/pki"
)

func main() {
	key, err := pki.New()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(key.PublicKeyToPemString())
	fmt.Println(key.PrivateKeyToPemString())
}*/
