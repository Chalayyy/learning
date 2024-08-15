let add1 x = x + 1  // signature is: int -> int​​ ​ 
let add x y = x + y // signature is: int -> int -> int​

let squarePlusOne x =
    let square = x * x
    square + 1

let areEqual x y = (x = y)

type AppleVariety =
    | Gala
    | Fuji
    | GrannySmith

type BananaVariety =
    | Cavendish
    | Golden
    | Plantain

type CherryVariety =
    | Maraschino
    | Fuji
    | Bing

type FruitSalad = {
    Apple: AppleVariety
    Banana: BananaVariety
    Cherries: CherryVariety
}

type FruitSnack =
| Apple of AppleVariety
| Banana of BananaVariety
| Cherries of CherryVariety

type Name = Name of string

type Person = {First:string; Last:string}

let me = {First="Charlie"; Last="Woudlntyouliketoknow"}
let wife = {First="Jen"; Last="Woudlntyouliketoknow"}

let {First=first; Last=last} = me // This is the same as below
let firstWife = wife.First
let lastWife = wife.Last

type OrderQuantity =
| UnitQuantity of int
| KilogramQuantity of float

let anOrderQtyInUnits = UnitQuantity 10
let anOrderQtyInKg = KilogramQuantity 2.5

let x = typedefof<OrderQuantity>

let printQuantity aOrderQty =
    match aOrderQty with
    | UnitQuantity uQty ->
        printfn "%i units" uQty
    | KilogramQuantity kgQty ->
    printfn "%g kg" kgQty

printQuantity anOrderQtyInUnits
printQuantity anOrderQtyInKg


// define some types
type CustomerId = CustomerId of int
type OrderId = OrderId of int

// define some values
let customerId = CustomerId 42
let orderId = OrderId 42

// try to compare them -- compiler error!
printfn "%b" (orderId = customerId)
//                      ^ This expression was expected to
//                      have type 'OrderId'


// construct
let customerId = CustomerId 42
// deconstruct
let (CustomerId innerValue) = customerId
//              ^ innerValue is set to 42
printfn "%i" innerValue // prints "42"

// deconstruct
let processCustomerId (CustomerId innerValue) = printfn "innerValue is %i" innerValue
// function signature
// val processCustomerId: CustomerId -> unit
