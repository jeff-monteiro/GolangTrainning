fun main(){
    println("Welcome to the Bytebank")
    val holder: String = "Jefferson Monteiro"
    val accountNumber: Int = 1045
    var balanceAccount: Double = 0.00
    balanceAccount = 100.00
    balanceAccount += 20.00
    balanceAccount -= 1000

    println("Holder: $holder")
    println("Account number: $accountNumber")
    println("Balance: $balanceAccount")

    // The WHEN expression can be used to build this example.
    if(balanceAccount > 0.00){
        println("Account has positive balance")
    } else if(balanceAccount == 0.00){
        println("Neutral balance")
    } else {
        println("Negative balance")
    }
}
