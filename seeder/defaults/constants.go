package defaults

import (
	"time"

	"github.com/gofrs/uuid"
)

// Date is a global date
var Date = time.Unix(1530720000, 0)

// Users
var User1 = uuid.Must(uuid.FromString("6c3d9452-9714-445c-af70-20fc018fb8e6"))
var User2 = uuid.Must(uuid.FromString("674de1ad-80b7-4a8b-86c2-de3e58816e85"))

// Stores
var Store1 = uuid.Must(uuid.FromString("bdd2246c-9900-4bf2-afd9-b4d9a8c63d2a"))
var Store2 = uuid.Must(uuid.FromString("e75f95cf-4b13-4474-b2e7-94eac9d5ee27"))

// Purchases
var Purchase1 uint64 = 1
var Purchase2 uint64 = 2
var Purchase3 uint64 = 3

// Reviews
var Review1 = uuid.Must(uuid.FromString("b7b47c12-df18-487e-a240-01d7777e6137"))
var Review2 = uuid.Must(uuid.FromString("4f518635-f5fa-4084-a66a-0aa3d5c763ec"))
var Review3 = uuid.Must(uuid.FromString("6c8e055c-dd24-459b-aa68-e51de100e227"))
var Review4 = uuid.Must(uuid.FromString("bf7ee74a-61cd-4483-873b-63f9cbe983ed"))
var Review5 = uuid.Must(uuid.FromString("f097f069-b590-437c-98c5-c0ebe1eb273a"))
var Review6 = uuid.Must(uuid.FromString("89b3803a-32a1-483a-b4dc-fb69fbb4e64f"))
