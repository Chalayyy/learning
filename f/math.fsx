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

let {First=first; Last=last} = me // This is the same as below


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